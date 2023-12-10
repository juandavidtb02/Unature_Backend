package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetPostsHandler(c *gin.Context) {
	// Maneja la lógica de la respuesta aquí
	conn, _ := Connection.GetConnection()

	var publicaciones []Models.Publicacion
	if err := conn.Preload("Usuario").Find(&publicaciones).Error; err != nil {
		log.Println("Error al obtener publicaciones:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener publicaciones"})
		return
	}
	c.JSON(http.StatusOK, publicaciones)
}
