package Migrate

import (
	"GORM/Connection"
	"GORM/Models"
)

func Init() {

	conn, _ := Connection.GetConnection()

	conn.AutoMigrate(&Models.Rol{})
	conn.AutoMigrate(&Models.Usuario{})
	conn.AutoMigrate(&Models.Imagen{})
	conn.AutoMigrate(&Models.Publicacion{})
	conn.AutoMigrate(&Models.Identificacion{})
	conn.AutoMigrate(&Models.Aprobacion{})
}
