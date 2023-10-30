package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	// Acceder a la sesión
	c.HTML(http.StatusOK, "login.html", nil)
}
func Login(c *gin.Context) {
	fmt.Print("hola mundano")
}
