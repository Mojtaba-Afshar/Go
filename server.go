package main

import (
	"github.com/Mojtaba-Afshar/ekart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.AppRoutes(server)
	server.GET("/", home)

	server.Run(":4000")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcom to gin",
	})
}
