package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed:", time.Since(start))
	}()
	channel := make(chan string)
	go throwNinjaStars(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}

}

func throwNinjaStars(channel chan string) {
	numStarsToThrow := 20
	for i := 0; i < numStarsToThrow; i++ {
		//generate a random score between 1 and 10
		score := 1 + rand.Intn(10)
		channel <- fmt.Sprintf("You scored %d points!", score)
	}
	//this allows the for loop in main() to exit
	close(channel)
}
