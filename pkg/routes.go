package pkg

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success")
}

func GetAllTodos(c *gin.Context) {
	var db, _ = openDB()
	log.Printf("db: %v", db)
	todos, _ := getAllTodos(db)
	log.Printf("todos: %v", todos)
	defer db.Close()

	
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
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

func PostTodo(c *gin.Context) {
	var newTodo TodoItem

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context) {
	idToDelete := c.Param("id")
	var todoAfterDelete []TodoItem

	for _, item := range todos {
		if item.ID != idToDelete {
			todoAfterDelete = append(todoAfterDelete, item)
		}
	}

	todos = todoAfterDelete

	c.IndentedJSON(http.StatusOK, "Success")
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var itemToUpdate TodoItem

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
