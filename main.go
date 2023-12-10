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
	Migrate.Init()
	r.GET("/post", Handlers.GetPostHandler)
	r.GET("/rol", Handlers.GetRolesHandler)
	r.POST("/rol", Handlers.CreateRoleHandler)
	r.POST("/post", Handlers.CreatePostHandler)
	r.POST("/signup", Handlers.CreateUserHandler)
	r.DELETE("/post/:id", Handlers.DeletePostHandler)

	fmt.Printf("El servidor está escuchando en el puerto %d...\n", puerto)
	err := r.Run(fmt.Sprintf(":%d", puerto))
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
