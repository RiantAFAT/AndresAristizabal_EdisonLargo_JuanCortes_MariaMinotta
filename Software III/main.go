package main

import (
	"ToolHub/handlers"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (db *sql.DB) {
	// Configuración de la conexión
	dsn := "root:root@tcp(localhost:3306)/ToolHub"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Manejar el error
		log.Println("openerror")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		// Manejar el error
		log.Println("error ping")
	}
	//db.Exec("INSERT INTO Usuario (nombre, correo, contraseña) VALUES (?, ?, ?)", "Hernan", "correo@example.com", "root")
	return db
}

func main() {
	var usuarioLog int

	db := OpenConnection()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/principal", handlers.PrincipalPage)
	r.GET("/register", handlers.RegisterPage)
	r.GET("/detalles/:nombre", handlers.DetallesPage)

	r.POST("/registrarUsuario", func(c *gin.Context) {
		handlers.RegistrarUsuario(c, db)
	})
	r.POST("/loguearUsuario", func(c *gin.Context) {
		usuarioLog = handlers.LoguearUsuario(c, db)
		fmt.Print("usuario logueado: ", usuarioLog, " ")
	})
	r.POST("/loguearInvitado", func(c *gin.Context) {
		usuarioLog = handlers.LoguearInvitado(c, db)
		fmt.Print("usuario logueado: ", usuarioLog, " ")
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
