package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	Action string             `bson:action`
	Date   primitive.DateTime `bson:date`
}
