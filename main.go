package main

import (
	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
