package main

import "fmt"

func main() {
	slice := []int{5, 3, 8, 2, 1, 4}
	fmt.Println("\n--- Unsorted ---\n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n---  Sorted  ---\n\n", slice, "\n")
}

func bubbleSort(slice []int) {
	n := len(slice)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--
	}
}
