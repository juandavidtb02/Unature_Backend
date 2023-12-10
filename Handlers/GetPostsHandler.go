package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostsHandler(c *gin.Context) {
	conn, err := Connection.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error de conexión a la base de datos"})
		return
	}

	// Obtener el ID de la publicación desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe
	var publicacion Models.Publicacion
	if err := conn.First(&publicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
		return
	}

	// Devolver la publicación en caso de éxito
	c.JSON(http.StatusOK, publicacion)
}
