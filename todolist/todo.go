package todolist

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	"github.com/tux0010/Todo/database"
)

type ToDoItem struct {
	UUID      string    `json:"uuid"`
	Created   time.Time `json:"created"`
	Completed bool      `json:"completed"`
	Text      string    `json:"text"`
}

var DB database.DatabaseDriver
var TableName = os.Getenv("TODO_TABLE_NAME")

func NewToDoItem(text string) *ToDoItem {
	item := &ToDoItem{
		UUID:      uuid.NewV4().String(),
		Created:   time.Now(),
		Completed: false,
		Text:      text,
	}

	return item
}

func getAllToDoItemsQuery(tableName string) string {
	query := fmt.Sprintf("SELECT * FROM %s;", tableName)

	return query
}

func getCreateToDoItemQuery(tableName string, item *ToDoItem) string {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, created, completed, text) VALUES (",
		tableName,
	)

	query += fmt.Sprintf(
		"'%s', '%d', '%s', '%s'",
		item.UUID,
		item.Created.Unix(),
		strconv.FormatBool(item.Completed),
		item.Text,
	)

	query += ");"

	return query
}

func getDeleteToDoItemQuery(tableName string, uuid string) string {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE uuid ='%s';",
		tableName,
		uuid,
	)

	return query
}

func getToDoItemsFromDBRows(rows *sql.Rows) []ToDoItem {
	var toDoItemRows []ToDoItem

	for rows.Next() {
		item := ToDoItem{}
		rows.Scan(&item.UUID, &item.Created, &item.Completed, &item.Text)

		toDoItemRows = append(toDoItemRows, item)
	}

	return toDoItemRows
}

// GetAllToDoItems returns a list of all To-DO items stored
// in the database
func GetAllToDoItems() ([]ToDoItem, error) {
	query := getAllToDoItemsQuery(TableName)

	rows, err := DB.ExceuteQueryWithResponse(query)
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
	query := getCreateToDoItemQuery(TableName, item)

	return DB.ExecuteQuery(query)
}

// DeleteToDoItem will delete a To-Do item from the database
// given it's UUID
func DeleteToDoItem(uuid string) error {
	query := getDeleteToDoItemQuery(TableName, uuid)

	return DB.ExecuteQuery(query)
}
