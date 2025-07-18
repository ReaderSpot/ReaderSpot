package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

func RegistrarFA(ctx *gin.Context) {
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metodo incorrecto firma JWT")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	email, _ := claims["email"].(string)
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "ReaderSpot",
		AccountName: email,
	})
	var buffer bytes.Buffer
	img, _ := key.Image(200, 200)
	png.Encode(&buffer, img)
	display_qr := base64.StdEncoding.EncodeToString(buffer.Bytes())
	ctx.HTML(http.StatusOK, "configurar_FA.html", gin.H{
		"qr":     display_qr,
		"secret": key.Secret(),
	})
}
