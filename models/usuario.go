package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre        string        `gorm:"not null" json:"nombre" form:"nombre"`
	Email         string        `gorm:"unique;not null" json:"email" form:"email"`
	Password      string        `gorm:"not null" json:"-" form:"password"`
	IsAdmin       bool          `gorm:"default:false" json:"is_admin"`
	Transacciones []Transaccion `gorm:"foreignKey:UsuarioID"`
}
