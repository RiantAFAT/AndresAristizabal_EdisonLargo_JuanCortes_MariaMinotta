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

	// Realiza el conteo en la base de datos
	consult := "SELECT COUNT(*) FROM Usuario WHERE correo= ?;"

	// Ejecuta la consulta
	var totalIguales int
	erro := db.QueryRow(consult, correo).Scan(&totalIguales)

	if erro != nil {
		totalIguales = 0
	}
	fmt.Print("correos iguales: ", totalIguales, " ")

	if totalIguales == 0 {
		_, err := db.Exec("INSERT INTO Usuario (nombre, correo, contraseña,Rol) VALUES (?, ?, ?,?)", nombre, correo, contraseña, "Registrado")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Print("usuario registrado")
		c.JSON(http.StatusOK, gin.H{"exito": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"exito": false})

	}
	//c.String(http.StatusOK, "Usuario registrado con éxito")

}
