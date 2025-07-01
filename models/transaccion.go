package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaccion struct {
	gorm.Model
	Tipo        string     `gorm:"not null" json:"tipo"` // "compra" o "prestamo"
	LibroID     uint       `gorm:"not null" json:"libro_id"`
	UsuarioID   uint       `gorm:"not null" json:"usuario_id"`
	FechaInicio time.Time  `gorm:"not null" json:"fecha_inicio"`
	FechaFin    *time.Time `json:"fecha_fin"`

	Libro   Libro   `gorm:"foreignKey:LibroID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // <- Esta línea es necesaria
	Usuario Usuario `gorm:"foreignKey:UsuarioID"`                                             // <- Opcional, pero útil
}
