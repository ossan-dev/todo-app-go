package test

import (
	"reflect"
	"testing"

	"todo-app-go.com/v1/src/models"
)

func assertCollectionsEqual(t *testing.T, got, want []models.Todo) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestGetAll(t *testing.T) {
	t.Run("GetAllTodos should return all the todos added", func(t *testing.T) {
		todos := []models.Todo{
			{1, "FirstTodo", false},
			{2, "SecondTodo", false},
			{3, "ThirdTodo", false},
		}

		todoManager := &models.TodoManager{todos}

		assertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})

	t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
		todos := []models.Todo{}

		todoManager := &models.TodoManager{todos}

		assertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})
}
