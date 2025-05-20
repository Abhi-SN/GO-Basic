package main

import "fmt"

// func sayHello(ch chan string) {
// 	ch <- "Hello from goroutine"
// }

// func main() {
// 	ch := make(chan string) // Create a string channel

// 	go sayHello(ch) // Start goroutine

// 	message := <-ch // Wait to receive message
// 	fmt.Println("Received:", message)
// }

func sendValueTo_CH(ch chan int) {
	for i := 0; i <= 3; i++ {
		ch <- i
	}
	close(ch)
	// close(ch) signals that no more values will be sent.

}

func main1() {
	ch := make(chan int)
	go sendValueTo_CH(ch)
	for val := range ch {
		fmt.Println("Channel value is ", val)
		//Use range with channels to receive until the channel is closed.
	}
	fmt.Println("All Channel values received")
}
