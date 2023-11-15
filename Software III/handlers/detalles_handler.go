package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	// Importa tu paquete de base de datos
)

var herramientaDetallada HerramientaDetall

func DetallesPage(c *gin.Context, db *sql.DB) {
	nombre := c.Param("nombre")

	consult := "SELECT id, urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre FROM Herramienta WHERE nombre = ?"

	erro := db.QueryRow(consult, nombre).Scan(&herramientaDetallada.IdH, &herramientaDetallada.URLImagen, &herramientaDetallada.Descripcion, &herramientaDetallada.Categoria, &herramientaDetallada.Estado, &herramientaDetallada.Licencia, &herramientaDetallada.EnlaceOficial, &herramientaDetallada.Nombre)
	if erro != nil {
		fmt.Println("No se encontro:", erro)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
	}
	t, err := template.ParseFiles("./templates/detalles.html")
	if err != nil {
		fmt.Println("Error al parsear:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	// Renderizar la plantilla con la herramienta
	err = t.Execute(c.Writer, herramientaDetallada)
	if err != nil {
		fmt.Println("Error al renderizar:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}

}

func AñadirFavorito(idUser int, c *gin.Context, db *sql.DB) {

	// Realiza el conteo en la base de datos
	consult := "SELECT Rol FROM Usuario WHERE id= ?;"

	// Ejecuta la consulta
	var rolLogueado string
	erro := db.QueryRow(consult, idUser).Scan(&rolLogueado)
	if erro != nil {
		fmt.Print("no se encontro usuario")
	}

	if rolLogueado == "Registrado" {
		consulta := "SELECT COUNT(*) FROM Favorito WHERE usuario_id= ? AND herramienta_id= ?;"

		// Ejecuta la consulta
		var repetidos int
		repetidos = 0
		erro := db.QueryRow(consulta, idUser, herramientaDetallada.IdH).Scan(&repetidos)
		if erro != nil {
			repetidos = 0
		}

		if repetidos == 0 {
			_, err := db.Exec("INSERT INTO Favorito (usuario_id, herramienta_id) VALUES (?, ?)", idUser, herramientaDetallada.IdH)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			fmt.Print("Añadido a favoritos.")
			c.JSON(http.StatusOK, gin.H{"exito": true, "mensaje": "Herramienta añadida a favoritos."})
		} else {
			c.JSON(http.StatusOK, gin.H{"exito": false, "mensaje": "Esta herramienta ya se encuentra en su lista de favoritos."})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{"exito": false, "mensaje": "Se requiere un usuario registrado y logueado para esta funcion."})

	}

}

type HerramientaDetall struct {
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

var herramientasDetall []HerramientaPri
