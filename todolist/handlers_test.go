package todolist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/tux0010/Todo/database"
)

func Test_CreateTodoItemAPI_FailureMissingParameter(t *testing.T) {
	assert := assert.New(t)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"POST",
		"/todo",
		nil,
	)

	w := httptest.NewRecorder()
	CreateTodoItemAPI(w, req)

	assert.Equal(http.StatusBadRequest, w.Code)
}

func Test_CreateTodoItemAPI_Success(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectExec("^INSERT (.+)").WillReturnResult(result)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"POST",
		"/todo?text=Get%20produce%20from%20the%20grocery%20store",
		nil,
	)

	w := httptest.NewRecorder()
	CreateTodoItemAPI(w, req)

	assert.Equal(http.StatusCreated, w.Code)
}

func Test_CreateTodoItemAPI_FailureSQL(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewErrorResult(fmt.Errorf("some error"))
	mock.ExpectExec("^INSERT (.+)").WillReturnResult(result)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"POST",
		"/todo?text=Get%20produce%20from%20the%20grocery%20store",
		nil,
	)

	w := httptest.NewRecorder()
	CreateTodoItemAPI(w, req)

	assert.Equal(http.StatusInternalServerError, w.Code)
}

func Test_GetAllToDoItemsAPI_SuccessNoRows(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	rows := sqlmock.NewRows([]string{"id", "created", "text"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"GET",
		"/todo",
		nil,
	)

	w := httptest.NewRecorder()
	GetAllToDoItemsAPI(w, req)

	assert.Equal(http.StatusOK, w.Code)
}

func Test_GetAllToDoItemsAPI_SuccessRowsFound(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
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

	// Test the HTTP response
	req, _ := http.NewRequest(
		"GET",
		"/todo",
		nil,
	)

	w := httptest.NewRecorder()
	GetAllToDoItemsAPI(w, req)

	assert.Equal(http.StatusOK, w.Code)
}

func Test_DeleteToDoItemAPI_FailureSQL(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewErrorResult(fmt.Errorf("some error"))
	mock.ExpectExec("^DELETE (.+)").WillReturnResult(result)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"DELETE",
		"/todo/0",
		nil,
	)

	w := httptest.NewRecorder()
	DeleteToDoItemAPI(w, req)

	assert.Equal(http.StatusInternalServerError, w.Code)
}

func Test_DeleteToDoItemAPI_FailureNotFound(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewResult(0, 0)
	mock.ExpectExec("^DELETE (.+)").WillReturnResult(result)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"DELETE",
		"/todo/0",
		nil,
	)

	w := httptest.NewRecorder()
	DeleteToDoItemAPI(w, req)

	assert.Equal(http.StatusNotFound, w.Code)
}

func Test_DeleteToDoItemAPI_Success(t *testing.T) {
	assert := assert.New(t)

	// Initialize the mock DB
	conn, mock, err := sqlmock.New()
	assert.Nil(err)
	defer conn.Close()

	DB = database.NewDatabase(conn)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectExec("^DELETE (.+)").WillReturnResult(result)

	// Test the HTTP API
	req, _ := http.NewRequest(
		"DELETE",
		"/todo/0",
		nil,
	)

	w := httptest.NewRecorder()
	DeleteToDoItemAPI(w, req)

	assert.Equal(http.StatusNoContent, w.Code)
}
