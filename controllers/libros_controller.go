package controllers

import (
	"autonomo_dos/db"
	"autonomo_dos/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BuscarLibroJWT(c *gin.Context) {
	//user_id debe coincidir con el nombre asignado en la funcion jwt.MapClaims en la funcion login
	obtener_userID, existe := c.Get("user_id")
	if !existe {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No has iniciado sesión"})
		return
	}
	userID := obtener_userID.(uint)
	type LibroResponse struct {
		ID              uint      `json:"id"`
		Titulo          string    `json:"titulo"`
		Autor           string    `json:"autor"`
		Leido           bool      `json:"leido"`
		UrlPdf          string    `json:"url_pdf"`
		NombreCategoria string    `json:"nombre_categoria"`
		CreatedAt       time.Time `json:"created_at"`
	}
	var libroRespuesta []LibroResponse
	db := db.DB.Table("libros").Select("libros.id, libros.titulo, libros.autor, libros.leido, libros.url_pdf, COALESCE(categoria.nombre, '') as nombre_categoria").
		Joins("LEFT JOIN categoria ON libros.categoria_id = categoria.id").
		Where("libros.usuario_id = ?", userID).
		Find(&libroRespuesta)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener libros de la base de datos"})
		return
	}
	if len(libroRespuesta) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no ha registrado libros"})
	}
	c.JSON(http.StatusOK, libroRespuesta)
}

func CrearLibroJWT(c *gin.Context) {
	var libro models.Libro
	libro_error_json := c.ShouldBindJSON(&libro)
	if libro_error_json != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de libro incorrecto"})
		return
	}
	obtener_userID, existe := c.Get("user_id")
	if !existe {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No has iniciado sesión"})
		return
	}
	userID := obtener_userID.(uint)
	libro.UsuarioID = userID
	var categoria models.Categoria
	if libro.CategoriaID != 0 {
		db_categoria := db.DB.Where("(usuario_id = ? OR usuario_id IS NULL) AND id = ?", userID, libro.CategoriaID).First(&categoria)
		if db_categoria.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La categoría seleccionada no existe"})
			return
		}
	} else if libro.NombreCategoria != "" {
		db_categoria := db.DB.Where("(usuario_id = ? OR usuario_id IS NULL) AND nombre = ?", userID, libro.NombreCategoria).First(&categoria)
		if db_categoria.Error != nil {
			categoria = models.Categoria{
				Nombre:    libro.NombreCategoria,
				UsuarioID: &userID,
			}
			db_categoria := db.DB.Create(&categoria)
			if db_categoria.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar la categoria"})
				return
			}
		}
	}
	libro.CategoriaID = categoria.ID
	db_libro := db.DB.Create(&libro)
	if db_libro.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar el libro"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensaje": "Libro agregado"})
}

func ActualizarLibroJWT(c *gin.Context) {
}
