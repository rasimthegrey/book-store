package handler

import (
	"book-store/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBooks(c *gin.Context) {
	books, err := db.QueryAllBooks()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
	}
	c.IndentedJSON(200, books)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := db.QueryBookByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Bad request",
		})
	}

	c.JSON(200, gin.H{
		"Name":     book.BookName,
		"ISBN":     book.ISBN,
		"Category": book.BookCategory.CategoryName,
		"Author":   book.BookAuthor.AuthorName,
	})
}

func GetCategories(c *gin.Context) {
	categories, err := db.QueryCategories()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
	}
	c.IndentedJSON(200, categories)
}

func GetAuthors(c *gin.Context) {
	authors, err := db.QueryAuthors()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
	}
	c.IndentedJSON(200, authors)
}
