package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func CompletarCompraGET(ctx *gin.Context) {
	var carrito models.CarritoCompras
	var total float64
	var detallesCompra []models.DetallesLibro
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	token, _ := jwt.Parse(obtenerJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metodo incorrecto firma JWT")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	usuarioID := claims["usuarioID"]
	redisClient := db.RedisClient()
	key := fmt.Sprintf("carrito:%v", usuarioID)
	value, err := redisClient.Get(db.Ctx, key).Result()
	if err == redis.Nil {
		ctx.HTML(http.StatusOK, "completarcompra.html", gin.H{
			"Libros":  nil,
			"Mensaje": "No tienes productos",
		})
		return
	} else if err != nil {
		log.Print("Redis: ", err)
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio?error=no_funciona_redis")
		return
	}
	json.Unmarshal([]byte(value), &carrito)
	for i, item := range carrito {
		var libro models.Libro
		db.DB.First(&libro, "id = ?", carrito[i].ID)
		subtotal := libro.Precio * float64(item.Cantidad)
		total += subtotal
		detallesCompra = append(detallesCompra, models.DetallesLibro{
			ID:       libro.ID,
			Cantidad: item.Cantidad,
			Titulo:   libro.Titulo,
			Precio:   libro.Precio,
			Subtotal: subtotal,
		})
	}
	ctx.HTML(http.StatusOK, "completarcompra.html", gin.H{
		"Libros":  detallesCompra,
		"Usuario": usuarioID,
		"Total":   total,
	})
}
