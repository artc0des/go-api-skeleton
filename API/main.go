package main

import (
	"github.com/gin-gonic/gin"
	"template.com/api/database"
	"template.com/api/routes"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
