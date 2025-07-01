package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	//Elimina la cookie del cliente
	ctx.SetCookie("usuarioID", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}
