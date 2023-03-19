package main

import (
	"github.com/gin-gonic/gin"

	"go-rest-api/services"
)

func main() {
	r := gin.New()

	r.GET("/books", services.ListBooksHandler)
	r.POST("/books", services.CreateBookHandler)
	r.DELETE("/books/:id", services.DeleteBookHandler)
	r.GET("/users", services.ListALLUser)

	r.Run()

	services.LoadConfig()
}
