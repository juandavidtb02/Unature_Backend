package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetRolesHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener todos los roles de la base de datos
	var roles []Models.Rol
	if err := conn.Find(&roles).Error; err != nil {
		log.Println("Error al obtener roles:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener roles"})
		return
	}

	// Devolver los roles en formato JSON
	c.JSON(http.StatusOK, roles)
}
