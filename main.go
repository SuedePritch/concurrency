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
	//iteratate through the channel
	//this example is a ninja throwing stars at a target
	numStarsToThrow := 500
	channel := make(chan string, numStarsToThrow)
	go throwNinjaStars(channel, numStarsToThrow)
	for i := 0; i < numStarsToThrow; i++ {
		fmt.Println(<-channel)
	}

}

func throwNinjaStars(channel chan string, numStarsToThrow int) {
	for i := 0; i < numStarsToThrow; i++ {
		//generate a random score between 1 and 10
		score := 1 + rand.Intn(10)
		channel <- fmt.Sprintf("You scored %d points!", score)
	}
}
