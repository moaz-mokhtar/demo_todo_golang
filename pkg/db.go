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
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Priority); err != nil {
			log.Fatal(err)
		}
		log.Printf("Scaned todo: id %s, description: %s, priority: %d\n", todo.ID, todo.Description, todo.Priority)
		todos = append(todos, todo)
	}
	if !rows.NextResultSet() {
		log.Printf("finished sets: %v", rows.Err())
	}

	log.Printf("Scaned todo: id %v", todos)

	return
}

// Function to get all todos
func getTodoById(db *sql.DB, id string) (TodoItem, error) {
	log.Printf("Get Todo By Id: %s", id)

	var todo TodoItem
	err := db.QueryRow("SELECT * FROM todos where id=?;", id).Scan(&todo.ID, &todo.Description, &todo.Priority)

	if err == sql.ErrNoRows {
		log.Printf("no todo item with id %v\n", id)
		return todo, err
	} else {
		log.Printf("No sql.ErrNoRows found but found : %s", err.Error())
		return todo, nil
	}
}
