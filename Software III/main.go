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
	/*db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/Chatgpt.png", "un modelo de lenguaje desarrollado por OpenAI.Esta diseñado para comprender y generar texto en lenguaje natural en una variedad de temas y contextos", "chatbot", "vigente", "3.5 Gratuito, 4 Licencia", "https://chat.openai.com/", "ChatGpt")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/ChatSonic.jpg", "es un chatbot de inteligencia artificial desarrollado por Neuroflash, una empresa de IA con sede en España. Se basa en la tecnología GPT, la misma que se utiliza en ChatGPT, y está entrenado con un conjunto de datos masivo de texto y código", "chatbot", "vigente", "Gratuito o Licencia", "https://writesonic.com/chat", "ChatSonic")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/Claude.png", "es un asistente de IA desarrollado por Anthropic, una empresa de IA con sede en Estados Unidos. Se basa en el modelo de lenguaje Claude 2, que está entrenado con un conjunto de datos masivo de texto y código", "Asistente", "vigente", "Gratuito", "https://claude.ai/", "Claude")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/BingCreator.jpg", "es una herramienta de IA de Microsoft que permite a los usuarios crear contenido creativo, como imágenes, texto y música. Se basa en el modelo de lenguaje de gran tamaño DALL-E 2, que está entrenado con un conjunto de datos masivo de imágenes y texto", "Generador", "vigente", "Gratuito", "https://www.bing.com/create?toWww=1&redig=143AACB8726144D5B4F942652F93FC1D", "BingCreator")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/ChatPDF.jpg", "es un chatbot de IA que puede resumir y responder a preguntas sobre cualquier PDF. Es una herramienta útil para estudiantes, investigadores y profesionales que necesitan acceder rápidamente a la información de los documentos PDF", "Asistente", "vigente", "Gratuito", "https://www.chatpdf.com/", "ChatPDF")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/bard.png", "es un gran modelo de lenguaje (LLM) de Google AI, entrenado en un conjunto de datos masivo de texto y código. Puede generar texto, traducir idiomas, escribir diferentes tipos de contenido creativo y responder a tus preguntas de manera informativa", "chatbot", "vigente", "Gratuito", "https://bard.google.com/?hl=es", "Bard")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/AIVA AI.png", "es un asistente de composición musical impulsado por IA desarrollado por AIVA Technologies. Utiliza algoritmos de aprendizaje automático para generar piezas musicales originales en varios géneros, incluidos música clásica, jazz, pop y electrónica", "Generador", "vigente", "Gratuito", "https://www.aiva.ai/", "AIVA AI")
	db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/HumataAi.jpg", "Es la herramienta de inteligencia artificial que te permite obtener respuestas instantáneas, generar informes rápidamente y comprender documentos técnicos y legales de forma eficiente. Con capacidades avanzadas de procesamiento de lenguaje natural, Humata facilita la adquisición de nuevos conocimientos.", "Asistente", "vigente", "Gratuito o Licencias", "https://www.humata.ai/", "Humata AI")
	*/ //db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/DeepDream.png", "Es un algoritmo desarrollado por Google que utiliza redes neuronales convolucionales para analizar e interpretar imágenes. A diferencia de otros modelos de visión por computadora que se utilizan para reconocimiento de objetos, DeepDream se destaca por su capacidad para resaltar y enfatizar patrones visuales dentro de una imagen.", "Generador", "vigente", "Prueba free y Licencias", "https://deepdreamgenerator.com/", "Deep Dream")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/DialogFlow.png", "Es una plataforma de desarrollo de chatbots y asistentes virtuales, basada en la nube, que utiliza procesamiento del lenguaje natural (NLP) para comprender y responder a las consultas de los usuarios de manera interactiva.", "chatbot", "vigente", "Cobro por sesion", "https://cloud.google.com/dialogflow", "DialogFlow")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/MicrosoftBotFramework.png", "Es un lienzo de creación visual de código abierto para que desarrolladores y equipos multidisciplinarios diseñen y creen experiencias conversacionales con Language Understanding, QnA Maker y una composición sofisticada de respuestas de bot (Language Generation).", "Generador", "vigente", "Codigo abierto", "https://dev.botframework.com/", "Microsoft Bot Framework Composer")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/IBMWatson.jpg", "Es una plataforma de inteligencia artificial conversacional líder en el mercado diseñada para ayudarle a superar la fricción del servicio de asistencia tradicional y ofrecer experiencias excepcionales a clientes potenciales, clientes y empleados.", "Asistente", "vigente", "Licencia prueba y de pago", "https://www.ibm.com/mx-es/products/watsonx-assistant", "IBM Watson Assistant")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/Siri.jpg", "Es un asistente virtual desarrollado por Apple que utiliza procesamiento del lenguaje natural para responder a comandos de voz en dispositivos Apple, realizando diversas tareas y proporcionando información.", "Asistente", "vigente", "Integrado en dispositivos Apple", "https://www.apple.com/es/siri/", "Siri")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/GoogleAssistant.png", "Es un asistente virtual desarrollado por Google que responde a comandos de voz y realiza diversas funciones, como buscar información, controlar dispositivos del hogar inteligente y ejecutar aplicaciones en dispositivos Android y otros dispositivos compatibles.", "Asistente", "vigente", "Gratuito en dispositivos compatibles", "https://assistant.google.com/intl/es_es/", "Google Assistant")
	//db.Exec("INSERT INTO Herramienta (urlimagen, descripcion, categoria, estado, licencia, enlaceOficial, nombre) VALUES (?, ?, ?, ?, ?, ?, ?)", "/static/imagenes/RunWayML.png", "Es una plataforma que permite a los usuarios crear fácilmente modelos de inteligencia artificial para generar contenido, como imágenes, audio y video, sin necesidad de conocimientos profundos en programación.", "Generador", "vigente", "Gratuito o Licencia", "https://runwayml.com/", "RunwayML")
	return db
}

func main() {
	var usuarioLog int

	db := OpenConnection()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/login", handlers.LoginPage)
	r.GET("/principal", func(c *gin.Context) {
		handlers.PrincipalPage(c, db)
	})
	r.GET("/register", handlers.RegisterPage)
	r.GET("/detalles/:nombre", func(c *gin.Context) {
		handlers.DetallesPage(c, db)
	})
	r.GET("/favoritos", func(c *gin.Context) {
		handlers.FavoritosPage(c, db, usuarioLog)
	})

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

	r.POST("/busqueda", handlers.BuscarHerramientas)
	r.POST("/cerrarSesion", func(c *gin.Context) {
		usuarioLog = handlers.CerrarSesion(c)
		fmt.Print("usuario logueado: ", usuarioLog, " ")
	})
	r.POST("/añadirFavorito", func(c *gin.Context) {
		handlers.AñadirFavorito(usuarioLog, c, db)
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
