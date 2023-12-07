package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreatePostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevaPublicacion Models.Publicacion
	if err := c.ShouldBindJSON(&nuevaPublicacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear la publicación en la base de datos
	if err := conn.Create(&nuevaPublicacion).Error; err != nil {
		log.Println("Error al crear la publicación:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la publicación"})
		return
	}

	// Devolver la publicación creada en formato JSON
	c.JSON(http.StatusCreated, nuevaPublicacion)
}
