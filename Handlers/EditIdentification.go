package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditIdentification(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la identificacion desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la identificacion existe antes de intentar editarla
	var existingIdentificacion Models.Identificacion
	if err := conn.First(&existingIdentificacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Identificacion no encontrada"})
		return
	}

	// Obtener los datos actualizados del cuerpo de la solicitud
	var datosActualizados Models.Identificacion
	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar la identificacion en la base de datos
	if err := conn.Model(&existingIdentificacion).Updates(&datosActualizados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al editar la identificacion"})
		return
	}

	c.JSON(http.StatusOK, existingIdentificacion)
}
