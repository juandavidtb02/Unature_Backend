package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIdentificationCount(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtén el ID de la publicacion desde los parámetros de la URL
	publicacionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de publicacion no válido"})
		return
	}

	// Consulta la base de datos para obtener el número de identificaciones asociadas a la publicacion
	var count int64
	result := conn.Model(&Models.Identificacion{}).Where("publicacion_id = ?", publicacionID).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el número de aprobaciones"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"identificaciones": count})
}
