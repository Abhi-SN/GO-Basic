package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("---- Web request ---")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("google response", response)
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)

	fmt.Println("lets go", content)

}

// fmt.Println("Welcome to handling the URL's in golang")
// 	fmt.Println(myurl)
// 	result, error := url.Parse(myurl)
// 	if error != nil {
// 		panic(error)
// 	}
// 	fmt.Println(result.Scheme)
// 	fmt.Println(result.Hostname())
// 	fmt.Println(result.Path)
// 	fmt.Println(result.Port())
// 	fmt.Println(result.RawQuery)

// 	qparams := result.Query()
// 	fmt.Printf("the type of qparams is %T \n", qparams)
// 	fmt.Println(qparams)
// 	fmt.Println(qparams["learn"])

// 	PartsOfUrl := &url.URL{
// 		Scheme:  "https",
// 		Host:    "www.w3schools.com",
// 		Path:    "/golang/Golang_example.asp",
// 		RawPath: "topic=Maps",
// 	}
// 	anotherUrl := PartsOfUrl.String()
// 	fmt.Println(anotherUrl)
