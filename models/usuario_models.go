package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Username string  `gorm:"unique; not null" json:"username"`
	Nombre   string  `gorm:"not null" json:"nombre"`
	Apellido string  `gorm:"not null" json:"apellido"`
	Email    string  `gorm:"unique; not null" json:"email"`
	Password string  `gorm:"not null" json:"-"`
	Libro    []Libro `json:"libro"` //relacion uno a muchos db
}
