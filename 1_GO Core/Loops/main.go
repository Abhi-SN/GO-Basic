package main

import "fmt"

func main() {
	// 1. Traditional for loop (like C/Java)
	fmt.Println("Traditional for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. For as a while loop (no init and post statement)
	fmt.Println("For as while loop:")
	j := 1
	for j <= 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	// 3. For with range to iterate over slice or map
	fmt.Println("For with range over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}

	fmt.Println("For with range over map:")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
