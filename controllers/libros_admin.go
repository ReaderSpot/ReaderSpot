package controllers

import (
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

func GetLibrosAdmin(ctx *gin.Context) {
	type Inicio struct {
		Libros []models.Libro
	}
	var libros []models.Libro
	db.DB.Find(&libros)
	data := Inicio{Libros: libros}
	ctx.HTML(http.StatusOK, "libros_admin.html", data)
}
