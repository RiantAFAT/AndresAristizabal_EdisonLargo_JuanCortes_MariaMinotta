package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPage(c *gin.Context) {
	// Acceder a la sesión
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegistrarUsuario(c *gin.Context, db *sql.DB) {
	//fmt.Print("hola mundano")

	nombre := c.PostForm("nombre")
	correo := c.PostForm("correo")
	contraseña := c.PostForm("contraseña")

	// Realiza la inserción en la base de datos
	_, err := db.Exec("INSERT INTO Usuario (nombre, correo, contraseña) VALUES (?, ?, ?)", nombre, correo, contraseña)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//c.String(http.StatusOK, "Usuario registrado con éxito")
	fmt.Print("Usuario Registrado con exito")
	c.Redirect(http.StatusSeeOther, "/login")
}
