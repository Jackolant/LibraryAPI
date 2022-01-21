package main

import (
	"fmt"

	"LibraryAPI/api"
	"LibraryAPI/book"
	"LibraryAPI/data"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello There")
	book.Print()
	api.Print()
	data.Print()
}

func run() {
	router := gin.Default()
	router.Run("localhost:8080")
}
