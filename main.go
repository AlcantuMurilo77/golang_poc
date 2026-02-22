package main

import (
	"example.com/my_go_proj/db"
	"example.com/my_go_proj/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic("Could not start server")
	}
}
