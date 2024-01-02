package routes

import (
	"net/http"

	"gitlab.com/alura-courses-code/golang/web-crud/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.HandleIndex)
}
