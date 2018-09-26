// package main

// import (
// 	"fmt"
// 	"math"
// )

// func mergeSort(slice []int) []int {
// 	if len(slice) == 1 {
// 		return slice
// 	}
// 	middleIndex := int(math.Floor(float64(len(slice) / 2)))
// 	firstHalf := slice[:middleIndex]
// 	secondHalf := slice[middleIndex:]

// 	return merge(mergeSort(firstHalf), mergeSort(secondHalf))
// }

// func merge(s1, s2 []int) []int {
// 	fmt.Println("21 Working")
// 	fmt.Println("I'm s1", s1[0])
// 	fmt.Println("I'm s2", s2[0])
// 	fmt.Println("I'm len", len(s1), "\t", len(s2))
// 	fmt.Println("derrrrrrrr", s1, " ---- ", s2)
// 	// for len(s1) > 0 && len(s2) > 0 {
// 	// 	fmt.Println("23 Working")
// 	// 	if s1[0] < s2[0] {
// 	// 		result = append(result, s1[0])
// 	// 	} else {
// 	// 		result = append(result, s2[0])
// 	// 	}
// 	// }
// 	fmt.Println("29 Working")

// 	if len(s1) > 0 {
// 		result = append(result, s1...)
// 	} else {
// 		result = append(result, s2...)
// 	}
// 	fmt.Println("28 Working")

// 	return result
// }

// func main() {
// 	fmt.Println(mergeSort([]int{4, 3, 2, 1}))
// }
