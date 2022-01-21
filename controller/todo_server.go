package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"todo-app-go.com/v1/database"
)

type TodoServer struct {
	todoStore database.TodoStore
	http.Handler
}

func NewTodoServer(todoStore database.TodoStore) *TodoServer {
	server := new(TodoServer)

	server.todoStore = todoStore

	router := mux.NewRouter()

	controller := &TodoController{todoStore}

	router.HandleFunc("/api/todos", controller.AddTodo).Methods(http.MethodPost)
	router.HandleFunc("/api/todos", controller.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc("/api/todos/{id:[0-9]+}", controller.GetTodoById).Methods(http.MethodGet)

	server.Handler = router

	return server
}
