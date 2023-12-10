package Models

import "gorm.io/gorm"

type Publicacion struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey;autoIncrement;not null;unique" json:"id_publicacion"`
	Titulo      string  `gorm:"type:varchar(50);not null" json:"titulo"`
	Descripcion string  `gorm:"type:text;not null" json:"descripcion"`
	Ubicacion   int     `gorm:"not null" json:"ubicacion"`
	Reportes    int     `gorm:"not null" json:"reportes"`
	UsuarioID   uint    `gorm:"not null" json:"id_usuario"`
	UrlImagen   string  `gorm:"type:varchar(100);not null" json:"url_imagen"`
	Usuario     Usuario `gorm:"foreignKey:UsuarioID"`
	//Identificaciones []Identificacion `gorm:"foreignKey:PublicacionID"`
}
