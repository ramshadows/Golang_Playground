package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Goroutines let your program
work on several diﬀerent tasks at once. Your goroutines
can coordinate their work using channels, which let them
send data to each other and synchronize so that one
goroutine doesn’t get ahead of another.

Note: Go routines statements can’t be used with return
values

Note2: Not only do channels allow you to send values
from one goroutine to another, they ensure the sending
goroutine has sent the value before the receiving goroutine
attempts to use it.

The only practical way to use a channel is to communicate
from one goroutine to another goroutine. So to
demonstrate channels, we’ll need to be able to do a few
things:
Create a channel.
Write a function that receives a channel as a
parameter. We’ll run this function in a separate
goroutine, and use it to send values over the
channel.
Receive the sent values in our original goroutine.
Each channel only carries values of a particular type, so
you might have one channel for int values, and another
channel for values with a struct type. To declare a variable
that holds a channel, you use the chan keyword, followed by
the type of values that channel will carry.

Note3:

A send operation blocks the sending goroutine until another
goroutine executes a receive operation on the same
channel. And vice versa: a receive operation blocks the
receiving goroutine until another goroutine executes a send
operation on the same channel. This behavior allows
goroutines to synchronize their actions—that is, to
coordinate their timing.
*/

func main() {
	// Declares a variable to hold a channel of float64 values
	//var myChannel chan float64

	// Actually create the channel
	//myChannel = make(chan float64)

	// Declare and Create a channel at aonce

	myChannel1 := make(chan float64)

	// To send a value on a channel, you use the <- operator
	// It looks like an arrow pointing from the value you’re sending to the
	// channel you’re sending it on.
	// Here: myChannel1 is receiving a value of 3.14

	myChannel1 <- 3.14 // 3.14 is the value to send to myChannel1

	// You also use the <- operator to receive values from a channel
	// you place the arrow to the left of the channel you’re receiving from.

	<-myChannel1

	// Declare and create a channel of string values
	myChannel2 := make(chan string)

	// Pass the channel to the greeting function now running a new goroutine
	go greeting(myChannel2)

	// We could also store the received value in a variable
	receivedValue := myChannel2

	// Print format now should receive a value from the channel
	fmt.Println(<-myChannel2) //Output: "Hello"
	fmt.Println(receivedValue)

	pageSize := make(chan int)

	go responseSize("https://www.google.com", pageSize)
	go responseSize("https://www.google.com", pageSize)
	go responseSize("https://www.google.com", pageSize)

	fmt.Println(<-pageSize)

	// Updating to use slice and for loop
	urls := []string{
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
	}

	for _, url := range urls {
		go responseSize(url, pageSize)

	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-pageSize)

	}

	pages := make(chan Page)

	for _, url := range urls {
		go responseSize2(url, pages)

	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s, %d\n", page.URL, page.Size)

	}

}

// Function greeting takes a channel as parameter
func greeting(myChan chan string) {
	// Send a value over the channel
	myChan <- "Hello"
}

func responseSize(url string, channel chan int) {
	fmt.Println("Getting ", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Instead of returning the page size - we send it over the channel
	channel <- len(body)
}

// Updating our channel to carry a struct

type Page struct {
	URL  string
	Size int
}

func responseSize2(url string, channel chan Page) {
	fmt.Println("Getting ", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Send back a Page with both current url and size
	channel <- Page{URL: url, Size: len(body)}
}
