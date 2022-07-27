package main

import (
	"example.com/m/api"
	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	router := gin.Default()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.Group("/").Use(api.AuthMiddleware)

	router.Run()
}
