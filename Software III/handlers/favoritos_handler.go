package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

var listaHerramientasUser []int

func FavoritosPage(c *gin.Context, db *sql.DB, idUser int) {

	herramientasFav = []HerramientaFav{}
	herramientasTotal = []HerramientaFav{}
	listaHerramientasUser = []int{}

	if len(herramientasTotal) == 0 {
		// Realiza la consulta a la base de datos
		rows, err := db.Query("SELECT * FROM Herramienta")
		if err != nil {

			fmt.Print("tabla vacia")
		}
		defer rows.Close()

		for rows.Next() {
			var r HerramientaFav
			err := rows.Scan(&r.IdH, &r.URLImagen, &r.Descripcion, &r.Categoria, &r.Estado, &r.Licencia, &r.EnlaceOficial, &r.Nombre)
			if err != nil {
				fmt.Println("Error al escanear herramientas:", err)
			}
			herramientasTotal = append(herramientasTotal, r)
		}
		defer rows.Close()
	}

	// Realiza la consulta a la base de datos
	query := "SELECT herramienta_id FROM Favorito WHERE usuario_id = ?"
	rowss, err := db.Query(query, idUser)
	if err != nil {

		fmt.Print("tabla vacia")
		defer rowss.Close()
	}
	defer rowss.Close()

	for rowss.Next() {
		var idHe int
		var esta bool
		esta = true

		err := rowss.Scan(&idHe)
		if err != nil {
			fmt.Println("Error al escanear herramientas:", err)
		}

		if len(listaHerramientasUser) == 0 {
			listaHerramientasUser = append(listaHerramientasUser, idHe)
		} else {
			for _, herrUser := range listaHerramientasUser {
				if herrUser != idHe {
					esta = false

				} else {
					esta = true
					break
				}
			}
		}

		if esta == false {
			listaHerramientasUser = append(listaHerramientasUser, idHe)
		}

	}

	/*fmt.Println("Elementos de la lista:")
	for _, elemento := range listaHerramientasUser {
		fmt.Println(elemento)
	}
	fmt.Println("Elementos de la lista:")
	for _, elemento := range herramientasTotal {
		fmt.Println(elemento)
	}*/

	for _, idHerr := range listaHerramientasUser {
		for _, herr := range herramientasTotal {
			if idHerr == herr.IdH {
				herramientasFav = append(herramientasFav, herr)
			} else {

			}
		}

	}

	t, err := template.ParseFiles("./templates/favoritos.html")
	if err != nil {
		fmt.Println("Error al parsear:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	// Renderizar la plantilla con la lista de herramientas
	err = t.Execute(c.Writer, herramientasFav)
	if err != nil {
		fmt.Println("Error al renderizar:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}
}

var herramientasFav []HerramientaFav
var herramientasTotal []HerramientaFav

type HerramientaFav struct {
	IdH           int
	URLImagen     string
	Descripcion   string
	Categoria     string
	Estado        string
	Licencia      string
	EnlaceOficial string
	Nombre        string
	// Agrega otros campos relacionados con el producto según sea necesario
}
