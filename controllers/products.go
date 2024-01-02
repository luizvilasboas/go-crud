package controllers

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.com/alura-courses-code/golang/web-crud/db"
	"gitlab.com/alura-courses-code/golang/web-crud/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

var database = db.ConnectDatabase()

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts(database)
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

	database.Close()
}
