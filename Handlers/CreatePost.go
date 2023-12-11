package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreatePostHandler(c *gin.Context) {
	conn, _ := Connection.GetConnection()

	var nuevaPublicacion Models.Publicacion
	if err := c.ShouldBindJSON(&nuevaPublicacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el ID del usuario se proporciona en el JSON
	if nuevaPublicacion.UsuarioID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no proporcionado"})
		return
	}

	// Verificar si el usuario existe
	var usuario Models.Usuario
	if err := conn.First(&usuario, nuevaPublicacion.UsuarioID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Asignar el usuario a la publicación
	nuevaPublicacion.Usuario = usuario

	// Crear la publicación en la base de datos
	if err := conn.Create(&nuevaPublicacion).Error; err != nil {
		log.Println("Error al crear la publicación:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la publicación"})
		return
	}

	// Devolver la publicación creada en formato JSON
	c.JSON(http.StatusCreated, nuevaPublicacion)
}
