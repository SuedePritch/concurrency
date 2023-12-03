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
	//channel select statement
	//the select statement lets a goroutine wait on multiple communication operations.
	//A select blocks until one of its cases can run, then it executes that case.
	//It chooses one at random if multiple are ready.
	//The default case in a select is run if no other case is ready.
	ninja1, ninja2 := make(chan string), make(chan string)
	go captainElect(ninja1, "Ninja1")
	go captainElect(ninja2, "Ninja2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	default:
		fmt.Println("No one is the captain")
	}
	roughlyFair()
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}
func roughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int
	for i := 0; i < 1000; i++ {
		select {
		case <-ninja1:
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		}
	}
	fmt.Printf("ninja1 %d vs %d ninja2\n", ninja1Count, ninja2Count)
}
