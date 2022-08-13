package main

import (
	"github.com/lsandi20/urlshortener-api/controllers"
)

func main() {
	router := controllers.AssignRoutes()
	router.Run("localhost:8080")
}
