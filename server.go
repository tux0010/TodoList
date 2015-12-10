package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tux0010/Todo/todolist"
)

func main() {
	// Ensure all the env vars have been set properly
	tableName := os.Getenv("TODO_TABLE_NAME")
	if len(tableName) == 0 {
		fmt.Println(`"TODO_TABLE_NAME" env var must be set`)
		os.Exit(-1)
	}

	// Register handlers
	r := mux.NewRouter()
	r.HandleFunc("/todo", todolist.GetAllToDoItemsAPI).Methods("GET")
	r.HandleFunc("/todo", todolist.CreateTodoItemAPI).Methods("POST")
	r.HandleFunc("/todo/{uuid}", todolist.DeleteToDoItemAPI).Methods("DELETE")

	// NOTE: Can use Negroni to add logging, etc
	http.ListenAndServe(":8080", r)
}
