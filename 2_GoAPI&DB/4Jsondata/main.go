package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Course Name"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json Data")
	encodeJson()

	decodeJsondata()

}

func encodeJson() {
	DetailedCourses := []course{
		{"Golang", 100, "YouTube", "Abcd1234", []string{"Web-dev", "Backend"}},
		{"Python", 00, "YouTube", "Aba234", []string{"Web-dev", "Backend"}},
		{"Java", 2200, "YouTube", "Abcd1d", nil},
	}
	// Package this data as json data
	finalJson, err := json.MarshalIndent(DetailedCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJsondata() {
	jsonDatafromData := []byte(`
	{
		"Course Name": "Golang",
		"Price": 100,
		"Platform": "YouTube",
		"tags": ["Web-dev","Backend"]
	}`)
	var decodeCourse course

	checkvaildate := json.Valid(jsonDatafromData)
	if checkvaildate {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDatafromData, &decodeCourse)
		fmt.Printf("The decode json data is %#v\n", decodeCourse)
	} else {
		fmt.Println("json was not valid")
	}

	var jsondecodeData map[string]interface{}
	json.Unmarshal(jsonDatafromData, &jsondecodeData)
	for k, v := range jsondecodeData {
		fmt.Printf("the key is %v and the value is %v and type is %T\n", k, v, v)
	}
}
