package main

import (
	"os"

	"algocrux/config"
	"algocrux/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDatabase()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000","https://algocruxx.vercel.app"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from AlgoCrux",
		})
	})

	server.POST("/api/user/signup", controllers.Signup)
	server.POST("/api/user/login", controllers.Login)
	server.GET("/api/user/get-user", controllers.GetUser)
	server.GET("/api/user/developers", controllers.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}