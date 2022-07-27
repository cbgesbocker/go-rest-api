package controllers

import (
	"example.com/m/src/models"
	"github.com/gin-gonic/gin"
)

func App() {
	models.ConnectDatabase()
	router := gin.Default()

	authRoutes := router.Group("/")
	authRoutes.Use(authMiddleware())
	authRoutes.GET("/books", FindBooks)
	authRoutes.POST("/books", CreateBook)
	authRoutes.GET("/books/:id", FindBook)
	authRoutes.PUT("/books/:id", UpdateBook)
	authRoutes.DELETE("/books/:id", DeleteBook)

	router.Run()

}
