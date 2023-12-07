package Models

import "gorm.io/gorm"

type Publicacion struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey;autoIncrement;not null;unique" json:"id_publicacion"`
	Descripcion string  `gorm:"type:text;not null" json:"descripcion"`
	Ubicacion   int     `gorm:"not null" json:"ubicacion"`
	Reportes    int     `gorm:"not null" json:"reportes"`
	UsuarioID   uint    `gorm:"not null" json:"id_usuario"`
	ImagenID    uint    `gorm:"not null" json:"id_imagen"`
	Usuario     Usuario `gorm:"foreignKey:UsuarioID"`
	Imagen      Imagen  `gorm:"foreignKey:ImagenID"`
	//Identificaciones []Identificacion `gorm:"foreignKey:PublicacionID"`
}
