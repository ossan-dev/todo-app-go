package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/src/crud"
)

type InMemoryTodoStore struct{}

func (i *InMemoryTodoStore) GetTodoById(id int) string {
	return "FirstTodo"
}

func main() {
	server := &crud.TodoServer{&InMemoryTodoStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
