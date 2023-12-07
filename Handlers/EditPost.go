package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditPostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la publicación desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe antes de intentar editarla
	var existingPost Models.Publicacion
	if err := conn.First(&existingPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
		return
	}

	// Obtener los datos actualizados del cuerpo de la solicitud
	var datosActualizados Models.Publicacion
	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar la publicación en la base de datos
	if err := conn.Model(&existingPost).Updates(&datosActualizados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al editar la publicación"})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}
