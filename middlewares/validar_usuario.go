package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidarUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getToken, err := ctx.Cookie("usuarioID")
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		obtenerDatos := token.Claims
		datosCookie, ok := obtenerDatos.(jwt.MapClaims)
		if !ok {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		obtenerID, ok := datosCookie["usuarioID"].(float64)
		if !ok {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		usuarioID := uint(obtenerID)

		ctx.Set("usuarioID", usuarioID)
		ctx.Set("claims", datosCookie)

		ctx.Next()
	}
}
