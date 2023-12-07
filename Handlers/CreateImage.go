package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateImageHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevaImagen Models.Imagen
	if err := c.ShouldBindJSON(&nuevaImagen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la imagen en la base de datos
	if err := conn.Create(&nuevaImagen).Error; err != nil {
		log.Println("Error al crear la imagen:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la imagen"})
		return
	}

	// Devolver la imagen creada en formato JSON
	c.JSON(http.StatusCreated, nuevaImagen)
}
