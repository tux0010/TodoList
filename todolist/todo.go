package todolist

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tux0010/Todo/database"
)

type ToDoItem struct {
	Id      int       `json:"id"`
	Created time.Time `json:"created"`
	Text    string    `json:"text"`
}

var DB database.DatabaseDriver
var TableName string

func init() {
	TableName = os.Getenv("TODO_TABLE_NAME")
}

// NewToDoItem returns a pointer to a new ToDoItem object
func NewToDoItem(text string) *ToDoItem {
	item := &ToDoItem{
		Created: time.Now(),
		Text:    text,
	}

	return item
}

func getAllToDoItemsQuery(tableName string) string {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	return query
}

func getCreateToDoItemQuery(tableName string) string {
	query := fmt.Sprintf(
		"INSERT INTO %s (created, text) VALUES (?,?)",
		tableName,
	)

	return query
}

func getDeleteToDoItemQuery(tableName string, id string) string {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id ='%s'",
		tableName,
		id,
	)

	return query
}

func getToDoItemsFromDBRows(rows *sql.Rows) []ToDoItem {
	var toDoItemRows []ToDoItem

	for rows.Next() {
		item := ToDoItem{}
		rows.Scan(&item.Id, &item.Created, &item.Text)

		toDoItemRows = append(toDoItemRows, item)
	}

	return toDoItemRows
}

// GetAllToDoItems returns a list of all To-DO items stored
// in the database
func GetAllToDoItems() ([]ToDoItem, error) {
	query := getAllToDoItemsQuery(TableName)

	rows, err := DB.ExecuteQueryWithResponse(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	todoItems := getToDoItemsFromDBRows(rows.(*sql.Rows))

	return todoItems, nil
}

// CreateToDoItem will create a new To-Do item and save it
// in the database
func CreateToDoItem(text string) error {
	item := NewToDoItem(text)
	query := getCreateToDoItemQuery(TableName)

	return DB.ExecuteQuery(query, item.Created, item.Text)
}

// DeleteToDoItem will delete a To-Do item from the database
// given it's ID
func DeleteToDoItem(id string) error {
	query := getDeleteToDoItemQuery(TableName, id)

	return DB.ExecuteQuery(query)
}
