package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/model"
)

type TodoServer struct {
	todoStore database.TodoStore
	http.Handler
}

func NewTodoServer(todoStore database.TodoStore) *TodoServer {
	server := new(TodoServer)

	server.todoStore = todoStore

	router := http.NewServeMux()
	router.Handle("/api/todos", http.HandlerFunc(server.todosHandler))

	server.Handler = router

	return server
}

func (t *TodoServer) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(t.todoStore.GetAllTodos())
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		var todo model.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			panic(fmt.Sprintf("invalid json data, err: %v", err))
		}
		t.todoStore.AddTodo(todo)
		w.WriteHeader(http.StatusCreated)
	}
}
