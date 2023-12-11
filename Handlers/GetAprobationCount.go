package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAprobationCount(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtén el ID de la identificación desde los parámetros de la URL
	identificacionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de identificación no válido"})
		return
	}

	// Consulta la base de datos para obtener el número de aprobaciones asociadas a la identificación
	var count int64
	result := conn.Model(&Models.Aprobacion{}).Where("identificacion_id = ?", identificacionID).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"aprobaciones": count})
}
