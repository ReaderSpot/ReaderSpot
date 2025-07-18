package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func ComprarLibro(ctx *gin.Context) {
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metodo incorrecto firma JWT")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"].(float64)
	redisClient := db.RedisClient()
	key := fmt.Sprintf("carrito:%v", usuarioID)
	value, err := redisClient.Get(db.Ctx, key).Result()
	if err == redis.Nil {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio/completarcompra?error=no_tienes_productos_para_pagar")
	}
	var carrito models.CarritoCompras
	json.Unmarshal([]byte(value), &carrito)
	for i, item := range carrito {
		var libro models.Libro
		db.DB.First(&libro, "id=?", carrito[i].ID)
		transaccion := models.Transaccion{
			Tipo:        "compra",
			LibroID:     libro.ID,
			UsuarioID:   usuarioID,
			FechaInicio: time.Now(),
			Precio:      libro.Precio,
			Cantidad:    item.Cantidad,
		}
		db.DB.Create(&transaccion)
	}
	redisClient.Del(db.Ctx, key)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio/libros/adquiridos")
}
