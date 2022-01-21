package api

import (
	"LibraryAPI/book"
	"errors"
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

	foundBook, err := GetBookHelper(newBook.ID)
	if err != nil {
		book.Books = append(book.Books, newBook)
		context.IndentedJSON(http.StatusCreated, newBook)
		return
	} else {
		context.IndentedJSON(400, "Book with that ID already exists: "+foundBook.Title)
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

	book, err := GetBookHelper(idNumber)
	if err == nil {
		context.IndentedJSON(http.StatusOK, book)
		return
	} else {
		context.IndentedJSON(404, "Book with that ID is not found")
		return
	}
}

//Helper function to loop trhough and find the book at that index
func GetBookHelper(id int) (book.Book, error) {
	for _, a := range book.Books {
		if a.ID == id {
			return a, nil
		}
	}
	return book.Book{}, errors.New("NOT_FOUND")
}

func DeleteBookById(context *gin.Context) {
	id := context.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(400, "Error getting id as a number")
		return
	}
	//Delete the book at the index
	book.Books = append(book.Books[:index-1], book.Books[index:]...)
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
