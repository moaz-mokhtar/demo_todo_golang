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

