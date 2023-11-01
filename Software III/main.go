package main

import (
	"ToolHub/handlers"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int
	Nombre     string
	correo     string
	Contraseña string
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "C:/Users/bklg1/OneDrive/Escritorio/Escritorio/Universidad/sotfware proyecto/AndresAristizabal_EdisonLargo_JuanCortes_MariaMinotta/Software III/base de datos/basededatos.db")

	db.Exec("CREATE TABLE IF NOT EXISTS usuario (id INTEGER PRIMARY KEY AUTOINCREMENT,nombre VARCHAR(30) NOT NULL,correo VARCHAR(30) NOT NULL,contraseña VARCHAR(30) NOT NULL )")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Bases de datos conectada")

	return db
}

func main() {

	db := dbConn()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/principal", handlers.PrincipalPage)
	r.GET("/register", handlers.RegisterPage)

	r.POST("/login", handlers.Login)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
