package main

import (
	"fmt"
	"time"
)

/*
Interact with more than one channel simultaneously by using the select keyword.

A select statement works like a switch statement but for channels. It blocks the program's execution until it receives an event to process. If it gets more than one event, it chooses one at random.

An essential aspect of the select statement is that it finishes its execution after it processes an event. If you want to wait for more events to happen, you might need to use a loop.
*/

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	// Use loop because the select statement ends as soon as it 
	// receives an event, but we're still waiting for the process 
	// function to finish.
	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}
