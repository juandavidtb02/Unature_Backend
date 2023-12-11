package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GlobalData(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var posts []Models.Publicacion
	if err := conn.Preload("Usuario").Find(&posts).Error; err != nil {
		log.Println("Error al obtener las identificaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	var iden []Models.Identificacion
	if err := conn.Preload("Usuario").Preload("Publicacion").Find(&iden).Error; err != nil {
		log.Println("Error al obtener las identificaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	var user []Models.Usuario
	if err := conn.Preload("Rol").Find(&user).Error; err != nil {
		log.Println("Error al obtener usuarios:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	publicacionJSON := map[string]interface{}{
		"num_publicaciones":    len(posts),
		"num_identificaciones": len(iden),
		"num_usuarios":         len(user),
	}

	c.JSON(http.StatusOK, publicacionJSON)
}
