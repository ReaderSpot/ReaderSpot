package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Username string  `gorm:"unique; not null" json:"username" form:"username"`
	Nombre   string  `gorm:"not null" json:"nombre" form:"nombre"`
	Apellido string  `gorm:"not null" json:"apellido" form:"apellido"`
	Email    string  `gorm:"unique; not null" json:"email" form:"email"`
	Password string  `gorm:"not null" json:"-" form:"password"`
	Libro    []Libro `json:"libro" form:"libro"` //relacion uno a muchos db
}
