package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No hay token"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		JWT_SECRET := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//token.Method valida que la firma de JWT sea HS256
			//_ no almacena el valor debido a que solo verifica
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWT_SECRET), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Token invalido"})
			c.Abort()
			return
		}
		datos_token, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pueden leer los datos del token"})
			c.Abort()
			return
		}
		//user_id es el nombre asignado en la funcion jwt.MapClaims en la funcion login
		userID, ok := datos_token["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User id incorrecto"})
			c.Abort()
			return
		}
		c.Set("user_id", uint(userID))
		c.Next()
	}
}
