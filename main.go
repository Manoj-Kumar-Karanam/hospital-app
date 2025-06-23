package main

import (
	"fmt"
	"hospitalapp/config"
	"hospitalapp/models"
	"hospitalapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})
	r := gin.Default()
	routes.SetRoutes(r)
	r.Run(":8080")
}