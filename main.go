package main

import (
	"GORM/Handlers"
	"GORM/Migrate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//Migrate.Init()
	puerto := 8080
	r := gin.Default()
<<<<<<< HEAD

	r.GET("/posts", Handlers.GetPostsHandler)
	r.GET("/post/:id", Handlers.GetPostHandler)
=======
	Migrate.Init()
	r.GET("/post", Handlers.GetPostHandler)
>>>>>>> 200f79f8259af8e27230ebcad604214ac8f51dfe
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
