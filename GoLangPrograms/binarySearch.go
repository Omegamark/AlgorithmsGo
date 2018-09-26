package main

import "fmt"

// Binary Search must be performed on an already sorted list.
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		// Establish the initial median of the list.
		median := (low + high) / 2

		if haystack[median] < needle {
			// Check if the needle is higher than the initial median.
			// If so, eliminate the everything before the median.
			low = median + 1
		} else {
			// If not, eliminate everything after the median.
			high = median - 1
		}
	}
	// If the low is the length of the haystack, then there are 0 items in the list and there is no need to search.
	// The way the search operates, low will = needle
	// if low == len(haystack), the end of the list was reached and the item being searched was not found.
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	fmt.Println(len(haystack))

	return true
}

func main() {
	items := []int{0, 1, 2, 9, 20, 31, 63, 45, 70, 100}
	fmt.Println(binarySearch(70, items))
}
