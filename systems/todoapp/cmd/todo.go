package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tomitahisaki/cs-lab/systems/todoapp/internal/infra"
	"github.com/tomitahisaki/cs-lab/systems/todoapp/internal/usecase"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(2)
	}

	// wire: infra -> usecase
	repo, err := infra.NewFileTaskRepo("./db/tasks.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "repo init error:", err)
		os.Exit(1)
	}
	uc := usecase.NewTaskUsecase(repo)

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		titleFlag := addCmd.String("title", "", "task title")
		_ = addCmd.Parse(os.Args[2:])

		// 引数でも受けられるように（todo add Buy milk）
		title := strings.TrimSpace(strings.Join(addCmd.Args(), " "))
		if *titleFlag != "" {
			title = *titleFlag
		}
		if title == "" {
			fmt.Fprintln(os.Stderr, "usage: todo add [-title \"...\"] | todo add <title words...>")
			os.Exit(2)
		}

		t, err := uc.Add(title)
		if err != nil {
			fmt.Fprintln(os.Stderr, "add error:", err)
			os.Exit(1)
		}
		fmt.Printf("added: #%d %s\n", t.ID, t.Title)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		_ = listCmd.Parse(os.Args[2:])

		tasks, err := uc.List()
		if err != nil {
			fmt.Fprintln(os.Stderr, "list error:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("(no tasks)")
			return
		}
		for _, t := range tasks {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("[%s] #%d %s\n", mark, t.ID, t.Title)
		}

	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		_ = doneCmd.Parse(os.Args[2:])

		if doneCmd.NArg() < 1 {
			fmt.Fprintln(os.Stderr, "usage: todo done <id>")
			os.Exit(2)
		}
		id, err := strconv.Atoi(doneCmd.Arg(0))
		if err != nil || id <= 0 {
			fmt.Fprintln(os.Stderr, "done: id must be a positive integer")
			os.Exit(2)
		}

		if err := uc.Done(id); err != nil {
			fmt.Fprintln(os.Stderr, "done error:", err)
			os.Exit(1)
		}
		fmt.Printf("done: #%d\n", id)

	case "help", "-h", "--help":
		printUsage()

	default:
		fmt.Fprintln(os.Stderr, "unknown command:", os.Args[1])
		printUsage()
		os.Exit(2)
	}
}

func printUsage() {
	fmt.Println(`todo - simple CLI todo

Usage:
  todo <command> [options]

Commands:
  add   [-title "Write tests"] | todo add Write tests
  list
  done  <id>

Examples:
  todo add -title "Write tests"
  todo add Ship code
  todo list
  todo done 1`)
}
