package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"todolist/internal/handler"
	"todolist/internal/repository"
	"todolist/internal/router"
	"todolist/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewSQLiteRepo(db)
	svc := service.NewTodoService(repo)
	h := handler.NewStdHandler(svc)
	r := router.NewStdRouter(h)

	log.Println("server running on :8080")
	http.ListenAndServe(":8080", r)
}
