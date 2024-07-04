package main

import (
	database "supermarket-api/src/database"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database.ConnectDatabase()

	server.Run(":8000")
}
