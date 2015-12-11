package todolist

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTodoItemAPI(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if len(text) == 0 {
		http.Error(w, "Missing 'text' GET parameter", http.StatusBadRequest)
		return
	}

	if err := CreateToDoItem(text); err != nil {
		log.Println(err.Error())
		http.Error(w, "Error creating TODO item", http.StatusInternalServerError)
		return
	}
}

func GetAllToDoItemsAPI(w http.ResponseWriter, r *http.Request) {
	todoItems, err := GetAllToDoItems()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to retrieve TODO items", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(todoItems)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error getting TODO items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteToDoItemAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := DeleteToDoItem(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to delete TODO item", http.StatusInternalServerError)
	}
}
