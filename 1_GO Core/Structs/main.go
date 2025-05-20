package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
	Check bool
}

func main() {
	abhi := User{Name: "abhishek", Email: "abhi@.com", Age: 23, Check: true}
	fmt.Println("--- structs ---")
	fmt.Printf("Lets go with abhi %+v \n", abhi)
	fmt.Println("Lets go with Structs", abhi)
	fmt.Printf("My Name is %v and Email address is %v", abhi.Name, abhi.Email)

}
