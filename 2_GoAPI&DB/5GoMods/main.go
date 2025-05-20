package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Lets understand Mods, Mux , Tidy etc")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("Get")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func Greeter() {
	fmt.Println("Welcome to GO lang video")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Lets Go Details using Mux <h1>"))
}
