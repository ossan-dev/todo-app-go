package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetByStatus(t *testing.T) {
	t.Run("get completed todos return non-empty collections", func(t *testing.T) {
		todos := []models.Todo{
			{Id: 1, Description: "FirstTodo", IsCompleted: false},
			{Id: 3, Description: "ThirdTodo", IsCompleted: true},
			{Id: 2, Description: "SecondTodo", IsCompleted: false},
		}

		todoManager := &TodoManager{todos}

		got := todoManager.GetByStatus(true)
		want := []models.Todo{
			{Id: 3, Description: "ThirdTodo", IsCompleted: true},
		}
		utils.AssertCollectionsEqual(t, got, want)
	})

	t.Run("get completed todos return empty collections", func(t *testing.T) {
		todos := []models.Todo{
			{Id: 1, Description: "FirstTodo", IsCompleted: false},
			{Id: 3, Description: "ThirdTodo", IsCompleted: false},
			{Id: 2, Description: "SecondTodo", IsCompleted: false},
		}

		todoManager := &TodoManager{todos}

		got := todoManager.GetByStatus(true)
		utils.AssertCollectionsEqual(t, got, []models.Todo{})
	})
}
