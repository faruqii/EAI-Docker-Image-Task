package main

import (
	"github.com/faruqii/EAI-Docker-Image-Task/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve the form
	r.GET("/", controller.HandleForm)

	// Handle the form submission
	r.POST("/", controller.Home)

	// Serve static files
	r.Static("/static", "./assets")

	// Start the server
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
