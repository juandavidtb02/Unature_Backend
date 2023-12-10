package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetPostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	// Obtener el ID de la publicación desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe
	var publicacion Models.Publicacion
	if err := conn.First(&publicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
		return
	}

	// Si la publicación existe, obtener todas las publicaciones y devolverlas en formato JSON
	var publicaciones []Models.Publicacion
	if err := conn.Preload("Usuario").Find(&publicaciones).Error; err != nil {
		log.Println("Error al obtener publicaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener publicaciones"})
		return
	}

	// Devolver todas las publicaciones en formato JSON
	c.JSON(http.StatusOK, publicaciones)
}
