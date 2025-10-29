package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	const numbers = "0123456789"

	for _, n := range numbers {
		if !strings.ContainsRune(s, n) {
			fmt.Printf("%c\n", n)
		}
	}
}
