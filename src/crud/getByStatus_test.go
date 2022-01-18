package crud

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func TestGetByStatus(t *testing.T) {
	todos := []models.Todo{
		{Id: 1, Description: "FirstTodo", IsCompleted: false},
		{Id: 2, Description: "SecondTodo", IsCompleted: false},
		{Id: 3, Description: "ThirdTodo", IsCompleted: true},
	}

	todoManager := &TodoManager{todos}

	got := todoManager.GetByStatus(true)
	want := todos[2:]

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
