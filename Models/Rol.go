package Models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey; autoIncrement;not null;unique" json:"id_rol"`
	TipoRol string `gorm:"type:varchar(20);not null" json:"tipo_rol"`
	//Usuarios []Usuario `gorm:"foreignKey:RolID"`
}
