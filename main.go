package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=luiz password=root dbname=alura_loja host=localhost sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := ConnectDatabase()

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawProducts, err := db.Query("SELECT name, description, price, quantity FROM products")

		if err != nil {
			panic(err)
		}

		products := []Product{}

		for rawProducts.Next() {
			var name, description string
			var price float64
			var quantity int

			err = rawProducts.Scan(&name, &description, &price, &quantity)
			if err != nil {
				panic(err)
			}

			products = append(products, Product{name, description, price, quantity})
		}

		temp.ExecuteTemplate(w, "Index", products)
	})

	fmt.Println(`Server is running in port http://localhost:8080`)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
