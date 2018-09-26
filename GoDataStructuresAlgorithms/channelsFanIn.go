package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}

// This arrow is a write only channel.
func producer(ch chan<- int, name string) {
	for {
		// Sleep for some random time
		sleep()

		// generate a random number
		n := rand.Intn(100)

		// send the message
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

// This arrow is a read only channel.
func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func main() {
	// chA & chB are the fanOut
	chA := make(chan int)
	chB := make(chan int)
	// chC is the fanIn or Orchestration Channel.
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chA, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
