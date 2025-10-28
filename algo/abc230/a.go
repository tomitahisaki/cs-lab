package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const title = "AGC"
	if n < 42 {
		fmt.Printf("%s%03d\n", title, n)
	} else {
		fmt.Printf("%s%03d\n", title, n+1)
	}
}
