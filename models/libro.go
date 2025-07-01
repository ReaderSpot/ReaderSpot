package models

import "gorm.io/gorm"

type Libro struct {
	gorm.Model
	Titulo      string  `gorm:"not null" json:"titulo" form:"titulo"`
	Autor       string  `gorm:"not null" json:"autor" form:"autor"`
	Descripcion string  `json:"descripcion" form:"descripcion"`
	Categoria   string  `gorm:"not null" form:"categoria"`
	Disponible  bool    `gorm:"default:true" json:"disponible" form:"disponible"`
	Precio      float64 `json:"precio" form:"precio"`
	Portada     string  `form:"portada" json:"portada"`
}
