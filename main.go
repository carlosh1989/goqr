package main

import (
	"github.com/carlosh1989/goqr/controllers"
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models" // new
)

//"github.com/skip2/go-qrcode"

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/books", controllers.FindBooks)         //new
	r.POST("books", controllers.CreateBook)        //new
	r.GET("/books/:id", controllers.FindBook)      //new
	r.PATCH("/books/:id", controllers.UpdateBook)  //new
	r.DELETE("/books/:id", controllers.DeleteBook) //new
	r.Run("127.0.0.1:8000")
}
