package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener todos los roles de la base de datos
	var user []Models.Usuario
	if err := conn.Preload("Rol").Find(&user).Error; err != nil {
		log.Println("Error al obtener usuarios:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	// Devolver los roles en formato JSON
	c.JSON(http.StatusOK, user)
}
