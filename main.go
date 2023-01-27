package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", health)
	router.GET("/todos", getAllTodos)
	router.GET("/todo/:id", getTodoById)
	router.POST("/todo", postTodo)
	router.DELETE("/todo/:id", deleteTodo)
	router.PUT("/todo/:id", updateTodo)

	router.Run("localhost:3000")
}

type todoItem struct {
	ID          string `json:"id" uri:"id" binding:"required"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

// todo slice to seed record todo data.
var todos = []todoItem{
	{
		ID:          "1",
		Description: "Learn Go language",
		Priority:    1,
	},
	{
		ID:          "2",
		Description: "Participate in Open Source with Go",
		Priority:    3,
	},
	{
		ID:          "3",
		Description: "Create sample project with Go",
		Priority:    1,
	},
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success")
}

func getAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")

	// log.Printf("getTodoById=>Context type: %T\t -- Value: %v\n", c, c)
	log.Printf("getTodoById=>id type: %T\t -- Value: %v\n", id, id)

	for _, item := range todos {
		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func postTodo(c *gin.Context) {
	var newTodo todoItem

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	idToDelete := c.Param("id")
	var todoAfterDelete []todoItem

	for _, item := range todos {
		if item.ID != idToDelete {
			todoAfterDelete = append(todoAfterDelete, item)
		}
	}

	todos = todoAfterDelete

	c.IndentedJSON(http.StatusOK, "Success")
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var itemToUpdate todoItem

	if err := c.BindJSON(&itemToUpdate); err != nil {
		return
	}

	for index, item := range todos {
		if item.ID == id {
			todos[index] = itemToUpdate
			c.IndentedJSON(http.StatusOK, itemToUpdate)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
