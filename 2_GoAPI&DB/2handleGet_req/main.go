package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const myurl = "http://localhost:3000/message"
	handleGetRequest(myurl)
}

func handleGetRequest(my_url string) {
	response, err := http.Get(my_url)
	if err != nil {
		panic(err)
	}
	fmt.Println("the status code is", response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("the length of content is =", response.ContentLength)
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		var responseData strings.Builder
		byteCount, _ := responseData.Write(data)
		fmt.Println("the length using stringsBuilder is", byteCount)
		fmt.Println(responseData.String())
		// dataContent := string(data)
		// fmt.Println("the data is ", dataContent)

	}

}
