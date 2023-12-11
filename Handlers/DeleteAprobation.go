package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteAprobation(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la aprobacion a eliminar desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la aprobacion existe antes de intentar eliminarla
	var existingAprobacion Models.Aprobacion
	if err := conn.First(&existingAprobacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aprobacion no encontrada"})
		return
	}

	// Eliminar la publicación de la base de datos
	if err := conn.Delete(&existingAprobacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la Aprobacion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aprobacion eliminada correctamente"})
}
