package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActualizarLibro(ctx *gin.Context) {
	//Obtiene el ID de la URL
	id := ctx.Param("id")
	var libro models.Libro
	db.DB.First(&libro, id)
	tituloRequest := ctx.PostForm("titulo")
	autorRequest := ctx.PostForm("autor")
	descripcionRequest := ctx.PostForm("descripcion")
	categoriaRequest := ctx.PostForm("categoria")
	precioStrRequest := ctx.PostForm("precio")
	precioRequest, _ := strconv.ParseFloat(precioStrRequest, 64)
	disponibleRequest := ctx.PostForm("disponible") == "true"
	libro.Titulo = tituloRequest
	libro.Autor = autorRequest
	libro.Descripcion = descripcionRequest
	libro.Categoria = categoriaRequest
	libro.Precio = precioRequest
	libro.Disponible = disponibleRequest
	portada, err := ctx.FormFile("portada")
	if err == nil && portada != nil && portada.Filename != "" {
		nombrePortada := filepath.Base(portada.Filename)
		path := fmt.Sprintf("static/portadas/%s", nombrePortada)
		if err := ctx.SaveUploadedFile(portada, path); err == nil {
			libro.Portada = "/" + path
		}
	}
	db.DB.Save(&libro)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros?success=libro_actualizado")
}
