package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	const maximumNumber = 1000000000 // 10^9

	result := 0.0

	for i := 0; i <= m; i++ {
		result += math.Pow(float64(n), float64(i))
	}
	if result <= maximumNumber {
		fmt.Println(int(result))
	} else {
		fmt.Println("inf")
	}
}
