package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey; not null" json:"libro_id"`
	Nombre    string `gorm:"unique; not null" json:"nombre_categoria"`
	Libros    []Libro
	UsuarioID *uint
	Usuario   Usuario `gorm:"foreignKey:UsuarioID"`
}
