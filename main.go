package main

import (
	"log"
	"net/http"

	"gitlab.com/alura-courses-code/golang/web-crud/routes"
)

func main() {
	routes.LoadRoutes()

	serverAddr := "localhost:8080"
	log.Printf("Server is running at http://%s\n", serverAddr)

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
