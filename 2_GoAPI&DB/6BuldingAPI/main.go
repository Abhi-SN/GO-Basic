package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
}

func main() {
	fmt.Println("---- < Learning API Build with Get request >-----")
	r := mux.NewRouter()
	r.HandleFunc("/hello", HandlerGet).Methods("GET")
	r.HandleFunc("/hello/{name}", HandleGetPathVariable).Methods("GET")
	r.HandleFunc("/helo/{name}", HandleGetQueryParameter).Methods("GET")
	// http.HandleFunc("/hello", HandlerGet)
	// Start the server
	log.Println("🚀 Server is running on http://localhost:3080")
	log.Fatal(http.ListenAndServe(":3080", r))

}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	response := Message{
		Status:   "Success",
		Messages: "Hello welcome to Get request Output",
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleGetPathVariable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	response := Message{
		Status:   "Success",
		Messages: "Hello Welcome to Get request with A path variable output " + name + "!",
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleGetQueryParameter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	// Get query parameters
	title := r.URL.Query().Get("title")
	lang := r.URL.Query().Get("ln")
	// Set default values if query params are empty
	if title == "" {
		title = ""
	}
	if lang == "" {
		lang = "en"
	}
	// Building responsive output
	response := Message{
		Status:   "Success",
		Messages: "The Name is  " + name + "  title is  " + title + "  and language " + lang,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
