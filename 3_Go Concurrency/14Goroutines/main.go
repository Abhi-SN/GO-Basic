package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchUserData() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println(" User data fetched")
}

func fetchOrderHistory() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println(" Order history fetched")
}

func fetchProductList() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Product list fetched")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Start fetching concurrently
	go fetchUserData()
	go fetchOrderHistory()
	go fetchProductList()

	// Wait for all goroutines to complete (temporary way)
	time.Sleep(2 * time.Second)

	fmt.Println(" All API data fetched")
}

// import (
// 	"fmt"
// 	"time"
// )

// func printMessage(msg string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(msg, i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go printMessage("Goroutine 1")
// 	go printMessage("Goroutine 2")
// 	go printMessage("Goroutine 3")

// 	printMessage("Main Function")
// }
