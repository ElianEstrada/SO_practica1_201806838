package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"server/config"
	"server/models"
)

func GetVehicles(c *gin.Context) {

	var vehicles []*models.Vehicle
	collection := config.Connection("vehicle")

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This Collection doesn't exist",
			"details": err,
		})
	}

	for cursor.Next(context.TODO()) {
		var vehicle models.Vehicle

		err := cursor.Decode(&vehicle)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "This model is not match with collection",
				"details": err,
			})
		}

		vehicles = append(vehicles, &vehicle)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This collection can not be required",
			"details": err,
		})
	}

	_ = cursor.Close(context.TODO())

	c.JSON(http.StatusOK, &vehicles)

}
