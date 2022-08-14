package services

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lsandi20/urlshortener-api/config"
	"github.com/lsandi20/urlshortener-api/database"
	"github.com/lsandi20/urlshortener-api/models"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RedirectToClient(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, config.Configs["CLIENT_URL"])
}

func RedirectURL(c *gin.Context) {
	short := c.Param("short")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.Db.Collection("links")

	var link models.Link

	err := coll.FindOne(ctx, bson.M{"short": short}).Decode(&link)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Error while getting link")
		return
	}

	c.Redirect(http.StatusMovedPermanently, link.Original)
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
		c.JSON(http.StatusBadRequest, "Error while getting links")
		return
	}

	c.JSON(http.StatusOK, urlShorternerResult)
}

func CreateLink(c *gin.Context) {
	var newLink models.Link

	if err := c.BindJSON(&newLink); err != nil {
		log.Println(err)
		return
	}

	short, err := shortid.Generate()

	if err != nil {
		panic(err)
	}

	newLink.Short = short
	newLink.Clicks = 0

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.Db.Collection("links")

	result, err := coll.InsertOne(ctx, newLink)

	if err != nil {
		panic(err)
	}

	newLink.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusOK, newLink)
}
