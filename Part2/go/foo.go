// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

func number_server(add_number <-chan int, get_number chan<- int, quit chan bool) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.

		case increment := <-add_number:
			i += increment
		case get_number <- i:
		case <-quit:
			return
		}
	}
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	finished <- true
		return
	//TODO: signal that the goroutine is finished
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}
	finished <- true
		return
	//TODO: signal that the goroutine is finished
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	add_number 	:= make(chan int)
	get_number 	:= make(chan int)
	quit 		:= make(chan bool)
	finished 	:= make(chan bool)
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.

	// TODO: Spawn the required goroutines
	go number_server(add_number, get_number, quit)
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)

	// TODO: block on finished from both "worker" goroutines
	<-finished
	<-finished
	
	Println("The magic number is:", <-get_number)
	quit<-true
}
