package main

import "fmt"

func sqauredNum(ch chan int) {
	for num := range ch {
		sqaure := num * num
		fmt.Println("square :- ", sqaure)
	}
}
func main2() {
	ch := make(chan int)
	go sqauredNum(ch)
	ch <- 2
	ch <- 4
	ch <- 6
	close(ch)

	fmt.Scanln() // Wait for Enter to exit program
	fmt.Println("end")
}
