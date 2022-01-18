package crud

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func TestGetAll(t *testing.T) {
	t.Run("GetAllTodos should return all the todos added", func(t *testing.T) {
		todos := []models.Todo{
			{Id: 1, Description: "FirstTodo", IsCompleted: false},
			{Id: 2, Description: "SecondTodo", IsCompleted: false},
			{Id: 3, Description: "ThirdTodo", IsCompleted: false},
		}

		todoManager := &TodoManager{todos}

		assertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})

	t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
		todos := []models.Todo{}

		todoManager := &TodoManager{todos}

		assertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})
}

func assertCollectionsEqual(t *testing.T, got, want []models.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
