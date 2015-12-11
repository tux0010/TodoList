package todolist

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
