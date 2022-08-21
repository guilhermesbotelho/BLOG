package main

import (
	"net/http"

	"github.com/guilhermesbotelho/BLOG/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
