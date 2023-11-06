package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// Importa tu paquete de base de datos
)

func DetallesPage(c *gin.Context) {
	nombre := c.Param("nombre")
	herramienta, encontrado := herramientas[nombre]
	if !encontrado {
		c.String(http.StatusNotFound, "Producto no encontrado")
		return
	}

	// Renderiza la página de detalles del producto con los datos del producto

	c.HTML(http.StatusOK, "detalles.html", gin.H{
		"herramienta": herramienta})
	// Pasa datos a la plantilla si es necesario

}

type Herramienta struct {
	Nombre      string
	Descripcion string
	Categoria   float64
	URLImagen   string
	Estado      string
	Lincencia   string
	URL         string // Otros campos
}

var herramientas = map[string]Herramienta{
	"ChatGpt": Herramienta{
		Nombre: "ChatGpt",
		Descripcion: `ChatGPT es un chatbot de inteligencia artificial desarrollado por OpenAI que se especializa en el diálogo. Está compuesto por los modelos GPT-4 y GPT-3.5 de OpenAI, que han sido entrenados con un conjunto de datos masivo de texto y código.
		ChatGPT es capaz de mantener conversaciones con los usuarios de manera natural e informativa. Puede responder a preguntas, generar texto creativo y realizar tareas como traducciones y resúmenes.
		ChatGPT se puede usar para una variedad de propósitos, como:
		-Atención al cliente: ChatGPT se puede usar para responder a preguntas de los clientes y proporcionar asistencia.
		-Educación: ChatGPT se puede usar para proporcionar instrucción y apoyo a los estudiantes.
		-Entretenimiento: ChatGPT se puede usar para crear juegos, historias y otros contenidos creativos`,
		Categoria: 4.5,
		URLImagen: "/static/imagenes/Chatgpt.png",
		Estado:    "Activo",
		Lincencia: "La licencia de ChatGPT es de código abierto, lo que significa que cualquiera puede usar, modificar y redistribuir el código fuente del modelo",
		URL:       "https://chat.openai.com/?model=text-davinci-002-render-sha",
	},
	"ChatSonic": Herramienta{
		Nombre:      "ChatSonic",
		Descripcion: "Descripción del Producto 2",
		Categoria:   3.5,
		URLImagen:   "/static/imagenes/ChatSonic.jpg",
		//Estado:
		//Lincencia:
		//URL:
	},
	"Claude": Herramienta{
		Nombre:      "Claude",
		Descripcion: "Descripción del Producto 3",
		Categoria:   4.0,
		URLImagen:   "/static/imagenes/Claude.png",
		//Estado:
		//Lincencia:
		//URL:
	},
	"BingCreator": Herramienta{
		Nombre:      "BingCreator",
		Descripcion: "Descripción del Producto 4",
		Categoria:   4.8,
		URLImagen:   "/static/imagenes/BingCreator.jpg",
		//Estado:
		//Lincencia:
		//URL:
	},
	"ChatPDF": Herramienta{
		Nombre:      "ChatPDF",
		Descripcion: "Descripción del Producto 5",
		Categoria:   4.0,
		URLImagen:   "/static/imagenes/ChatPDF.jpg",
		//Estado:
		//Lincencia:
		//URL:
	},
	"Bard.png": Herramienta{
		Nombre:      "bard",
		Descripcion: "Descripción del Producto 6",
		Categoria:   3.8,
		URLImagen:   "/static/imagenes/bard.png",
		//Estado:
		//Lincencia:
		//URL:
	},
	"HumataAi": Herramienta{
		Nombre:      "HumataAi",
		Descripcion: "Descripción del Producto 7",
		Categoria:   4.2,
		URLImagen:   "/static/imagenes/HumataAi.jpg",
		//Estado:
		//Lincencia:
		//URL:
	},
	"AIVA AI": Herramienta{
		Nombre:      "AIVA AI",
		Descripcion: "Descripción del Producto 8",
		Categoria:   2.8,
		URLImagen:   "/static/imagenes/AIVA AI.png",
		//Estado:
		//Lincencia:
		//URL:
	},

	// Agrega más productos según sea necesario
}
