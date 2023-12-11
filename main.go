package main

import (
	"GORM/Handlers"
	"GORM/Middleware"
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

	//Migrate.Init()
	// Rutas públicas (sin protección)
	r.GET("/posts", Handlers.GetPostsHandler)
	r.GET("/post/:id", Handlers.GetPostHandler)
	r.GET("/post", Handlers.GetPostsHandler)
	r.GET("/rol", Handlers.GetRolesHandler)
	r.GET("/user", Handlers.GetUserHandler)
	r.GET("/publicacion/:id/identificaciones", Handlers.GetIdentication)
	r.GET("/publicacion/:id/identificaciones/count", Handlers.GetIdentificationCount)
	r.GET("/identificacion/:id/aprobaciones", Handlers.GetAprobationCount)
	r.GET("/publications", Handlers.GetPublications)
	r.POST("/signup", Handlers.CreateUserHandler)
	r.POST("/login", Handlers.LoginHandler)
	// Rutas protegidas (requieren autenticación)

	r.Use(Middleware.AuthMiddleware())
	{
		r.DELETE("/post/:id", Handlers.DeletePostHandler)
		r.PUT("/post/:id", Handlers.EditPostHandler)
		r.POST("/post", Handlers.CreatePostHandler)
		r.POST("/identificacion", Handlers.CreateIdentification)

		r.PUT("/identificacion/:id", Handlers.EditIdentification)
		r.DELETE("/identificacion/:id", Handlers.DeleteIdentification)

		r.POST("/aprobacion", Handlers.CreateAprobation)
		r.DELETE("/aprobacion/:id", Handlers.DeleteAprobation)

		r.DELETE("/publication/:id", Handlers.DeletePublication)
		r.PUT("/publication/:id", Handlers.EditPublication)
	}
	r.Use(Middleware.AuthMiddleware("admin"))
	{
		r.POST("/rol", Handlers.CreateRoleHandler)

	}
	fmt.Printf("El servidor está escuchando en el puerto %d...\n", puerto)
	err := r.Run(fmt.Sprintf(":%d", puerto))
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
