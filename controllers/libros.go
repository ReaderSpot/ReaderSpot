package controllers

import (
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetLibros(ctx *gin.Context) {
	var libros []models.Libro
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	JWT, _ := jwt.Parse(obtenerJWT, func(JWT *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	var isAdmin bool
	if datosJWT, ok := JWT.Claims.(jwt.MapClaims); ok && JWT.Valid {
		isAdmin = datosJWT["isAdmin"].(bool)
	}
	db.DB.Where("disponible = ?", true).Find(&libros)
	ctx.HTML(http.StatusOK, "inicio.html", gin.H{
		"Libros":  libros,
		"IsAdmin": isAdmin,
	})
}
