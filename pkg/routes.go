package pkg

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success")
}

func GetAllTodos(c *gin.Context) {
	log.Printf("Route: Get all Todos `GET /todos`")
	var db, err = openDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}

	todos, err := getAllTodos(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}

	log.Printf("todos: %v", todos)
	defer db.Close()

	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	idStr := c.Param("id")
	log.Printf("Route: Get Todo by Id `GET /todo/:id=%v`", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	db, err := openDB()
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

func NewTodo(c *gin.Context) {
	log.Printf("Route: new todo `POST /todo`")

	var newTodo TodoItem
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	db, err := openDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	id, err := insertTodo(db, newTodo)
	log.Printf("New todo added id: %v", id)

	defer db.Close()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	} else {
		log.Printf("New todo inserted is: %v", newTodo)
		c.IndentedJSON(http.StatusOK, newTodo)
	}
}

func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	log.Printf("Route: delete a todo `DELETE /todo/:id=%v`", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	db, err := openDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	count, err := deleteTodo(db, id)
	defer db.Close()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		if count > 0 {
			message := fmt.Sprintf("Delete todo.id: %v done successfully. Count of rows affected by detele are: %v", idStr,count)
			log.Print(message)
			c.IndentedJSON(http.StatusOK, gin.H{"message": message})
		} else {
			message := fmt.Sprintf("No rows found to delete todo.id: %v. Count of affected rows is: %v", idStr, count)
			log.Print(message)
			c.IndentedJSON(http.StatusOK, gin.H{"message": message})
		}
	}
}

func UpdateTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, "Not implemented")
}
