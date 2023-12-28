package handler

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()

	book := router.Group("/book")
	{
		book.GET("/", GetBooks)
		book.GET("/:id", GetBookByID)
	}

	category := router.Group("/category")
	{
		category.GET("/", GetCategories)
	}

	author := router.Group("/author")
	{
		author.GET("/", GetAuthors)
	}

	router.Run("localhost:8080") //
}
