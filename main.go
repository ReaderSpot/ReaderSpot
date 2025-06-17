package main

import (
	"autonomo_dos/controllers"
	"autonomo_dos/db"
	"autonomo_dos/middlewares"
	"autonomo_dos/models"

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
	gin := gin.Default()
	gin.POST("/registro", controllers.RegistrarUsuario)
	gin.POST("/login", controllers.LoginUsuario)
	autenticar := gin.Group("/auth")
	autenticar.Use(middlewares.LoginMiddleware())
	autenticar.GET("/libros", controllers.BuscarLibroJWT)
	autenticar.POST("/libros", controllers.CrearLibroJWT)
	gin.Run()
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
