package main

func fibonacci(position int) func() int {
	for i := 0; i < position; i++ {
		current, next, nextnext := 0, 1, 1
		return func() int {
			ret := current

			current = next
			next = nextnext
			nextnext = current + next
			return ret
		}()
	}
}

func main() {
	fibonacci(10)

}
