package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func PrincipalPage(c *gin.Context, db *sql.DB) {

	if len(herramientasPri) == 0 {
		// Realiza la consulta a la base de datos
		rows, err := db.Query("SELECT * FROM Herramienta")
		if err != nil {

			fmt.Print("tabla vacia")
		}
		defer rows.Close()

		for rows.Next() {
			var r HerramientaPri
			err := rows.Scan(&r.IdH, &r.URLImagen, &r.Descripcion, &r.Categoria, &r.Estado, &r.Licencia, &r.EnlaceOficial, &r.Nombre)
			if err != nil {
				fmt.Println("Error al escanear herramientas:", err)
			}
			herramientasPri = append(herramientasPri, r)
		}
	}
	t, err := template.ParseFiles("./templates/principal.html")
	if err != nil {
		fmt.Println("Error al parsear:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	// Renderizar la plantilla con la lista de herramientas
	err = t.Execute(c.Writer, herramientasPri)
	if err != nil {
		fmt.Println("Error al renderizar:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}
}

func BuscarHerramientas(c *gin.Context) {
	// Obtén el término de búsqueda del formulario HTML
	searchTerm := c.PostForm("search")
	herramientasBusqueda = herramientasBusqueda[:0]

	// Realiza una búsqueda en tus datos o base de datos para encontrar el producto por nombre
	BuscarProducto(searchTerm)

	t, err := template.ParseFiles("./templates/principal.html")
	if err != nil {
		fmt.Println("Error al parsear:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}

	// Renderizar la plantilla con la lista de herramientas
	err = t.Execute(c.Writer, herramientasBusqueda)
	if err != nil {
		fmt.Println("Error al renderizar:", err)
		c.String(http.StatusInternalServerError, "Error interno del servidor")
		return
	}
}

func CerrarSesion(c *gin.Context) int {
	usuarioDesloguea := 0
	c.HTML(http.StatusOK, "login.html", nil)
	return usuarioDesloguea
}

func BuscarProducto(searchTerm string) {

	for _, herramientaPri := range herramientasPri {
		if strings.ToLower(herramientaPri.Nombre) == strings.ToLower(searchTerm) || strings.ToLower(herramientaPri.Categoria) == strings.ToLower(searchTerm) {
			herramientasBusqueda = append(herramientasBusqueda, herramientaPri)
		}
	}

}

type HerramientaPri struct {
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

var herramientasBusqueda []HerramientaPri
var herramientasPri []HerramientaPri

/*var herramientasPri = []HerramientaPri{
	{"ChatGpt", "un modelo de lenguaje desarrollado por OpenAI.Estoy diseñado para comprender y generar texto en lenguaje natural en una variedad de temas y contextos", 4.5, "/static/imagenes/Chatgpt.png"},
	{"ChatSonic", "es un chatbot de inteligencia artificial desarrollado por Neuroflash, una empresa de IA con sede en España. Se basa en la tecnología GPT, la misma que se utiliza en ChatGPT, \n y está entrenado con un conjunto de datos masivo de texto y código", 3.5, "/static/imagenes/ChatSonic.jpg"},
	{"Claude", "es un asistente de IA desarrollado por Anthropic, una empresa de IA con sede en Estados Unidos. Se basa en el modelo de lenguaje Claude 2, que está entrenado con un conjunto de datos masivo de texto y código", 4.0, "/static/imagenes/Claude.png"},
	{"BingCreator", "es una herramienta de IA de Microsoft que permite a los usuarios crear contenido creativo, como imágenes, texto y música. Se basa en el modelo de lenguaje de gran tamaño DALL-E 2, que está entrenado con un conjunto de datos masivo de imágenes y texto", 4.8, "/static/imagenes/BingCreator.jpg"},
	{"ChatPDF", "es un chatbot de IA que puede resumir y responder a preguntas sobre cualquier PDF. Es una herramienta útil para estudiantes, investigadores y profesionales que necesitan acceder rápidamente a la información de los documentos PDF", 4.0, "/static/imagenes/ChatPDF.jpg"},
	{"bard", "es un gran modelo de lenguaje (LLM) de Google AI, entrenado en un conjunto de datos masivo de texto y código. Puede generar texto, traducir idiomas, escribir diferentes tipos de contenido creativo y responder a tus preguntas de manera informativa", 3.8, "/static/imagenes/bard.png"},
	{"AIVA AI", "es un asistente de composición musical impulsado por IA desarrollado por AIVA Technologies. Utiliza algoritmos de aprendizaje automático para generar piezas musicales originales en varios géneros, incluidos música clásica, jazz, pop y electrónica", 2.8, "/static/imagenes/AIVA AI.png"},

	// Agrega más productos según tus necesidades
}
*/
