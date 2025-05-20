package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	fmt.Println("rate our pizza")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numrating, err := strconv.ParseInt(strings.TrimSpace(input), 32, 64)
	if err != nil {
		fmt.Print("Error occured", err)
	} else {
		fmt.Println("number rating is ", numrating+1)
	}

}
