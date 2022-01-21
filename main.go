package main

import (
	"LibraryAPI/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", api.GetAllBooks)
	router.GET("/books/:id", api.GetBookById)
	router.POST("/books", api.PostBook)
	router.PUT("/books/:id", api.UpdateBookById)
	router.DELETE("/books/:id", api.DeleteBookById)

	router.Run("localhost:8080")
}
