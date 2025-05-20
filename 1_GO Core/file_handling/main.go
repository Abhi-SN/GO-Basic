package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(">-------welcome to go lang files -------> ")
	content := "Data to entered in the files --> lets go"
	file, err := os.Create("./lets_go.txt")
	checkErrorNil(err)
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()

	readfile("./lets_go.txt")

}

func readfile(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(b))
}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
