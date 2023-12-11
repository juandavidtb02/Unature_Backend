package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Reportes(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	publicacionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de publicación no válido"})
		return
	}

	var post Models.Publicacion
	if err := conn.Preload("Usuario").Where("id = ?", publicacionID).Find(&post).Error; err != nil {
		log.Println("Error al obtener las identificaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las identificaciones"})
		return
	}

	// Aumentar en 1 el número de reportes
	post.Reportes++
	conn.Save(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Número de reportes aumentado correctamente"})
}
