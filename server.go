package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tux0010/Todo/database"
	"github.com/tux0010/Todo/todolist"
)

func main() {
	// Ensure all the env vars have been set properly
	tableName := os.Getenv("TODO_TABLE_NAME")
	if len(tableName) == 0 {
		fmt.Println(`"TODO_TABLE_NAME" env var must be set`)
		os.Exit(-1)
	}

	// Initialize the Database connection (sqlite3 in this case)
	conn, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	todolist.DB = database.NewDatabase(conn)
	defer conn.Close()

	// Register handlers
	r := mux.NewRouter()
	r.HandleFunc("/todo", todolist.GetAllToDoItemsAPI).Methods("GET")
	r.HandleFunc("/todo/new", todolist.CreateTodoItemAPI).Methods("POST")
	r.HandleFunc("/todo/{id}", todolist.DeleteToDoItemAPI).Methods("DELETE")

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)
	n.Run(":8080")
}
