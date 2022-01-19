package main

import (
	"log"
	"net/http"

	"todo-app-go.com/v1/src/crud/servers"
	"todo-app-go.com/v1/src/crud/stores"
)

func main() {
	store := stores.NewInMemoryTodoStore()
	server := servers.NewTodoServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
