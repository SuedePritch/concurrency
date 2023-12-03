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
	evilNinjas := []string{"Hattori", "Yoshi", "Kuma", "Momochi"}
	for _, ninja := range evilNinjas {
		// go keyword is used to create a goroutine
		// the attack function will be called in a goroutine
		// on its own it will fininsh before the goroutines are done
		go attack(ninja)
	}
	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Attacking", target)
}
