package controllers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lsandi20/urlshortener-api/services"
)

func AssignRoutes() (router *gin.Engine) {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "access-control-allow-origin", "acess-control-allow-headers", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	router.GET("/links", services.GetLinks)
	router.POST("/link", services.CreateLink)
	router.GET("/:short", services.RedirectURL)
	return
}
