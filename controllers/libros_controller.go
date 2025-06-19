package controllers

import (
	"autonomo_dos/db"
	"autonomo_dos/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
)

func BuscarLibroJWT(c *gin.Context) {
	log.Print("Funcion buscar libro")
	//user_id debe coincidir con el nombre asignado en la funcion jwt.MapClaims en la funcion login
	obtener_userID, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}
	token, err := jwt.Parse(obtener_userID, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}
	datosInterface := token.Claims
	datos, ok := datosInterface.(jwt.MapClaims)
	if ok {
		getUserID := datos["user_id"].(float64)
		userID := uint(getUserID)
		log.Printf("User ID: %d", userID)
		c.Set("user_id", userID)
	}
	userID := datos["user_id"].(float64)
	type LibroRespuesta struct {
		ID              uint      `json:"id"`
		Titulo          string    `json:"titulo"`
		Autor           string    `json:"autor"`
		Leido           bool      `json:"leido"`
		UrlPdf          string    `json:"url_pdf"`
		NombreCategoria string    `json:"nombre_categoria"`
		CreatedAt       time.Time `json:"created_at"`
	}
	//Obtiene las categorias del usuario para el menu agregar libro
	var Categoria []models.Categoria
	db_categoria := db.DB.Where("usuario_id = ?", userID).Find(&Categoria)
	if db_categoria.Error != nil {
		c.String(http.StatusInternalServerError, "No se pudo cargar las categorias")
		return
	}
	//Obtiene los libros del usuario para la ruta /auth/libros
	var Libros []LibroRespuesta
	db := db.DB.Table("libros").Select("libros.id, libros.titulo, libros.autor, libros.leido, libros.url_pdf, COALESCE(categoria.nombre, '') as nombre_categoria").
		Joins("LEFT JOIN categoria ON libros.categoria_id = categoria.id").
		Where("libros.usuario_id = ?", userID).
		Find(&Libros)
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener libros de la base de datos"})
		return
	}
	c.HTML(http.StatusOK, "libros.html", gin.H{
		"Libros":    Libros,
		"Categoria": Categoria,
	})
}

func CrearLibroJWT(c *gin.Context) {
	var libro models.Libro
	libro_error := c.ShouldBindWith(&libro, binding.Form)
	log.Print("Libro registrado: ", libro)
	if libro_error != nil {
		c.Redirect(http.StatusSeeOther, "/valid/libros/agregar?error=formato_incorrecto")
		return
	}
	obtener_userID, existe := c.Get("user_id")
	if !existe {
		c.Redirect(http.StatusSeeOther, "/login?error=no_has_iniciado_sesion")
		return
	}
	userID := obtener_userID.(uint)
	libro.UsuarioID = userID
	db_libro := db.DB.Create(&libro)
	if db_libro.Error != nil {
		c.Redirect(http.StatusSeeOther, "/valid/libros/agregar?error=no_se_pudo_agregar_el_libro")
		return
	}
	c.Redirect(http.StatusSeeOther, "/valid/libros?success=libro_agregado")
}

func ActualizarLibroJWT(c *gin.Context) {
}

// Arreglar funcion de eliminar libros con el frontend, tal vez problema de redirecciones
func BorrarLibroJWT(c *gin.Context) {
	var libro models.Libro
	libroIDUrl := c.Param("id")
	libroID, err := strconv.ParseUint(libroIDUrl, 10, 64)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/valid/libros?error=id_incorrecto")
		return
	}
	db := db.DB.Delete(&libro, libroID)
	if db.Error != nil {
		c.Redirect(http.StatusSeeOther, "/valid/libros?error=no_se_pudo_borrar_libro")
		return
	}
	c.Redirect(http.StatusSeeOther, "/valid/libros?success=libro_eliminado")
}
