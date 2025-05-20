package main

import "fmt"

func main() {
	fmt.Println("-- Pointers --")
	// var ptr *int
	mynumber := 01
	var ptr = &mynumber
	fmt.Println("the values of pointers", ptr)
	fmt.Println("the values of pointers", *ptr)

	*ptr = *ptr * 2
	fmt.Println(mynumber)

}
