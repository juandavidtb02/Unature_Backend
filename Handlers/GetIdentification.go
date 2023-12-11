package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIdentication(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtén el ID de la publicación desde los parámetros de la URL
	publicacionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de publicación no válido"})
		return
	}

	// Consulta la base de datos para obtener las identificaciones asociadas a la publicación
	var identificaciones []Models.Identificacion
	result := conn.Where("publicacion_id = ?", publicacionID).Find(&identificaciones)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	c.JSON(http.StatusOK, identificaciones)
}
