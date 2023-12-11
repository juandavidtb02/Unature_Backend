package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	// Conectar a la base de datos
	conn, _ := Connection.GetConnection()

	// Obtener los datos del cuerpo de la solicitud
	var nuevoUsuario Models.Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el rol ya existe o crear uno nuevo
	var rol Models.Rol
	if err := conn.FirstOrCreate(&rol, Models.Rol{TipoRol: "usuario"}).Error; err != nil {
		log.Println("Error al obtener/crear el rol:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener/crear el rol"})
		return
	}

	// Encriptar la contraseña antes de almacenarla en la base de datos
	hashedContraseña, err := bcrypt.GenerateFromPassword([]byte(nuevoUsuario.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al encriptar la contraseña:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}

	nuevoUsuario.Password = string(hashedContraseña)
	nuevoUsuario.Rol = rol // Asignar el rol directamente

	// Crear el usuario en la base de datos
	if err := conn.Create(&nuevoUsuario).Error; err != nil {
		log.Println("Error al crear el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Devolver el usuario creado en formato JSON
	c.JSON(http.StatusCreated, nuevoUsuario)
}
