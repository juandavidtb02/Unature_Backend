// Middleware/AuthMiddleware.go
package Middleware

import (
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
		c.Set("rol", claims.Role) // Agrega el rol al contexto

		// Si no se proporcionan roles, permitir el acceso
		if len(roles) == 0 {
			c.Next()
			return
		}

		// Verificar el rol del usuario
		for _, allowedRole := range roles {
			if claims.Role == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Acceso no autorizado"})
		c.Abort()
		return
	}
}
