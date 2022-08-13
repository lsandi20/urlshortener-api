package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lsandi20/urlshortener-api/services"
)

func AssignRoutes() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/links", services.GetLinks)
	router.GET("/:short", services.RedirectURL)
	return
}
