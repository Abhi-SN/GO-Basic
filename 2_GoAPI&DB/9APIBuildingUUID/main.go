package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Define the person structure
type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory slice to store users
var users []person

// Response message structure
type Message struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

func main() {
	fmt.Println("-----< Go API with UUID >-----")

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/UserPost", handlePostUser).Methods("POST")
	r.HandleFunc("/UserInfo", GetUserData).Methods("GET")
	r.HandleFunc("/UserPut/{id}", UpdateUserData).Methods("PUT")
	r.HandleFunc("/UserDelete/{id}", DeleteUser).Methods("DELETE")

	log.Println("🚀 Server running on http://localhost:3080")
	log.Fatal(http.ListenAndServe(":3080", r))
}

// POST: Create user
func handlePostUser(w http.ResponseWriter, r *http.Request) {
	var user person
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name and Age are required", http.StatusBadRequest)
		return
	}

	user.ID = uuid.New().String()
	users = append(users, user)

	response := Message{
		Status:   "Success",
		Messages: "User " + user.Name + " (Age: " + strconv.Itoa(user.Age) + ") created successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message
		ID string `json:"id"`
	}{
		Message: response,
		ID:      user.ID,
	})
}

// GET: Fetch all users
func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// PUT: Update user by ID
func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedUser person
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil || updatedUser.Age <= 0 || updatedUser.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	found := false
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := Message{
		Status:   "Success",
		Messages: "User updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DELETE: Delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	found := false
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := Message{
		Status:   "Success",
		Messages: "User deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
