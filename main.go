/*
Autor:
Tarea:
Fecha:
*/
package main

import (
	"sistema_de_gestion_de_libros_electronicos/controllers"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/middlewares"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConectarDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(models.Libro{}, models.Usuario{}, models.Transaccion{})
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	r.GET("/registro", func(ctx *gin.Context) {
		ctx.HTML(200, "registro.html", nil)
	})
	r.POST("/registro", controllers.Registro)
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "iniciar_sesion.html", nil)
	})
	r.POST("/login", controllers.Iniciar_sesion)
	r.Static("/static", "./static")
	r.GET("/verificarFA", func(ctx *gin.Context) {
		ctx.HTML(200, "verificar_FA.html", nil)
	})
	r.POST("/verificarFA", controllers.VerificarFA)
	//Endpoints a los que solo puede acceder un usuario autenticado /autenticado/ruta
	autenticado := r.Group("/autenticado")
	autenticado.Use(middlewares.ValidarUsuario())
	{
		autenticado.GET("/inicio", controllers.GetLibros)
		autenticado.GET("/logout", controllers.Logout)
		autenticado.POST("/inicio/libros/comprar/:id", controllers.ComprarLibro)
		autenticado.GET("/inicio/libros/adquiridos", controllers.GetLibrosAdquiridos)
		autenticado.GET("/inicio/libros/leer/:id", controllers.LeerLibro)
		autenticado.GET("/inicio/configuracion/seguridad/configurarFA", controllers.RegistrarFA)
		autenticado.POST("/inicio/configuracion/seguridad/configurarFA", controllers.ConfigurarFA)
		isAdmin := autenticado.Group("/admin")
		isAdmin.Use(middlewares.ValidarAdmin())
		isAdmin.GET("/libros", controllers.GetLibrosAdmin)
		isAdmin.GET("/libros/agregar", func(ctx *gin.Context) {
			ctx.HTML(200, "crear_libro.html", nil)
		})
		isAdmin.POST("/libros/agregar", controllers.AgregarLibro)
		isAdmin.GET("/libros/editar/:id", controllers.MostrarEditarLibro)
		isAdmin.POST("/libros/editar/:id", controllers.ActualizarLibro)
		isAdmin.POST("/libros/eliminar/:id", controllers.EliminarLibro)
		isAdmin.GET("/logout", controllers.Logout)
	}
	r.Run()
}
