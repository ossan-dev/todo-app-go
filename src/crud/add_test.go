package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestAdd(t *testing.T) {
	todos := []models.Todo{
		{Id: 1, Description: "FirstTodo", IsCompleted: false},
		{Id: 2, Description: "SecondTodo", IsCompleted: false},
		{Id: 3, Description: "ThirdTodo", IsCompleted: false},
	}

	todoManager := &TodoManager{todos}

	todo := models.Todo{
		Id:          4,
		Description: "Fourth todo",
		IsCompleted: false,
	}
	todoManager.Add(todo)

	got, err := todoManager.GetById(todo.Id)

	utils.AssertNoError(t, err)
	utils.AssertTodosEqual(t, got, &todo)
}
