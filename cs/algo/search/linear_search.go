package main

import (
	"fmt"
)

// LinearSearch in Go (index version)
//
// Returns the index of the first element that matches the target.
// Returns -1 if the target is not found.
//
// Time Complexity: O(N)
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1 // index does not become -1, so we use -1 to indicate "not found"
}

// LinearSearchValue in Go (value version)
//
// Returns the *value* of the first element that matches the target.
// Returns -1 if the target is not found.
//
// Time Complexity: O(N)
func LinearSearchValue(arr []int, target int) (int, bool) {
	for _, v := range arr {
		if v == target {
			return v, true
		}
	}
	return -1, false // value possibly -1, so we return a bool to indicate "not found"
}

func main() {
	arr := []int{10, 20, 30, 40, 50}

	// --- Demo ---
	fmt.Println(LinearSearch(arr, 30))      // Output: 2
	fmt.Println(LinearSearchValue(arr, 30)) // Output: 30
}
