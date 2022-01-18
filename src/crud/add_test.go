package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestAdd(t *testing.T) {
	todos := []models.Todo{
		models.NewTodo(1, "FirstTodo", false),
		models.NewTodo(2, "SecondTodo", false),
		models.NewTodo(3, "ThirdTodo", true),
	}

	todoManager := &TodoManager{todos}

	todo := models.NewTodo(4, "Fourth todo", false)
	todoManager.Add(todo)

	got, err := todoManager.GetById(todo.Id)

	utils.AssertNoError(t, err)
	utils.AssertTodosEqual(t, got, &todo)
}
