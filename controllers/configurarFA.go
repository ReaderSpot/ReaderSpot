package controllers

import (
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

func ConfigurarFA(ctx *gin.Context) {
	var usuario models.Usuario
	secret := ctx.PostForm("secret")
	codigo := ctx.PostForm("codigo")
	obtenerJWT, err := ctx.Cookie("usuarioID")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"].(float64)
	valid := totp.Validate(codigo, secret)
	if valid {
		db.DB.Model(&usuario).Where("id = ?", usuarioID).Updates(map[string]interface{}{
			"is_2fa":     true,
			"secret_2fa": secret,
		})
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio?successful=cuenta_seguridad")
	} else if !valid {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/configuracion/seguridad/2fa?error=codigo_incorrecto")
	}
}
