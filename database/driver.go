package database

import (
	"context"
	"log"

	"github.com/lsandi20/urlshortener-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbClient *mongo.Client
var Db *mongo.Database

func InitDatabase() {
	uri := config.Configs["MONGODB_URI"]
	dbName := config.Configs["MONGODB_DATABASE"]

	if uri == "" {
		log.Fatal("'MONGODB_URI' not set in .env")
	}

	var err error
	DbClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	Db = DbClient.Database(dbName)

}
