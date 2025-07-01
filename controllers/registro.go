package controllers

import (
	"fmt"
	"net/http"
	"sistema_de_gestion_de_libros_electronicos/db"
	"sistema_de_gestion_de_libros_electronicos/models"
	"sistema_de_gestion_de_libros_electronicos/utils"

	"github.com/gin-gonic/gin"
)

func Registro(ctx *gin.Context) {
	//Obtiene el usuario del cliente
	var usuario models.Usuario
	err_usuario := ctx.ShouldBind(&usuario)
	fmt.Print(usuario)
	if err_usuario != nil {
		ctx.Redirect(http.StatusSeeOther, "/login?error=datos_incorrectos")
		return
	}
	//Codifica la contraseña para la db
	hash, _ := utils.EncryptPassword(usuario.Password)
	//Asigna la contraseña codificada a la db mediante el modelo usuario
	usuario.Password = hash
	err_db := db.DB.Create(&usuario).Error
	if err_db != nil {
		ctx.Redirect(http.StatusSeeOther, "/login?error=no_se_pudo_registrar_el_usuario")
		return
	}
	//Genera un JWT y lo asigna en la cookie del cliente
	tokenID, _ := utils.GetJWT(usuario.ID, usuario.IsAdmin)
	ctx.SetCookie("usuarioID", tokenID, 86400, "/", "localhost", false, true)
	ctx.Redirect(http.StatusSeeOther, "/autenticado/inicio")
}
