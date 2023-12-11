package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeletePublication(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la aprobacion a eliminar desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe antes de intentar eliminarla
	var existingPublicacion Models.Publicacion
	if err := conn.First(&existingPublicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicacion no encontrada"})
		return
	}

	// Eliminar la publicación de la base de datos
	if err := conn.Delete(&existingPublicacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la Publicacion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Publicacion eliminada correctamente"})
}
