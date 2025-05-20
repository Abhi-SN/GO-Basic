package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type person struct {
	Name string
	Age  int
}

var users []person

type Message struct {
	Status   string `json:"status"`
	Messages string `json:"Messages"`
}

func main() {
	fmt.Println("-----< Post Request API Building >-----")
	r := mux.NewRouter()
	// r := mux.NewRouter()
	// r.HandleFunc("/user", handlePostUser).Methods("POST")
	r.HandleFunc("/UserPost", handlePostUser).Methods("POST")
	r.HandleFunc("/UserInfo", GetUserData).Methods("GET")
	r.HandleFunc("/UserPut/{name}", UpdateUserData).Methods("PUT")
	r.HandleFunc("/UserDel/{name}", HandleDeleteUser).Methods("DELETE")

	log.Println("🚀 Server running on http://localhost:3080")
	log.Fatal(http.ListenAndServe(":3080", r))
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	//Parse json
	var user person
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the fields
	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name and Age values are invalid", http.StatusBadRequest)
		return
	}
	// store post values in Memory
	users = append(users, user)

	// Build and send response
	response := Message{
		Status:   "Success",
		Messages: "The user name is  " + user.Name + " User Age is " + strconv.Itoa(user.Age),
	}
	w.Header().Set("content-type", "Application/json")
	json.NewEncoder(w).Encode(response)

}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	json.NewEncoder(w).Encode(users)

}

func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var updatedUser person
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil || updatedUser.Age <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// found := false

	var found bool = false
	for i, user := range users {
		if user.Name == name {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			found = true
			break

		}

	}
	if !found {
		http.Error(w, "Input not found", http.StatusNotFound)
		return
	}
	updateResponse := Message{
		Status:   "Success",
		Messages: "Updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateResponse)
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	// 1st Method
	// index := -1
	// for i, user := range users {
	// 	if user.Name == name {
	// 		index = i
	// 	}
	// }
	// if index == -1 {
	// 	http.Error(w, "User not found", http.StatusNotFound)
	// 	return
	// }
	// // Delete Users
	// users = append(users[:index], users[index+1:]...)

	// user by their Name via a DELETE request

	found := false
	for i, user := range users {
		if user.Name == name {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "User not found", http.StatusBadRequest)
		return

	}

	response := Message{
		Status:   "Success",
		Messages: "User" + name + " Deleted Successfully",
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}
