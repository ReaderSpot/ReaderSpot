package controllers

import (
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"sistema_de_gestion_de_libros_electronicos/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

func VerificarFA(ctx *gin.Context) {
	var usuario models.Usuario
	obtenerJWT, err := ctx.Cookie("usuarioID")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"].(float64)
	codigo := ctx.PostForm("codigo")
	db.DB.First(&usuario, usuarioID)
	secret := usuario.Secret_2fa
	valid := totp.Validate(codigo, secret)
	if valid {
		isAuthenticated2fa := true
		tokenID, _ := utils.GetJWT(usuario.ID, usuario.IsAdmin, usuario.Email, usuario.Is_2fa, isAuthenticated2fa)
		ctx.SetCookie("usuarioID", tokenID, 3600, "/", "localhost", false, true)
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio")
		if usuario.IsAdmin {
			ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros")
		} else {
			ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio")
		}
	} else if !valid {
		ctx.Redirect(http.StatusSeeOther, "/login?error=codigo_incorrecto")
	}
}
