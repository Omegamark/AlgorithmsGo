package main

import "fmt"

func fib(n int) int {
	// Create a map to store the last 2 numbers.
	fn := make(map[int]int)

	for i := 0; i <= n; i++ {
		var f int
		// Prepopulate the map with 1 and 1 in the first 2 positions.
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	// Set the current fibonacci number on the map fn at [i]
	// fn is a map with a key of the position in sequence and a value of the fibonacci number itself.
	fmt.Printf("%v", fn)

	return fn[n]
}

func main() {
	x := fib(3)
	println("\n", x)
}
