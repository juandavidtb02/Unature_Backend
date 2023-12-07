package main

import (
	"GORM/Handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//Migrate.Init()
	puerto := 8080
	r := gin.Default()

	r.GET("/posts", Handlers.GetPostsHandler)
	r.GET("/post/:id", Handlers.GetPostHandler)
	r.GET("/rol", Handlers.GetRolesHandler)

	r.POST("/rol", Handlers.CreateRoleHandler)
	r.POST("/image", Handlers.CreateImageHandler)
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
