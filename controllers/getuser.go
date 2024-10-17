package controllers

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "algocrux/config"
    "algocrux/models"
    "algocrux/utils"
)

func GetUser(c *gin.Context) {
    tokenString, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in cookies"})
        return
    }

    token, claims, err := utils.ValidateJWT(tokenString)
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    githubUsername := claims.GithubUsername

    collection := config.DB.Collection("users")
    var user models.UserModel
    err = collection.FindOne(context.Background(), bson.M{"github_username": githubUsername}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}