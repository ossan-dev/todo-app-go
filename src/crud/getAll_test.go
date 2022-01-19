package crud

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetAll(t *testing.T) {
	t.Run("GetAllTodos should return all the todos added", func(t *testing.T) {
		todos := []models.Todo{
			models.NewTodo(1, "FirstTodo", false),
			models.NewTodo(2, "SecondTodo", false),
			models.NewTodo(3, "ThirdTodo", true),
		}

		todoManager := &TodoManager{todos}

		utils.AssertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})

	t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
		todos := []models.Todo{}

		todoManager := &TodoManager{todos}

		utils.AssertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	})
}

func TestGetAllEndpoint(t *testing.T) {
	t.Run("returns todos", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "api/todos", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `[{"id": 1, "description": "FirstTodo", "isCompleted": false}]`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
