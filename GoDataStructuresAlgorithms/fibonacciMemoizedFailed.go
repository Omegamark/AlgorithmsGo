package main

func fibMemo(index int, cache []int) int {
	if val, ok := cache[index]; ok {
		return val
	} else if index < 3 {
		return 1
	} else {
		cache[index] = fibMemo(index-1, cache) + fibMemo(index-2, cache)
	}
	return cache[index]
}

func main() {
	fibMemo(20, []int{0})
}
