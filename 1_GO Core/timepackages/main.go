package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome time study")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
	createDate := time.Date(2024, time.February, 01, 10, 1, 1, 2320, time.UTC)
	fmt.Println("Created Time is ", createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
