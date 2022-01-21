package main

import (
	"LibraryAPI/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", api.GetAllBooks)
	router.GET("/book/:id", api.GetBookById)
	router.POST("/book", api.PostBook)
	router.PUT("/book/:id", api.UpdateBookById)
	router.DELETE("/book/:id", api.DeleteBookById)

	router.Run("localhost:8080")
}
