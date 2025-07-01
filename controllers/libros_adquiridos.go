package controllers

import (
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetLibrosAdquiridos(ctx *gin.Context) {
	var transacciones []models.Transaccion
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	JWT, _ := jwt.Parse(obtenerJWT, func(JWT *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	var usuarioID float64
	if datosJWT, ok := JWT.Claims.(jwt.MapClaims); ok && JWT.Valid {
		usuarioID = datosJWT["usuarioID"].(float64)
	}
	db.DB.Preload("Libro").Where("usuario_id = ?", usuarioID).Find(&transacciones)
	ctx.HTML(http.StatusOK, "libros_adquiridos.html", gin.H{
		"Transacciones": transacciones,
	})
}
