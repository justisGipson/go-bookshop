package main

import (
	"github.com/justisGipson/go-bookshop/controllers"
	"github.com/justisGipson/go-bookshop/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize gin
	r := gin.Default()

	// connect db
	models.ConnectDatabase()

	// routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// run server
	r.Run()
}
