package pkg

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "user:password@/tododb")
	if err == nil {
		log.Printf("Database opened successfully")
	}
	return db, err
}

// Function to get all todos
func getAllTodos(db *sql.DB) (todos []TodoItem, err error) {
	rows, err := db.Query("SELECT id, description, priority FROM todos")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo TodoItem
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Priority); err != nil {
			log.Fatal(err)
		}
		log.Printf("Scaned todo: id %d, description: %s, priority: %d\n", todo.Id, todo.Description, todo.Priority)
		todos = append(todos, todo)
	}
	if !rows.NextResultSet() {
		log.Printf("finished sets: %v", rows.Err())
	}

	log.Printf("Scaned todo: id %v", todos)

	return
}

// Function to get todo by id
func getTodoById(db *sql.DB, id int) (TodoItem, error) {
	log.Printf("Get Todo By Id: %d", id)

	var todo TodoItem
	err := db.QueryRow("SELECT * FROM todos where id=?;", id).Scan(&todo.Id, &todo.Description, &todo.Priority)

	if err == sql.ErrNoRows {
		log.Printf("no todo item with id %v\n", id)
		return todo, err
	} else {
		log.Printf("No sql.ErrNoRows found but found : %s", err.Error())
		return todo, nil
	}
}

// Function to insert a new todo item into the database.
// Returns id of the new todo item
func insertTodo(db *sql.DB, newTodo TodoItem) (int, error) {
	log.Printf("Insert a Todo: %v", newTodo)

	feedback, err := db.Exec("INSERT INTO todos (id, description, priority) VALUES (?, ?, ?);", newTodo.Id, newTodo.Description, newTodo.Priority)
	
	log.Printf("Insert exec feedback: %s", feedback)
	
	if err != nil {
		log.Printf("Error can't insert todo item %v\n", newTodo)
		return newTodo.Id, err
	}
	
	id, err := feedback.LastInsertId()
	log.Printf("feedback.LastInsertId(): %d", id)
	return newTodo.Id, err

}
