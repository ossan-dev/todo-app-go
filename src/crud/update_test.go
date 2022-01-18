package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestUpdate(t *testing.T) {
	todos := []models.Todo{
		models.NewTodo(1, "FirstTodo", false),
		models.NewTodo(2, "SecondTodo", false),
		models.NewTodo(3, "ThirdTodo", true),
	}

	todoManager := &TodoManager{todos}

	todoUpdated := models.NewTodo(2, "Updated text", false)
	todoManager.Update(todoUpdated)

	got, err := todoManager.GetById(todoUpdated.Id)

	utils.AssertNoError(t, err)
	utils.AssertTodosEqual(t, got, &todoUpdated)
}
