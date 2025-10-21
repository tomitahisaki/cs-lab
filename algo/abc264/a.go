package main

import "fmt"

func main() {
	var l, r int
	_, _ = fmt.Scan(&l, &r)

	const string = "atcoder"

	fmt.Println(string[l-1 : r])
}
