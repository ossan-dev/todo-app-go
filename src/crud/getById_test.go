package crud

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetById(t *testing.T) {
	todos := []models.Todo{
		models.NewTodo(1, "FirstTodo", false),
		models.NewTodo(2, "SecondTodo", false),
		models.NewTodo(3, "ThirdTodo", true),
	}

	todoManager := &TodoManager{todos}

	t.Run("todo present in collection", func(t *testing.T) {
		got, err := todoManager.GetById(1)

		utils.AssertTodosEqual(t, got, &todos[0])
		utils.AssertNoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := todoManager.GetById(4)

		utils.AssertTodosEqual(t, got, nil)
		utils.AssertError(t, err, ErrTodoNotFound)
	})
}

func TestGetByIdEndpoint(t *testing.T) {
	t.Run("first todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/1", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `{"id": 1, "description": "FirstTodo", "isCompleted": false}`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("first todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/2", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := `{"id": 2, "description": "SecondTodo", "isCompleted": true}`

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
