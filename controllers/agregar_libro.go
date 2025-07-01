package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

func AgregarLibro(ctx *gin.Context) {
	//Obtiene los datos del cliente
	var libro models.Libro
	ctx.ShouldBind(&libro)
	//Agregar portada
	portada, _ := ctx.FormFile("portada")
	nombrePortada := filepath.Base(portada.Filename)
	path := fmt.Sprintf("static/portadas/%s", nombrePortada)
	ctx.SaveUploadedFile(portada, path)
	libro.Portada = "/" + path
	//Agregar PDF libro
	libroArchivo, _ := ctx.FormFile("url_pdf")
	nombreArchivoLibro := filepath.Base(libroArchivo.Filename)
	pathLibro := fmt.Sprintf("static/libro/%s", nombreArchivoLibro)
	ctx.SaveUploadedFile(libroArchivo, pathLibro)
	libro.UrlPDF = "/" + pathLibro
	db := db.DB.Create(&libro)
	if db.Error != nil {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros?error=no_se_agrego_el_libro")
	}
	ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros")
}
