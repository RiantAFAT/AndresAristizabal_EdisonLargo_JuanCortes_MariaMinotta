package main

import (
	"ToolHub/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/principal", handlers.PrincipalPage)

	r.POST("/login", handlers.Login)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
