package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Original string             `bson:"original,omitempty" json:"original"`
	Short    string             `bson:"short,omitempty" json:"short"`
	Clicks   int64              `bson:"clicks" json:"clicks"`
}
