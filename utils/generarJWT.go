package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWT(usuarioID uint, isAdmin bool) (string, error) {
	//Asigna a una cookie el ID del usuario y si es Admin para validarlos en los middlewares
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuarioID": usuarioID,
		"isAdmin":   isAdmin, // <- asegúrate que `usuario.IsAdmin` sea un bool
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	//Obtiene el sercret del .env para validar el JWT
	tokenID, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenID, nil
}
