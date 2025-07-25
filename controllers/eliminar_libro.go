package controllers

import (
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"

	"github.com/gin-gonic/gin"
)

func EliminarLibro(ctx *gin.Context) {
	//Obtiene el ID de la URL
	id := ctx.Param("id")
	var libro models.Libro
	db.DB.First(&libro, id)
	db.DB.Unscoped().Delete(&libro)
	//Al borrar libro eliminar cache redis
	redisClient := db.RedisClient()
	key := "libros_disponibles"
	redisClient.Del(db.Ctx, key)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros?success=libro_eliminado")
}
