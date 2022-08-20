package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		return
	}

	for cursor.Next(context.TODO()) {
		var vehicle models.Vehicle

		err := cursor.Decode(&vehicle)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "This model is not match with collection",
				"details": err,
			})
			return
		}

		vehicles = append(vehicles, &vehicle)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This collection can not be required",
			"details": err,
		})

		return
	}

	_ = cursor.Close(context.TODO())

	c.JSON(http.StatusOK, &vehicles)
}

func GetVehicle(c *gin.Context) {

	collection := config.Connection("vehicle")

	var vehicle models.Vehicle

	id := c.Params.ByName("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This id does not exist on DataBase",
			"details": err,
		})
		return
	}

	err = collection.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&vehicle)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This Model dont match with data",
			"details": err,
		})

		return
	}

	if vehicle.Plate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can not find this vehicle",
		})

		return
	}

	c.JSON(http.StatusOK, &vehicle)
}

func AddVehicle(c *gin.Context) {

	var vehicle models.Vehicle

	if err := c.BindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This Body is empty",
			"details": err,
		})

		return
	}

	vehicle.ID = primitive.NewObjectID()

	collection := config.Connection("vehicle")
	insertResult, err := collection.InsertOne(context.TODO(), vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can not insert into Data Base",
			"details": err,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Vehicle create successfully",
		"result":  insertResult.InsertedID,
	})
}


