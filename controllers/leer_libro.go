package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LeerLibro(ctx *gin.Context) {
	id := ctx.Param("id")
	var libro models.Libro
	var transacciones models.Transaccion
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metodo incorrecto firma JWT")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"].(float64)
	wasBought := db.DB.Where("usuario_id = ? AND libro_id = ?", usuarioID, id).First(&transacciones)
	if wasBought.Error != nil {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio?error=no_compraste_el_libro")
	}
	db.DB.First(&libro, id)
	ctx.HTML(http.StatusOK, "leer_libro.html", libro)
}
