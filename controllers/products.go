package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gitlab.com/alura-courses-code/golang/web-crud/models"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	baseTemplate, err := template.ParseFiles(
		"templates/base.gohtml",
		"templates/index.gohtml",
	)

	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = baseTemplate.Execute(w, products)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		baseTemplate, err := template.ParseFiles(
			"templates/base.gohtml",
			"templates/new.gohtml",
		)

		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = baseTemplate.Execute(w, nil)
		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

func HandleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")

		newId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error parsing id:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		product, err := models.GetProductById(newId)
		if err != nil {
			log.Println("Error getting product:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		baseTemplate, err := template.ParseFiles(
			"templates/base.gohtml",
			"templates/edit.gohtml",
		)

		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = baseTemplate.Execute(w, product)
		if err != nil {
			log.Println("Error rendering template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		newId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error parsing id:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

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

		err = models.UpdateProduct(models.Product{Id: newId, Name: name, Description: description, Price: newPrice, Quantity: newQuantity})
		if err != nil {
			log.Println("Error updating product:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
