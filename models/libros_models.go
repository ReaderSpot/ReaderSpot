package models

import (
	"gorm.io/gorm"
)

type Libro struct {
	gorm.Model
	Titulo          string `gorm:"type:varchar(100);not null" json:"titulo" form:"titulo"`
	Autor           string `gorm:"type:varchar(100);not null" json:"autor" form:"autor"`
	Leido           bool   `gorm:"not null" json:"leido" form:"leido"`
	UrlPDF          string `gorm:"type:text;not null" json:"url_pdf" form:"url_pdf"`
	NombreCategoria string `json:"nombre_categoria" form:"nombre_categoria"`
	UsuarioID       uint
	CategoriaID     uint      `form:"categoria_id"`
	Usuario         Usuario   `gorm:"foreignKey:UsuarioID;constrain:OnDelete:CASCADE;" json:"-"`
	Categoria       Categoria `gorm:"foreignKey:CategoriaID;constrain:OnDelete:CASCADE;" json:"categoria,omitempty"`
}
