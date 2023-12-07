package Models

import "gorm.io/gorm"

type Identificacion struct {
	gorm.Model
	ID                   uint        `gorm:"primaryKey;autoIncrement;not null;unique" json:"id_identificacion"`
	Especie              string      `gorm:"type:varchar(20);not null" json:"especie"`
	Genero               string      `gorm:"type:varchar(20);not null" json:"genero"`
	NombreComun          string      `gorm:"type:varchar(20);not null" json:"nombre_comun"`
	InformacionAdicional string      `gorm:"type:text;not null" json:"informacion_adicional"`
	PublicacionID        uint        `gorm:"not null" json:"id_publicacion"`
	UsuarioID            uint        `gorm:"not null" json:"id_usuario"`
	Publicacion          Publicacion `gorm:"foreignKey:PublicacionID"`
	Usuario              Usuario     `gorm:"foreignKey:UsuarioID"`
	//Aprobaciones         []Aprobacion `gorm:"foreignKey:IdentificacionID"`
}
