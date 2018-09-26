package main

import "fmt"

func linearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{95, 78, 46, 58, 86, 34, 99, 151, 420, 320}
	fmt.Println(linearSearch(items, 420))
}
