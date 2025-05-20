package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("---> Post Requests <---")
	const my_url = "http://localhost:3000/post"
	PerformPostFormRequest(my_url)
	handlePostReq(my_url)

}
func handlePostReq(myurl string) {
	requestJson := strings.NewReader(`
	{
	"country" : "india",
	"id" : 90,
	"hostname":"http://localhost:3000/"
	}
	`)

	resp, err := http.Post(myurl, "application/json", requestJson)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func PerformPostFormRequest(myurl string) {
	data := url.Values{}
	data.Add("first", "Abhishek")
	data.Add("Id", "43")
	respo, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	dataByte, _ := ioutil.ReadAll(respo.Body)
	fmt.Println(string(dataByte))
	fmt.Println("---> Postform <---")

}
