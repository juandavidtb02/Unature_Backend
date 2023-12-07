package Models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey; autoIncrement;not null;unique" json:"id_user"`
	CorreoUsuario string `gorm:"type:varchar(50);not null" json:"correo_usuario"`
	Contraseña    string `gorm:"type:varchar(200);not null" json:"contraseña"`
	Semestre      int    `gorm:"not null" json:"semestre"`
	RolID         uint   `gorm:"not null" json:"id_rol"`
	Rol           Rol    `gorm:"foreignKey:RolID"`
	//Publicaciones    []Publicacion    `gorm:"foreignKey:UsuarioID"`
	//Aprobaciones     []Aprobacion     `gorm:"foreignKey:UsuarioID"`
	//Identificaciones []Identificacion `gorm:"foreignKey:UsuarioID"`
}
