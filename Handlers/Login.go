package Handlers

import (
	"GORM/Connection"
	"GORM/Models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Credentials struct {
	CorreoUsuario string `json:"email" binding:"required"`
	Contraseña    string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var secretKey = []byte("key123")

func GenerateToken(usuario Models.Usuario) (string, error) {
	claims := Claims{
		Username: usuario.CorreoUsuario,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expira en 1 hora
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginHandler(c *gin.Context) {
	var credentials Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credenciales inválidas"})
		return
	}

	conn, _ := Connection.GetConnection()

	var user Models.Usuario
	if err := conn.Where("correo_usuario = ?", credentials.CorreoUsuario).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Correo incorrecto"})
		return
	}

	// Comparar la contraseña proporcionada con la almacenada en la base de datos
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Contraseña)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectasss"})
		return
	}

	// Generar el token JWT
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID})
}
