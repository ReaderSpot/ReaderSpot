package controllers

import (
	"ReaderSpot/db"
	"ReaderSpot/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
		NombreCategoria string    `form:"nombre_categoria"`
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
	err := c.ShouldBind(&libro)
	if err != nil {
		log.Println("Error en el binding del formulario:", err)
		c.Redirect(http.StatusSeeOther, "/valid/libros/agregar?error=formato_incorrecto")
		return
	}
	log.Print("Libro recibido del formulario: ", libro)
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
	if !ok {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	userID := uint(datos["user_id"].(float64))
	log.Printf("User ID: %d", userID)

	libro.UsuarioID = userID
	result := db.DB.Create(&libro)
	if result.Error != nil {
		log.Println("Error al guardar el libro:", result.Error)
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

func AgregarCategoriaJWT(c *gin.Context) {
	var categoria models.Categoria
	err := c.ShouldBind(&categoria)
	//Obtiene el usr_id del JWT
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
	categoria.UsuarioID = uint(userID)
	//Agrega categoria
	db := db.DB.Create(&categoria)
	if db.Error != nil {
		c.Redirect(http.StatusSeeOther, "/valid/libros/agregar?error=No_se_pudo_agregar_la_categoria")
	}
	c.Redirect(http.StatusSeeOther, "/valid/libros?success=Se_agrego_la_categoria")
}
