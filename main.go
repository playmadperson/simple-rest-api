package main

import (
	"github.com/gin-gonic/gin"
	"github.com/playmadperson/simple-rest-api/controllers"
	"github.com/playmadperson/simple-rest-api/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateNewBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}
