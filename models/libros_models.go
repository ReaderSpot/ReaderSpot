package models

import (
	"gorm.io/gorm"
)

type Libro struct {
	gorm.Model
	Titulo string `gorm:"type:varchar(100);not null" json:"titulo"`
	Autor  string `gorm:"type:varchar(100);not null" json:"autor"`
	Leido  bool   `json:"leido"`
	UrlPDF string `gorm:"type:text;not null" json:"url_pdf"`
}
