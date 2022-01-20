package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/src/crud/servers"
)

func main() {
	store := database.NewInMemoryTodoStore()
	server := servers.NewTodoServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
