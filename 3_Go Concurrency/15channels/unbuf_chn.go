package main

import "fmt"

// Unbuffered Channel (Default) :--> Both sides must be ready or the program blocks.
func greet(ch chan string) {
	msg := <-ch // Wait to receive
	fmt.Println("Received :- ", msg)
}

func main3() {
	ch := make(chan string) // Unbuffered
	go greet(ch)

	ch <- "message sent from main function" // Send blocks until greet() receives
}
