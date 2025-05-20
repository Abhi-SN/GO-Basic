package main

import (
	"fmt"
	"log"
	"net/http"

	"modular-api/config"
	"modular-api/routes"
)

func main() {
	fmt.Println("Starting Modular Go API...")
	config.ConnectDB() // Establish DB connection

	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":3080", router))
}
