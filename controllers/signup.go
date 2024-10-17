package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
    "algocrux/config"
    "algocrux/models"
)

var validate = validator.New()

type SignupRequest struct {
    Name           string `json:"name" binding:"required,min=3,max=50"`
    Email          string `json:"email" binding:"required,email"`
    Password       string `json:"password" binding:"required,min=6,max=100"`
    GithubUsername string `json:"github_username" binding:"required,min=3,max=20"`
    ProfileUrl     string `json:"profile_url"`
}

func Signup(c *gin.Context) {
    var req SignupRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Unable to hash password",
        })
        return
    }

    collection := config.DB.Collection("users")
    var existingUser models.UserModel
    err = collection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&existingUser)

    if err == nil {
        c.JSON(http.StatusConflict, gin.H{
            "error": "User already exists with this email",
        })
        return
    }

    profileUrl := "https://github.com/" + req.GithubUsername

    user := models.UserModel{
        ID:             primitive.NewObjectID(),
        Name:           req.Name,
        Email:          req.Email,
        Password:       string(passwordHash),
        GithubUsername: req.GithubUsername,
        ProfileUrl:     profileUrl,
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
    }

    _, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error during signup",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User signup successful",
    })
}