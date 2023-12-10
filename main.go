package main

import (
	"GORM/Handlers"
	"GORM/Migrate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	puerto := 8080
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	Migrate.Init()
	r.GET("/posts", Handlers.GetPostsHandler)
	r.GET("/post/:id", Handlers.GetPostHandler)
	r.GET("/post", Handlers.GetPostsHandler)
	r.GET("/rol", Handlers.GetRolesHandler)
	r.POST("/rol", Handlers.CreateRoleHandler)
	r.POST("/post", Handlers.CreatePostHandler)
	r.POST("/signup", Handlers.CreateUserHandler)
	r.POST("/login", Handlers.LoginHandler)
	r.DELETE("/post/:id", Handlers.DeletePostHandler)
	r.PUT("/post/:id", Handlers.EditPostHandler)

	fmt.Printf("El servidor está escuchando en el puerto %d...\n", puerto)
	err := r.Run(fmt.Sprintf(":%d", puerto))
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
