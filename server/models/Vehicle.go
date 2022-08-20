package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	ID     primitive.ObjectID `bson:"_id"`
	Plate  string             `bson:"plate"`
	Make   string             `bson:"make"`
	Model  int64              `bson:model`
	Series string             `bson:series`
	Color  string             `bson:color`
}
