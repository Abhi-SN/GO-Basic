package main

import "fmt"

func sendData(ch chan int) {
	ch <- 10
	ch <- 21
	ch <- 21
	ch <- 232
	close(ch)
}
func receive(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:- ", value)
	}
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	receive(ch)
}
