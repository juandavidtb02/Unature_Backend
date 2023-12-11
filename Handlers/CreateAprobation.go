package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateAprobation(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevaAprobacion Models.Aprobacion
	if err := c.ShouldBindJSON(&nuevaAprobacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la aprobacion en la base de datos
	if err := conn.Create(&nuevaAprobacion).Error; err != nil {
		log.Println("Error al crear la Aprobación:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la aprobación"})
		return
	}

	// Devolver la publicación creada en formato JSON
	c.JSON(http.StatusCreated, nuevaAprobacion)
}
