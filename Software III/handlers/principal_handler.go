package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrincipalPage(c *gin.Context) {
	// Obtén el término de búsqueda del formulario HTML
	searchTerm := c.PostForm("search")

	// Realiza una búsqueda en tus datos o base de datos para encontrar el producto por nombre
	result := BuscarProductoPorNombre(searchTerm)

	// Pasa el resultado a la plantilla HTML
	c.HTML(http.StatusOK, "principal.html", gin.H{
		"Result": result,
	})
}

func BuscarProductoPorNombre(searchTerm string) *Producto {
	// Implementa la lógica de búsqueda real aquí, busca el producto por nombre
	// Devuelve el producto si se encuentra, o nil si no se encuentra
	// Esto es un ejemplo simplificado
	for _, producto := range Productos {
		if producto.Nombre == searchTerm {
			return &producto
		}
	}
	return nil
}

type Producto struct {
	Nombre       string
	Descripcion  string
	Calificacion float64
	URLImagen    string
	// Agrega otros campos relacionados con el producto según sea necesario
}

var Productos = []Producto{
	Producto{"ChatGpt", "un modelo de lenguaje desarrollado por OpenAI.Estoy diseñado para comprender y generar texto en lenguaje natural en una variedad de temas y contextos", 4.5, "/static/imagenes/Chatgpt.png"},
	Producto{"Producto 2", "Descripción del producto 2", 3.5, "product2.jpg"},
	// Agrega más productos según tus necesidades
}
