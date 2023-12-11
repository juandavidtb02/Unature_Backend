package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteIdentification(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la identificacion a eliminar desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la identificacion existe antes de intentar eliminarla
	var existingIdentificacion Models.Identificacion
	if err := conn.First(&existingIdentificacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Identificacion no encontrada"})
		return
	}

	// Eliminar la publicación de la base de datos
	if err := conn.Delete(&existingIdentificacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la identificacion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Identificación eliminada correctamente"})
}
