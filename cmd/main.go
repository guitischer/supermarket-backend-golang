package main

import (
	"supermarket-api/src/controllers"
	database "supermarket-api/src/database"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database.ConnectDatabase()

	server.GET("/product-types", controllers.GetAllProductTypes)
	server.POST("/product-types", controllers.CreateProductType)

	server.Run(":8000")
}
