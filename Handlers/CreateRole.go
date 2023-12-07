package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateRoleHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevoRol Models.Rol
	if err := c.ShouldBindJSON(&nuevoRol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear el rol en la base de datos
	if err := conn.Create(&nuevoRol).Error; err != nil {
		log.Println("Error al crear el rol:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el rol"})
		return
	}

	// Devolver la imagen creada en formato JSON
	c.JSON(http.StatusCreated, nuevoRol)
}
