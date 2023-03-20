package main

import (
	"fmt"

	"example.com/go-project/controllers"
	"example.com/go-project/database"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"
)

func main() {
	fmt.Println("starting aplication ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books/:id", controllers.ReadBook)
	r.GET("/books", controllers.ReadBooks)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run(":5000")
}
