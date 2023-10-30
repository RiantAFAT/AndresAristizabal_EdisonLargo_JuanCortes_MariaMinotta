package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrincipalPage(c *gin.Context) {
	// Acceder a la sesión
	c.HTML(http.StatusOK, "principal.html", nil)
}
