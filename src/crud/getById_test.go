package crud

import (
	"fmt"
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

type StubTodoStore struct {
	todos map[int]string
}

func (s *StubTodoStore) GetTodoById(id int) string {
	todo := s.todos[id]
	return todo
}

func TestGetByIdEndpoint(t *testing.T) {
	store := StubTodoStore{
		map[int]string{
			1: `{"id": 1, "description": "FirstTodo", "isCompleted": false}`,
			2: `{"id": 2, "description": "SecondTodo", "isCompleted": true}`,
		},
	}
	server := &TodoServer{&store}

	t.Run("first todo", func(t *testing.T) {
		request := newGetToDoByIdReq(1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), `{"id": 1, "description": "FirstTodo", "isCompleted": false}`)
	})

	t.Run("second todo", func(t *testing.T) {
		request := newGetToDoByIdReq(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utils.AssertResponseBody(t, response.Body.String(), `{"id": 2, "description": "SecondTodo", "isCompleted": true}`)
	})
}

func newGetToDoByIdReq(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/todos/%d", id), nil)
	return req
}
