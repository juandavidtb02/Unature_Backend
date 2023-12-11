package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditPublication(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la publicacion desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicacion existe antes de intentar editarla
	var existingPublicacion Models.Publicacion
	if err := conn.First(&existingPublicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicacion no encontrada"})
		return
	}

	// Obtener los datos actualizados del cuerpo de la solicitud
	var datosActualizados Models.Publicacion
	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar la identificacion en la base de datos
	if err := conn.Model(&existingPublicacion).Updates(&datosActualizados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al editar la publicación"})
		return
	}

	c.JSON(http.StatusOK, existingPublicacion)
}
