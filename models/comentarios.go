package models

import "gorm.io/gorm"

type Comentarios struct {
	gorm.Model
	Comentario string
	LibroID    uint
	UsuarioID  uint
	Libro      Libro   `gorm:"foreignKey:LibroID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Usuario    Usuario `gorm:"foreignKey: UsuarioID"`
}
