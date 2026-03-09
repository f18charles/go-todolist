package main

import (
	"log"
	"net/http"

	"go-todolist/internal/database"
	"go-todolist/internal/handler"
	"go-todolist/internal/repository"
	"go-todolist/internal/router"
	"go-todolist/internal/service"
)

func main() {
	collection, err := database.NewMongoDB("mongodb://localhost:27017", "todolist", "todos")
	if err != nil {
    	log.Fatal(err)
	}

	repo := repository.NewMongoRepo(collection)
	svc := service.NewTodoService(repo)
	h := handler.NewStdHandler(svc)
	r := router.NewStdRouter(h)

	log.Println("server running on :8080")
	http.ListenAndServe(":8080", r)
}
