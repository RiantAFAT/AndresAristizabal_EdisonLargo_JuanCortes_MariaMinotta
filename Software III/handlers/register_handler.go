package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPage(c *gin.Context) {
	// Acceder a la sesión
	c.HTML(http.StatusOK, "register.html", nil)
}
