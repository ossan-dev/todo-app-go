package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetByStatus(t *testing.T) {
	t.Run("get completed todos return non-empty collections", func(t *testing.T) {
		todos := []models.Todo{
			models.NewTodo(1, "FirstTodo", false),
			models.NewTodo(2, "SecondTodo", false),
			models.NewTodo(3, "ThirdTodo", true),
		}

		todoManager := &TodoManager{todos}

		got := todoManager.GetByStatus(true)
		want := []models.Todo{
			models.NewTodo(3, "ThirdTodo", true),
		}
		utils.AssertCollectionsEqual(t, got, want)
	})

	t.Run("get completed todos return empty collections", func(t *testing.T) {
		todos := []models.Todo{
			models.NewTodo(1, "FirstTodo", false),
			models.NewTodo(2, "SecondTodo", false),
			models.NewTodo(3, "ThirdTodo", false),
		}

		todoManager := &TodoManager{todos}

		got := todoManager.GetByStatus(true)
		utils.AssertCollectionsEqual(t, got, []models.Todo{})
	})
}
