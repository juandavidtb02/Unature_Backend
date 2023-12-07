package Models

import "gorm.io/gorm"

type Aprobacion struct {
	gorm.Model
	IdAprobacion     uint           `gorm:"primaryKey; autoIncrement; not null;unique" json:"id_aprobacion"`
	UsuarioID        uint           `gorm:"not null" json:"id_usuario"`
	IdentificacionID uint           `gorm:"not null" json:"id_identificacion"`
	Usuario          Usuario        `gorm:"foreignKey:UsuarioID"`
	Identificacion   Identificacion `gorm:"foreignKey:IdentificacionID"`
}
