package controllers

import (
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ComprarLibro(ctx *gin.Context) {
	//Obtiene el ID de la URL
	id := ctx.Param("id")
	var libro models.Libro
	db.DB.First(&libro, id)
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	JWT, _ := jwt.Parse(obtenerJWT, func(JWT *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	var usuarioID float64
	if datosJWT, ok := JWT.Claims.(jwt.MapClaims); ok && JWT.Valid {
		usuarioID = datosJWT["usuarioID"].(float64)
	}
	transaccion := models.Transaccion{
		Tipo:        "compra",
		LibroID:     libro.ID,
		UsuarioID:   uint(usuarioID),
		FechaInicio: time.Now(),
	}
	db.DB.Create(&transaccion)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio/libros/adquiridos")
}
