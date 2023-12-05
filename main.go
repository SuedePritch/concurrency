package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed:", time.Since(start))
	}()
	var evilNinjas = []string{"James", "Jenn", "John", "Jim"}
	var attackers = []string{"Attacker1", "Attacker2", "Attacker3"}
	var wg sync.WaitGroup

	for _, attacker := range attackers {
		wg.Add(1)
		go func(attacker string) {
			defer wg.Done()
			for _, evilNinja := range evilNinjas {
				wg.Add(1)
				go attack(attacker, evilNinja, &wg)
			}
		}(attacker)
	}

	wg.Wait()
}

func attack(attacker string, evilNinja string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s is attacking: %s\n", attacker, evilNinja)
}
