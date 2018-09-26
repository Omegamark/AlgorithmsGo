package main

import (
	"fmt"
	"math/rand"
	"time"
)

func echoWorker(in, out chan int) {
	for {

		n := <-in

		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

		out <- n
	}
}

func producer(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n", i)
		ch <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	// Produces workers indefinitely
	// go echoWorker(in, out)

	// Produce 4 workers simultaneously, essentially 4 instances of the function are running in parallel.
	for i := 0; i < 4; i++ {
		go echoWorker(in, out)
	}

	go producer(in)

	// Consuming the output channel.
	for n := range out {
		fmt.Printf("<- Recv job: %d\n", n)
	}
}
