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

func PrestamoLibro(ctx *gin.Context) {
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
	fechaInicio := time.Now()
	//Establece que los libros se devuelven en 1 semana
	fechaFin := fechaInicio.Add(7 * 24 * time.Hour)
	transaccion := models.Transaccion{
		Tipo:        "prestamo",
		LibroID:     libro.ID,
		UsuarioID:   uint(usuarioID),
		FechaInicio: fechaInicio,
		FechaFin:    &fechaFin,
	}
	db.DB.Create(&transaccion)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio/libros/adquiridos")
}
