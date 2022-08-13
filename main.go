package main

import (
	"context"
	"time"

	"github.com/lsandi20/urlshortener-api/config"
	"github.com/lsandi20/urlshortener-api/controllers"
	"github.com/lsandi20/urlshortener-api/database"
)

func main() {
	config.ReadConfig()

	database.InitDatabase()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err := database.DbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	port := config.Configs["PORT"]

	router := controllers.AssignRoutes()
	router.Run(":" + port)
}
