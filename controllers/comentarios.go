package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ComentariosPOST(ctx *gin.Context) {
	libroIDPost := ctx.PostForm("libro_id")
	libroID, _ := strconv.ParseUint(libroIDPost, 10, 64)
	texto := ctx.PostForm("comentario")
	obtenerJWT, err := ctx.Cookie("usuarioID")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"].(float64)
	comentario := models.Comentarios{
		Comentario: texto,
		LibroID:    uint(libroID),
		UsuarioID:  uint(usuarioID),
	}
	db.DB.Create(&comentario)
	ctx.Redirect(303, fmt.Sprintf("/autenticado/inicio/libro?id=%d", libroID))
}
