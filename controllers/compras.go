package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AgregarAlCarrito(ctx *gin.Context) {
	var carrito models.CarritoCompras
	libroIDStr := ctx.PostForm("book_id")
	libroID64, err := strconv.ParseUint(libroIDStr, 10, 64)
	if err != nil || libroID64 == 0 {
		log.Println("ID inválido al agregar al carrito:", libroIDStr)
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio?error=id_invalido")
		return
	}
	libroID := uint(libroID64)
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
	if err == nil {
		json.Unmarshal([]byte(value), &carrito)
	} else if err != redis.Nil {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio?error=no_se_agrego_a_comprar")
		return
	} else /*redis.Nil*/ {
		carrito = models.CarritoCompras{}
	}
	isCarrito := false
	for i, item := range carrito {
		if item.ID <= 0 {
			continue
		}
		if item.ID == libroID {
			carrito[i].Cantidad++
			isCarrito = true
			break
		}
	}
	if !isCarrito {
		carrito = append(carrito, models.Compras{ID: libroID, Cantidad: 1})
	}
	datos, _ := json.Marshal(carrito)
	redisClient.Set(db.Ctx, key, datos, 30*time.Minute)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio/libros/carrito")
}
