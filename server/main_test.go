package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHealth(t *testing.T) {
	r := SetUpRouter()
	r.GET("/", health)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllTodos(t *testing.T) {
	currentTodoList := getCurrentTodoList(t)

	r := SetUpRouter()
	r.GET("/todos", getAllTodos)
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody := loadJSONList(w.Body.String())

	t.Logf("currentTodoList type: %T\t -- Value: %v\n", currentTodoList, currentTodoList)
	t.Logf("w.Body type: %T\t -- Value: %v\n", w.Body, w.Body)
	t.Logf("responseBody type: %T\t -- Value: %v\n", responseBody, responseBody)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, currentTodoList, responseBody)
}

func TestGetTodoById(t *testing.T) {
	todoId := 2
	endpoint := fmt.Sprintf("/todo/%v", todoId)

	t.Logf("endpoint type: %T\t -- Value: %v\n", endpoint, endpoint)

	r := SetUpRouter()
	r.GET(endpoint, getTodoById)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("req type: %T\t -- Value: %v\n", req, req)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	t.Logf("w.Body type: %T\t -- Value: %v\n", w.Body, w.Body)

	// responseBody := loadJSONItem(w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	// assert.Equal(t, todoId, responseBody.ID)
}

func TestPostTodo(t *testing.T) {
	// Get list of current todo
	currentTodoList := getCurrentTodoList(t)
	t.Logf("currentTodoList type: %T\t -- Value: %v\n", currentTodoList, currentTodoList)

	// Create a new todo
	newTodo := todoItem{
		ID:          "4",
		Description: "Research learning resources for Go",
		Priority:    1,
	}
	t.Logf("newTodo type: %T\t -- Value: %v\n", newTodo, newTodo)
	newTodoBytes, _ := json.Marshal(newTodo)

	// Add new todo to the current todo list to be able to test final results
	expectedTodo := append(currentTodoList, newTodo)
	t.Logf("expectedTodo type: %T\t -- Value: %v\n", expectedTodo, expectedTodo)

	// Test `POST /todo`
	r := SetUpRouter()
	r.POST("/todo", postTodo)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewReader(newTodoBytes))
	t.Logf("req type: %T\t -- Value: %v\n", req, req)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	t.Logf("w.Body type: %T\t -- Value: %v\n", w.Body, w.Body)

	var responseBody []todoItem
	var bodyString = w.Body.String()
	json.Unmarshal([]byte(bodyString), &responseBody)
	t.Logf("responseBody type: %T\t -- Value: %v\n", responseBody, responseBody)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedTodo, responseBody)
}

func TestDeleteTodo(t *testing.T) {
	idToDelete := "1"
	endpoint := fmt.Sprintf("/todo/%v", idToDelete)

	t.Logf("endpoint type: %T\t -- Value: %v\n", endpoint, endpoint)

	r := SetUpRouter()
	r.DELETE(endpoint, deleteTodo)
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	t.Logf("req type: %T\t -- Value: %v\n", req, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	t.Logf("w.Body type: %T\t -- Value: %v\n", w.Body, w.Body)

	responseBody := loadJSONItem(w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, idToDelete, responseBody.ID)
}

func TestUpdateTodo(t *testing.T) {
	// TODO

}

func loadJSONList(s string) []todoItem {
	var todoList []todoItem
	json.Unmarshal([]byte(s), &todoList)
	return todoList
}

func loadJSONItem(s string) todoItem {
	var todo todoItem
	json.Unmarshal([]byte(s), &todo)
	return todo
}

func getCurrentTodoList(t *testing.T) []todoItem {
	t.Log("getcurrentTodoList")

	routerGet := SetUpRouter()
	// Get list of current todo
	routerGet.GET("/todos", getAllTodos)
	reqGet, _ := http.NewRequest("GET", "/todos", nil)
	recorderGet := httptest.NewRecorder()
	routerGet.ServeHTTP(recorderGet, reqGet)
	currentTodoList := loadJSONList(recorderGet.Body.String())

	t.Logf("getCurrentTodoList.currentTodoList type: %T\t -- Value: %v\n", currentTodoList, currentTodoList)

	return currentTodoList
}

func isTodoAvailable(t *testing.T, id string) bool {
	list := getCurrentTodoList(t)
	for _, todo := range list {
		if todo.ID == id {
			return true
		}
	}
	return false
}
