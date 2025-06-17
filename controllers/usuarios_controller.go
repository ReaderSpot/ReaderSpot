package controllers

import (
	"autonomo_dos/db"
	"autonomo_dos/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetPerfilUsuario(w http.ResponseWriter, r *http.Request) {

}
func ActualizarUsuario(w http.ResponseWriter, r *http.Request) {

}
func EliminarUsuario(w http.ResponseWriter, r *http.Request) {

}
func RegistrarUsuario(c *gin.Context) {
	var usuario models.Usuario
	err_usuario := c.ShouldBindJSON(&usuario)
	if err_usuario != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos incorrectos"})
		return
	}
	hash, err := hashearPassword(usuario.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el hash para la contraseña"})
		return
	}
	usuario.Password = hash
	err_db := db.DB.Create(&usuario).Error
	if err_db != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo registrar el usuario"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario registrado",
	})
}

func LoginUsuario(c *gin.Context) {
	var usuarioEntrante struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err_json := c.ShouldBindJSON(&usuarioEntrante)
	if err_json != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON incorrecto"})
		return
	}
	var usuario models.Usuario
	db := db.DB.Where("email = ?", usuarioEntrante.Email).First(&usuario)
	if db.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No existe una cuenta registrada con el email"})
		return
	}
	if !checkHashPassword(usuarioEntrante.Password, usuario.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": usuario.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar JWT"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Login exitoso",
		"token":   tokenString,
	})
}

func hashearPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
