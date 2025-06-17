package main

import (
	"autonomo_dos/controllers"
	"autonomo_dos/db"
	"autonomo_dos/middlewares"
	"autonomo_dos/models"
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
	router := gin.Default()
	/*	|------------------------------|
		|Crea la configuracion del CORS|
		|------------------------------|
	*/
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
	authenticator := router.Group("/auth")
	authenticator.POST("/registro", controllers.RegistrarUsuario)
	authenticator.POST("/login", controllers.LoginUsuario)
	validate := router.Group("/valid")
	/*
		-----------------------------------------
		Middleware para verificar si esta logueado
		-----------------------------------------
	*/
	validate.Use(middlewares.LoginMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "GOOD",
		})
	})
	validate.GET("/libros", controllers.BuscarLibroJWT)
	validate.POST("/libros", controllers.CrearLibroJWT)
	router.Run(":8080")
}

/*
func main() {
	DB, err := db.ConectarDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer DB.Close()

	fmt.Println("Conexión exitosa a la base de datos MySQL.")
}
*/
