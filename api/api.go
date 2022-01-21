package api

import (
	"LibraryAPI/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, book.Books)
}

func PostBook(context *gin.Context) {
	var newBook book.Book

	if err := context.BindJSON(&newBook); err != nil {
		context.IndentedJSON(400, "Error inputting a book")
		return
	}

	book.Books = append(book.Books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookById(context *gin.Context) {
	id := context.Param("id")
	idNumber, convError := strconv.Atoi(id)
	if convError != nil {
		context.IndentedJSON(400, "Error getting id as a number")
		return
	}

	for _, a := range book.Books {
		if a.ID == idNumber {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func DeleteBookById(context *gin.Context) {
	id := context.Param("id")
	index, err := strconv.Atoi(id)
	index = index - 1
	book.Books = append(book.Books[:index], book.Books[index+1:]...)
	context.IndentedJSON(http.StatusOK, err)
	return
}

func UpdateBookById(context *gin.Context) {
	//get the id as a param to be updated
	id := context.Param("id")

	//convert to an int
	idNumber, convError := strconv.Atoi(id)
	//if the input is not a valid int
	if convError != nil {
		context.IndentedJSON(400, "Error getting id as a number")
		return
	}

	var updateBook book.Book

	//Convert body to json, if an error return 400
	if err := context.BindJSON(&updateBook); err != nil {
		context.IndentedJSON(400, "Error getting parames to update a book")
		return
	}

	//loop through and find the item to be updated
	for i, x := range book.Books {
		if x.ID == idNumber {
			book.Books[i].Title = updateBook.Title
			book.Books[i].Author = updateBook.Author
			book.Books[i].Pages = updateBook.Pages
			book.Books[i].Shelf = updateBook.Shelf
			book.Books[i].Genre = updateBook.Genre

			context.IndentedJSON(http.StatusOK, book.Books[i])
			return
		}
	}
}
