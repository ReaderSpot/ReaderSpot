package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func GetLibros(ctx *gin.Context) {
	var libros []models.Libro
	obtenerJWT, _ := ctx.Cookie("usuarioID")
	JWT, _ := jwt.Parse(obtenerJWT, func(JWT *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	var isAdmin bool
	if datosJWT, ok := JWT.Claims.(jwt.MapClaims); ok && JWT.Valid {
		isAdmin = datosJWT["isAdmin"].(bool)
	}
	redisClient := db.RedisClient()
	key := "libros_disponibles"
	cacheLibros, err := redisClient.Get(db.Ctx, key).Result()
	if err == redis.Nil {
		//redis.Nil redis no tiene la respuesta de db
		log.Print("Consultando a db para usar en redis")
		db.DB.Where("disponible = ?", true).Find(&libros)
		//JSON de la consulta db para redis cache
		librosJSON, _ := json.Marshal(libros)
		redisClient.Set(db.Ctx, key, librosJSON, 5*time.Minute)
	} else if err != nil {
		//No funciona redis, consulta db
		db.DB.Where("disponible = ?", true).Find(&libros)
	} else {
		log.Print("Usando db redis")
		//La consulta db es redis cache
		json.Unmarshal([]byte(cacheLibros), &libros)
	}

	ctx.HTML(http.StatusOK, "inicio.html", gin.H{
		"Libros":  libros,
		"IsAdmin": isAdmin,
	})
}
