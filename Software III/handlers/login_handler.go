package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	// Acceder a la sesión
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoguearUsuario(c *gin.Context, db *sql.DB) int {
	correo := c.PostForm("correo")
	contraseña := c.PostForm("contraseña")

	// Realiza el conteo en la base de datos
	consult := "SELECT id,contraseña FROM Usuario WHERE correo= ?;"

	// Ejecuta la consulta
	var idEncontrado int
	var contraseñaLog string
	idEncontrado = 0
	erro := db.QueryRow(consult, correo).Scan(&idEncontrado, &contraseñaLog)
	if erro != nil {
		idEncontrado = 0
	}
	//fmt.Print("idus: ", idEncontrado, "contra: ", contraseñaLog)

	if idEncontrado > 0 {
		if contraseña == contraseñaLog {
			c.JSON(http.StatusOK, gin.H{"estado": "exito", "mensaje": "Logueando..."})

		} else {
			c.JSON(http.StatusOK, gin.H{"estado": "fallido", "mensaje": "Contraseña incorrecta."})
			idEncontrado = 0
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"estado": "fallido", "mensaje": "No se encontro el correo."})
	}
	return idEncontrado
}

func LoguearInvitado(c *gin.Context, db *sql.DB) int {
	var idInvitado int
	// Realiza el conteo en la base de datos
	consult := "SELECT id FROM Usuario WHERE Rol='Invitado' LIMIT 1;"

	erro := db.QueryRow(consult).Scan(&idInvitado)

	if erro != nil {
		panic(erro)
	}
	c.JSON(http.StatusOK, gin.H{"estado": "exito", "mensaje": "Logueando como invitado..."})

	return idInvitado
}
