package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/src/crud"
)

func main() {
	server := &crud.TodoServer{&crud.InMemoryTodoStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
