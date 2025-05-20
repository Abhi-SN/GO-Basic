package main

import "fmt"

func main4() {
	ch := make(chan string, 2) // the channel can hold up to 2 string values before a sender (via ch <- "value") blocks.
	ch <- "Message 1"
	ch <- "Message 2"
	// ch <- "Message 3"

	// add in another channel will block since buffer capcity will be full
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
