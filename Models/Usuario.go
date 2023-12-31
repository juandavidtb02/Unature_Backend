package Models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey; autoIncrement;not null;unique" json:"id_user"`
	CorreoUsuario string `gorm:"type:varchar(50);not null" json:"correo_usuario"`
	Nombre        string `gorm:"type:varchar(50);not null" json:"nombre_usuario"`
	Programa      string `gorm:"type:varchar(50);not null" json:"programa"`
	Password      string `gorm:"type:varchar(200);not null" json:"contraseña"`
	Semestre      int    `gorm:"not null" json:"semestre"`
	RolID         uint   `gorm:"not null" json:"id_rol"`
	Rol           Rol    `gorm:"foreignKey:RolID"`
	//Publicaciones    []Publicacion    `gorm:"foreignKey:UsuarioID"`
	//Aprobaciones     []Aprobacion     `gorm:"foreignKey:UsuarioID"`
	//Identificaciones []Identificacion `gorm:"foreignKey:UsuarioID"`
}

// En el modelo Usuario

func (u *Usuario) CountPublicaciones(db *gorm.DB) int {
	var count int64
	db.Model(u).Where("usuario_id = ?", u.ID).Count(&count)
	return int(count)
}
