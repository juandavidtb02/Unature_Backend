package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

<<<<<<< HEAD
	// Obtener el ID de la publicación desde los parámetros de la URL
	id := c.Param("id")

	// Verificar si la publicación existe
	var publicacion Models.Publicacion
	if err := conn.First(&publicacion, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publicación no encontrada"})
=======
	var publicaciones []Models.Publicacion
	if err := conn.Preload("Usuario").Find(&publicaciones).Error; err != nil {
		log.Println("Error al obtener publicaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener publicaciones"})
>>>>>>> 200f79f8259af8e27230ebcad604214ac8f51dfe
		return
	}

	// Devolver la publicación en formato JSON
	c.JSON(http.StatusOK, publicacion)
}
