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
	examplechannel := make(chan bool)
	evilNinjas := []string{"Hattori", "Yoshi", "Kuma", "Momochi"}
	for _, ninja := range evilNinjas {
		// go keyword is used to create a goroutine
		// the attack function will be called in a goroutine
		// on its own main func will finish before the goroutines are done
		go attack(ninja, examplechannel)

		// we can use the <- operator to receive a value from a channel
		// <- is to read from a channel i.e. pulling it out of the channel
		fmt.Println(<-examplechannel)
	}
	time.Sleep(time.Second)
}

func attack(target string, attacked chan bool) {
	fmt.Println("Attacking", target)
	//when we want to add a value to a channel we use the <- operator pointing to the channel
	// <- is to write to a channel i.e. pushing it into the channel
	attacked <- true
}
