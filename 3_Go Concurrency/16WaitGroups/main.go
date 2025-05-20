package main

import (
	"fmt"
	"sync"
	"time"
)

func cleanRoom(room string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker starting to clean the room ", room)
	time.Sleep(2 * time.Second) // time taken compelete cleaning
	fmt.Println("Finished cleaning ", room)
}

func main1() {
	var wg sync.WaitGroup
	rooms := []string{"Living room", "dinning room", "Bedroom"}
	for _, val := range rooms {
		wg.Add(1)
		go cleanRoom(val, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All rooms cleaned. Locking the house!")
}
