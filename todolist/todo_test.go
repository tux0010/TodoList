package todolist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllToDoItemsQuery(t *testing.T) {
	assert := assert.New(t)

	expected := "SELECT * FROM todolist;"
	assert.Equal(expected, getAllToDoItemsQuery("todolist"))
}

func Test_GetCreateToDoItemQuery(t *testing.T) {
	assert := assert.New(t)

	item := NewToDoItem("Hello world!")
	expected := "INSERT INTO todolist (uuid, created, completed, text) VALUES ("
	expected += fmt.Sprintf(
		"'%s', '%d', '%s', '%s'",
		item.UUID,
		item.Created.Unix(),
		"false",
		"Hello world!",
	)
	expected += ");"

	assert.Equal(expected, getCreateToDoItemQuery("todolist", item))
}

func Test_GetDeleteToDoItemQuery(t *testing.T) {
	assert := assert.New(t)

	expected := "DELETE FROM todolist WHERE uuid ='dead-beef';"
	assert.Equal(expected, getDeleteToDoItemQuery("todolist", "dead-beef"))
}
