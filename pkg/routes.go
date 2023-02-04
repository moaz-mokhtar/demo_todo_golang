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
	id := c.Param("id")

	log.Printf("Route: Get Todo by Id `GET /todo/:id=%v`", id)

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

func NewTodo(c *gin.Context) {
	log.Printf("Route: new todo `POST /todo`")

	var newTodo TodoItem
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	var db, err = openDB()
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
	idToDelete := c.Param("id")
	log.Printf("Route: Delete a todo `DELETE /todo/:id=`", idToDelete)

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

	log.Printf("Route: Update a todo item `PUT /todo/:id=`", id)

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
