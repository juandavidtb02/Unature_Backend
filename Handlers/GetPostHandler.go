package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la publicación desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe
	var publicacion Models.Publicacion
	if err := conn.First(&publicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
		return
	}

	// Devolver la publicación en formato JSON
	c.JSON(http.StatusOK, publicacion)
}
