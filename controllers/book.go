package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nilerajput91/assig5bookapi/apihelpers"
	"github.com/nilerajput91/assig5bookapi/models"
)

// ListBook of all books
func ListBook(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBook(&book)

	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// AddNewBook from book
func AddNewBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.AddNewBook(&book)
	if err != nil {
		apihelpers.RespondJSON(c, 400, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// GetOneBook from book
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	err := models.GetOneBook(&book, id)

	if err != nil {
		apihelpers.RespondJSON(c, 400, book)

	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// PutOneBook on the book
func PutOneBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.GetOneBook(&book, id)

	if err != nil {
		apihelpers.RespondJSON(c, 400, book)
	}
	c.BindJSON(&book)
	err = models.PutOneBook(&book, id)

	if err != nil {
		apihelpers.RespondJSON(c, 400, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// DeleteBook function
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(&book, id)
	if err != nil {
		apihelpers.RespondJSON(c, 400, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}
