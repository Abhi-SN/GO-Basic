package main

import (
	"fmt"
	"sync"
	"time"
)

var tickets = 10
var mu sync.Mutex

func bookTickets(user string, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	if tickets > 0 {
		fmt.Printf("%s booked a ticket.  Tickets left: %d\n", user, tickets-1)
		tickets--
		time.Sleep(500 * time.Millisecond) // Simulate processing time
	} else {
		fmt.Printf("%s tried to book a ticket, but none were left. \n", user)
	}
}

func main() {
	var wg sync.WaitGroup
	users := []string{"Abhi", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi", "Ivan", "Judy", "Ken", "Laura"}

	for _, user := range users {
		wg.Add(1)
		go bookTickets(user, &wg)
	}

	wg.Wait()
	fmt.Println("All booking attempts finished!")
}
