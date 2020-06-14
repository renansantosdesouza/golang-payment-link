package main

import (
	"net/http"

	"github.com/payment-link/config"
	"github.com/payment-link/routes"
)

func main() {
	config.Initialize()
	routes.BindRoutes()
	http.ListenAndServe(":8000", nil)
}
