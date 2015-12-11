package todolist

import (
	"fmt"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/tux0010/Todo/database"
)

func Test_GetAllToDoItemsQuery(t *testing.T) {
	assert := assert.New(t)

	expected := "SELECT * FROM todolist"
	assert.Equal(expected, getAllToDoItemsQuery("todolist"))
}

func Test_GetCreateToDoItemQuery(t *testing.T) {
	assert := assert.New(t)

	expected := "INSERT INTO todolist (created, text) VALUES (?,?)"
	assert.Equal(expected, getCreateToDoItemQuery("todolist"))
}

func Test_GetDeleteToDoItemQuery(t *testing.T) {
	assert := assert.New(t)

	expected := "DELETE FROM todolist WHERE id ='dead-beef'"
	assert.Equal(expected, getDeleteToDoItemQuery("todolist", "dead-beef"))
}

func Test_GetAllToDoItems_NoRowsFound(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	rows := sqlmock.NewRows([]string{"id", "created", "text"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	items, err := GetAllToDoItems()
	assert.Nil(err)
	assert.Empty(items)
}

func Test_GetAllToDoItems_RowsFound(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	rows := sqlmock.NewRows([]string{"id", "created", "text"})
	items := make([]ToDoItem, 2)

	for i := 0; i < 2; i++ {
		text := fmt.Sprintf("Item number %d", i)
		item := NewToDoItem(text)
		items[i] = *item
		rows.AddRow(i, item.Created, text)
	}

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	resItems, err := GetAllToDoItems()
	assert.Nil(err)
	assert.NotEmpty(items)
	assert.Equal(len(items), len(resItems))

	for i, item := range resItems {
		assert.Equal(item.Created, resItems[i].Created)
		assert.Equal(item.Id, resItems[i].Id)
		assert.Equal(item.Text, resItems[i].Text)
	}
}

func Test_CreateToDoItem_Failure(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewErrorResult(fmt.Errorf("some error"))
	mock.ExpectExec("^INSERT (.+)").WillReturnResult(result)
	err = CreateToDoItem("something")
	assert.NotNil(err)
}

func Test_CreateToDoItem_Success(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectExec("^INSERT (.+)").WillReturnResult(result)
	err = CreateToDoItem("something")
	assert.Nil(err)
}

func Test_DeleteToDoItem_Failure(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewErrorResult(fmt.Errorf("some error"))
	mock.ExpectExec("^DELETE (.+)").WillReturnResult(result)
	err = DeleteToDoItem("some_ID")
	assert.NotNil(err)
}

func Test_DeleteToDoItem_Success(t *testing.T) {
	assert := assert.New(t)

	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectExec("^DELETE (.+)").WillReturnResult(result)
	err = DeleteToDoItem("some_ID")
	assert.Nil(err)
}
