package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sending Email confirmation")
	time.Sleep(1 * time.Second)
	fmt.Println("Email Sent")
}
func prepareShipment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Preparing Shipment")
	time.Sleep(2 * time.Second)
	fmt.Println("Shipment ready")

}

func ProcesPayment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing Payment")
	time.Sleep(1 * time.Second)
	fmt.Println("payment successful")
}
func generateInvoice(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Generating invoice...")
	time.Sleep(2 * time.Second)
	fmt.Println(" Invoice generated")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go prepareShipment(&wg)
	go sendEmail(&wg)
	go ProcesPayment(&wg)
	go generateInvoice(&wg)

	defer wg.Wait()
	fmt.Println("Order Placed")

}
