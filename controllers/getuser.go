package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"algocrux/config"
	"algocrux/models"
)

func GetUser(c *gin.Context) {
	username := c.Param("username")
	collection := config.DB.Collection("users")

	var user models.UserModel
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}