package Models

import "gorm.io/gorm"

type Imagen struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement;not null;unique" json:"id_imagen"`
	UrlImagen string `gorm:"type:varchar(100);not null" json:"url_imagen"`
}
