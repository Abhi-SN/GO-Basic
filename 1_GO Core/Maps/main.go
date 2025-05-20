package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Maps ---")

	languages := make(map[string]string)
	languages["Py"] = "python"
	languages["Rb"] = "Ruby"
	languages["js"] = "JavaScript"

	fmt.Println("List of Languages", languages)
	fmt.Println("Language of Py is", languages["Py"])
	delete(languages, "Rb")
	fmt.Println("Deleted RB", languages)

	for key, values := range languages {
		fmt.Printf("the key pair is %v and the value is %v \n", key, values)
	}

}
