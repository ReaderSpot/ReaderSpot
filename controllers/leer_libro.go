package controllers

import (
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

func LeerLibro(ctx *gin.Context) {
	id := ctx.Param("id")
	var libro models.Libro
	db.DB.First(&libro, id)
	ctx.HTML(http.StatusOK, "leer_libro.html", libro)
}
