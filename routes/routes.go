package routes

import (
	"net/http"

	"gitlab.com/alura-courses-code/golang/web-crud/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.HandleIndex)
	http.HandleFunc("/new", controllers.HandleNew)
	http.HandleFunc("/delete", controllers.HandleDelete)
	http.HandleFunc("/edit", controllers.HandleEdit)
}
