package services

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lsandi20/urlshortener-api/database"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
)

var short, err = shortid.Generate()

func RedirectURL(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/lsandi20")
}

func GetLinks(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.Db.Collection("links")

	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	var urlShorternerResult []bson.M

	if err = cursor.All(ctx, &urlShorternerResult); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, urlShorternerResult)
}
