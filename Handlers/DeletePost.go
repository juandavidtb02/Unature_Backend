package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeletePostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la publicación a eliminar desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe antes de intentar eliminarla
	var existingPost Models.Publicacion
	if err := conn.First(&existingPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
		return
	}

	// Eliminar la publicación de la base de datos
	if err := conn.Delete(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la publicación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Publicación eliminada correctamente"})
}
