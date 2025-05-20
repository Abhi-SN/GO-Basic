package main

import "fmt"

func main() {

	fmt.Println("slices")

	var fruitList = []string{"apple", "banana", "peach", "pineapple", "grapes"}

	fruitList = append(fruitList, "watermelon")
	fmt.Println("type of fruits", fruitList)
	// fruitList = append(fruitList[1:3])
	fmt.Println("type of fruits", fruitList)
	var index int = 2

	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)

}
