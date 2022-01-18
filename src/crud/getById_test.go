package crud

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func TestGetById(t *testing.T) {
	todos := []models.Todo{
		{1, "FirstTodo", false},
		{2, "SecondTodo", false},
		{3, "ThirdTodo", false},
	}

	todoManager := &TodoManager{todos}

	id := 1
	got := todoManager.GetById(id)
	want := todos[id-1]

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
