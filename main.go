package main

import (
	"ReaderSpot/controllers"
	"ReaderSpot/db"
	"ReaderSpot/middlewares"
	"ReaderSpot/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConectarDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(models.Libro{}, models.Categoria{}, models.Usuario{})
	if err != nil {
		panic(err)
	}
	/*	|------------------------------|
		|Crea la configuracion del CORS|
		|------------------------------|
	*/
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		//direccion de solicitudes permitidas
		AllowOrigins: []string{"http://readerspot.xyz"},
		/*	____________________________________
			|Metodos permitidos a la direccion |
			|__________________________________|
		*/
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		/*	____________________________________
			|Permite Headers personalizados	   |
			|__________________________________|
		*/
		AllowHeaders: []string{"Origin"},
		/*	_____________________________________________________
			|Permite exponer directamente los Headers al usuario |
			|____________________________________________________|
		*/
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		/*	______________________________________
			|Cache de cada solicitud echa a la API|
			|_____________________________________|
		*/
		MaxAge: 6 * time.Hour,
	}))

	/*
		-------------------
		endpoint de prueba
		-------------------
	*/
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/registro", func(c *gin.Context) {
		c.HTML(200, "registro.html", nil)
	})
	router.POST("/registro", controllers.RegistrarUsuario)
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	router.POST("/login", controllers.LoginUsuario)
	validate := router.Group("/valid")
	validate.Use(middlewares.AuthMiddleware())
	{
		validate.GET("/libros", controllers.BuscarLibroJWT)
		validate.POST("/libros", controllers.CrearLibroJWT)
		validate.POST("/libros/categoria_nueva", controllers.AgregarCategoriaJWT)
	}
	router.Run()
}
