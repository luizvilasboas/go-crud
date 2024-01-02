package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gitlab.com/alura-courses-code/golang/web-crud/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "Index", products)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "New", nil)

		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		newPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error parsing price:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		newQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error parsing quantity:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		err = models.InsertProduct(models.Product{Name: name, Description: description, Price: newPrice, Quantity: newQuantity})
		if err != nil {
			log.Println("Error inserting product:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error parsing id:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = models.DeleteProductById(newId)
	if err != nil {
		log.Println("Error deleting product:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
