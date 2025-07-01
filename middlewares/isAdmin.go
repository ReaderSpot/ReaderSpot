package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidarAdmin() gin.HandlerFunc {
	//Valida si en la cookie el valor de isAdmin es true
	return func(ctx *gin.Context) {
		valor, existe := ctx.Get("claims")
		if !existe {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "no autorizado"})
			ctx.Abort()
			return
		}

		claims, ok := valor.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "formato de claims inválido"})
			ctx.Abort()
			return
		}

		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok || !isAdmin {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "se requiere ser administrador"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
