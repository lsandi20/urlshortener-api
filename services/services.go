package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lsandi20/urlshortener-api/models"
)

var links = []models.Link{
	{ID: "1", Original: "https://github.com/lsandi20", Short: "http://localhost:8080/lsandi", Clicks: 0},
}

func RedirectURL(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/lsandi20")
}

func GetLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}
