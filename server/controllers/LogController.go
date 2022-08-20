package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"server/config"
	"server/models"
	"time"
)

func AddLog(c *gin.Context) {

	var log models.Log
	if err := c.BindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The Body is empty",
			"details": err,
		})

		return
	}

	log.Date = primitive.NewDateTimeFromTime(time.Now().Add(6 * time.Hour * -1))

	collection := config.Connection("log")
	insertResult, err := collection.InsertOne(context.TODO(), log)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The Log can not be created",
			"details": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Log has be create successfully",
		"result":  insertResult.InsertedID,
	})
}
