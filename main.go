package main

import (
	"inventory-book-myskill/app"
	"inventory-book-myskill/auth"
	"inventory-book-myskill/db"
	"inventory-book-myskill/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	router.GET("/", auth.HomeHandler)

	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
