// Middleware/AuthMiddleware.go
package Middleware

import (
	"GORM/Connection"
	"GORM/Models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Claims struct {
	Username string `json:"username"`
	Role     string
	jwt.StandardClaims
}

var secretKey = []byte("key123")

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// Verificar y extraer el token si comienza con "Bearer "
		if len(tokenString) >= 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Verificar el token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido"})
			c.Abort()
			return
		}

		// Agregar información del usuario al contexto para su uso posterior en las rutas protegidas
		c.Set("usuario", claims.Username)
		// Verificar el rol del usuario en la base de datos
		conn, _ := Connection.GetConnection()
		var user Models.Usuario
		if err := conn.Preload("Rol").Where("correo_usuario = ?", claims.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
			c.Abort()
			return
		}
		fmt.Println(roles)

		// Verificar el rol del usuario
		if len(roles) > 0 {
			for _, allowedRole := range roles {
				fmt.Println(user.Rol.TipoRol, allowedRole)
				if user.Rol.TipoRol == allowedRole {
					c.Next()
					return
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso no autorizado"})
			c.Abort()
			return
		}
		c.Next()
		return

	}
}
