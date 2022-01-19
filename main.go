package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/src/crud"
	"todo-app-go.com/v1/src/models"
)

type InMemoryTodoStore struct{}

func (i *InMemoryTodoStore) GetTodoById(id int) *models.Todo {
	return &models.Todo{1, "FirstTodo", false}
}

func main() {
	server := &crud.TodoServer{&InMemoryTodoStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
