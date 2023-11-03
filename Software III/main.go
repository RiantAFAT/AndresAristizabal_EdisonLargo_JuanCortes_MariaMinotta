package main

import (
	"ToolHub/handlers"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int
	Nombre     string
	correo     string
	Contraseña string
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "/base_de_datos/ToolHub.db")

	db.Exec("CREATE TABLE IF NOT EXISTS Usuario (nombre VARCHAR(30) NOT NULL,correo VARCHAR(30) NOT NULL,contraseña VARCHAR(30) NOT NULL )")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Bases de datos conectada")

	return db

}
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

	//db := dbConn()
	db := OpenConnection()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/principal", handlers.PrincipalPage)
	r.GET("/register", handlers.RegisterPage)

	r.POST("/login", handlers.Login)
	r.POST("/registrarUsuario", func(c *gin.Context) {
		handlers.RegistrarUsuario(c, db)
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
