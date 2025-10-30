package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	var s, t string
	fmt.Scan(&n, &m) // nolint:errcheck
	fmt.Scan(&s, &t) /// nolint:errcheck

	if strings.HasPrefix(t, s) {
		if strings.HasSuffix(t, s) {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	} else if strings.HasSuffix(t, s) {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
