package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		products := []Product{
			{Name: "Camiseta", Description: "Oakley", Price: 32, Quantity: 10},
			{Name: "Calça", Description: "Jeans", Price: 100, Quantity: 5},
			{Name: "Tênis", Description: "Nike", Price: 200, Quantity: 2},
			{Name: "Meia", Description: "Nike", Price: 10, Quantity: 100},
		}

		temp.ExecuteTemplate(w, "Index", products)
	})

	fmt.Println(`Server is running in port http://localhost:8080`)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
