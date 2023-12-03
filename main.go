package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed:", time.Since(start))
	}()

	// THIS IS A CHANNEL
	// channel := make(chan string)
	//cannot just use channel <- "First Message" without a go routine
	// this will create a deadlock as the channel cant hold information is just passes it through
	// channel <- "First Message"
	// go func() {
	// 	channel <- "First Message"
	// }()

	// THIS IS A BUFFERED CHANNEL
	// capacity is 1
	// channel := make(chan string, 1)
	// channel <- "First Message"

	channel := make(chan string, 2)
	channel <- "First Message"
	channel <- "Second Message"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
