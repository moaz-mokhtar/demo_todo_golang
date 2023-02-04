package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"training/demo_todo/pkg"
)

func main() {
	fmt.Println("Hello, world!")

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", pkg.Health)
	router.GET("/todos", pkg.GetAllTodos)
	router.GET("/todo/:id", pkg.GetTodoById)
	router.POST("/todo", pkg.NewTodo)
	router.DELETE("/todo/:id", pkg.DeleteTodo)
	router.PUT("/todo/:id", pkg.UpdateTodo)

	router.Run("localhost:3000")
}



