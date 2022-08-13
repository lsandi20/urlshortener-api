package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type link struct {
	ID       string `json:"id"`
	Original string `json:"original"`
	Short    string `json:"short"`
	Clicks   int64  `json:"clicks"`
}

var links = []link{
	{ID: "1", Original: "https://github.com/lsandi20", Short: "http://localhost:8080/lsandi", Clicks: 0},
}

func redirectURL(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/lsandi20")
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func main() {
	router := gin.Default()
	router.GET("/links", getLinks)
	router.GET("/:short", redirectURL)

	router.Run("localhost:8080")
}
