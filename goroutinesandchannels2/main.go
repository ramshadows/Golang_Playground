package main

/*
Concurrency is the composition of independent activities

A goroutine is a concurrent activity

A channel in Go is a communication mechanism between goroutines.
When you need to send a value from one goroutine to another, you use channels.
channel is a communication mechanism that sends and receives data, it also has a type.
That means you can send data for only the kind that the channel supports.
You use the keyword chan as the data type for a channel, but you also need to specify the
data type that will pass through the channel, like an int type.

e.g ch := make(chan int)

A channel can do two operations: send data and receive data.
you have to use the <- operator after the channel. When you want the channel to receive data,
you have to use the <- operator before the channel,

e.g

ch <- x // sends (or write) x through channel ch
x = <-ch // x receives (or reads) data sent to the channel ch
<-ch // receives data, but the result is discarded

channels are unbuffered by default. That means they accept a send operation only if there's a receive operation. Otherwise, the program will be blocked waiting forever.

Buffered channels send and receive data without blocking the program because a buffered channel behaves like a queue. You can limit the size of this queue when you create the channel, like this:

ch := make(chan string, 10)

Every time you send something to a buffered channel, the element is added to the queue. Then, a receive operation removes the element from the queue. When the channel is full, any send operation simply waits until there's space to hold the data. Conversely, if the channel is empty and there's a read operation, it's blocked until there's something to read. buffered channels decouple the send and receive operations. They don't block a program


Another operation that you can use in a channel is to close it. To close a channel, you use the built-in close() function,

e.g close(ch)

****** Channel directions ****
You can specify whether a channel is meant to send or receive data

e.g 

chan<- int // it's a channel to only send data
<-chan int // it's a channel to only receive data

func send(ch chan<- string, message string) {
    fmt.Printf("Sending: %#v\n", message)
    ch <- message
}

func read(ch <-chan string) {
    fmt.Printf("Receiving: %#v\n", <-ch)
}

#main()
ch := make(chan string, 1)
    send(ch, "Hello World!")
    read(ch)
#Output
Sending: "Hello World!"
Receiving: "Hello World!"

*/

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// creates an unbuffered channel
	// Unbuffered channels block the sending operation until there's someone ready to receive the data.
	// stopped as soon as it receives the first message. Need to loop to receive all data expected.
	// unbuffered channels are synchronize the sending and receiving operations.
	// Even though you're using concurrency, the communication is synchronous.

	ch := make(chan string)

	for _, api := range apis {
		// Fetches the apis concurrently using the checkAPI function
		go func(myAPI string, myChan chan string) {
			checkAPI(myAPI, myChan)

		}(api, ch)

	}

	// Need to loop to receive all data expected
	for i := 0; i < len(apis); i++ {
		// notice we are using Print() since the recieving channel
		// in checkAPI() func has end line charecter - \n
		fmt.Print(<-ch)

	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// Test buffered channels
	size := 4

	messages := []string{
		"One",
		"Two",
		"Three",
		"Four",
	}

	// create a buffered channel
	ch1 := make(chan string, size)

	for _, message := range messages {
		// Fetches the apis concurrently using the checkAPI function
		go func(myChan chan string, myMessage string) {
			send(myChan, myMessage)

		}(ch1, message)

	}
	

	// Need to loop to receive all data expected
	for i := 0; i < size; i++ {
		
		fmt.Println(<-ch1)

	}

}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func send(ch chan string, message string) {
	ch <- message
}
