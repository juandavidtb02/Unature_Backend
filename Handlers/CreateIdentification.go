package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateIdentification(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevaIdentificacion Models.Identificacion
	if err := c.ShouldBindJSON(&nuevaIdentificacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la identificacion en la base de datos
	if err := conn.Create(&nuevaIdentificacion).Error; err != nil {
		log.Println("Error al crear la identificacion:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la identificacion"})
		return
	}

	// Devolver la publicación creada en formato JSON
	c.JSON(http.StatusCreated, nuevaIdentificacion)
}
