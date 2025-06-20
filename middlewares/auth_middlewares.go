package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken, err := c.Cookie("user_id")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		token, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		datosInterface := token.Claims
		datos, ok := datosInterface.(jwt.MapClaims)
		if !ok {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		getUserID := datos["user_id"].(float64)
		userID := uint(getUserID)
		c.Set("user_id", userID)
		c.Next()
	}
}
