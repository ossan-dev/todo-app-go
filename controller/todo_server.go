package controller

import (
	"encoding/json"
	"io/ioutil"
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
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(t.todoStore.GetAllTodos())
	w.WriteHeader(http.StatusOK)
}

func (t *TodoServer) getAllTodos() []model.Todo {
	return []model.Todo{
		model.NewTodo(1, "First Todo", false),
		model.NewTodo(2, "Second Todo", true),
	}
}

func (t *TodoServer) todoCreation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo model.Todo
	json.Unmarshal(reqBody, &todo)
	t.todoStore.AddTodo(todo)
}
