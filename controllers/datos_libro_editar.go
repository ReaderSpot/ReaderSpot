package controllers

import (
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

/*
Esta funcion envia los datos del libro editado en el get de /admin/libros/editar para saber que datos se quieren cambiar
*/

func MostrarEditarLibro(ctx *gin.Context) {
	//Obtiene el ID de la URL
	id := ctx.Param("id")

	var libro models.Libro
	db.DB.First(&libro, id)

	ctx.HTML(http.StatusOK, "editar_libro.html", gin.H{
		"ID":          libro.ID,
		"Titulo":      libro.Titulo,
		"Autor":       libro.Autor,
		"Descripcion": libro.Descripcion,
		"Categoria":   libro.Categoria,
		"Precio":      libro.Precio,
		"Portada":     libro.Portada,
		"Disponible":  libro.Disponible,
	})
}
