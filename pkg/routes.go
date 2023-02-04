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
	
	todos, err := getAllTodos(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}

	log.Printf("todos: %v", todos)
	defer db.Close()

	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	
	log.Printf("route GET /todo/%v", id)

	var db, err = openDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}

	todo, err := getTodoById(db, id)
	defer db.Close()
	
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	} else {
		log.Printf("Todo by id is: %v", todo)
		c.IndentedJSON(http.StatusOK, todo)
	}
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
