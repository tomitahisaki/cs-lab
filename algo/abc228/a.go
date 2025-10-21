package main

import "fmt"

func main() {
	var s, t, x int
	fmt.Scan(&s, &t, &x) // nolint:errcheck

	if s < t {
		if s <= x && x < t {
			fmt.Println("Yes")
			return
		}
	} else {
		if s <= x || x < t {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
