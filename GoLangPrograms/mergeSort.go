package main

import (
	"time"
    "fmt"
    "math/rand"
)

func mergeSort(slice []int) []int {
	num := len(slice)

	if num == 1 {
		return slice
	}

	middleIndex := int(num / 2)

	var (
		left  = make([]int, middleIndex)
		right = make([]int, num-middleIndex)
	)

    for i := 0; i < num; i++ {
        if i < middleIndex {
            left[i] = slice[i]
        } else {
            right[i - middleIndex] = slice[i]
        }
    }

    return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
    result := make([]int, len(left) + len(right))

    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            fmt.Println("I'm left before being cut", left)

            left = left[1:]
        } else {
            result[i] = right[0]
            fmt.Println("I'm right before being cut", right)
            right = right[1:]
        }
        i++
    }

    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
    return result
}

func generateSlice(size int) []int {
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func main() {
    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    fmt.Println("\n---  Sorted  ---\n\n", mergeSort(slice))
}