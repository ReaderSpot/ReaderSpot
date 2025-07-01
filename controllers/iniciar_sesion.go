package controllers

import (
	"errors"
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"sistema_de_gestion_de_libros_electronicos/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Iniciar_sesion(ctx *gin.Context) {
	//Obtiene el usuario y la contraseña del cliente
	var usuarioRequest models.Usuario
	ctx.ShouldBind(&usuarioRequest)
	//Busca el usuario en la base de datos
	var usuario models.Usuario
	db := db.DB.Where("email = ?", usuarioRequest.Email).First(&usuario)
	if db.Error != nil {
		// Si hay error, no existe el usuario
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			ctx.Redirect(http.StatusSeeOther, "/?error=cuenta_inexistente")
			return
		}
		ctx.Redirect(http.StatusSeeOther, "/?error=error_db")
		return
	}
	//Verifica si la contraseña es correcta con la funcion en el modulo utils
	if !utils.VerfiicarPassword(usuarioRequest.Password, usuario.Password) {
		ctx.Redirect(http.StatusSeeOther, "/login?error=contraseña_incorrecta")
		return
	}
	//Genera un JWT y lo asigna en la cookie del cliente
	tokenID, _ := utils.GetJWT(usuario.ID, usuario.IsAdmin)
	ctx.SetCookie("usuarioID", tokenID, 86400, "/", "localhost", false, true)
	//Valida si el usuario es Admin
	if usuario.IsAdmin {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/admin/libros")
	} else {
		ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio")
	}
}
