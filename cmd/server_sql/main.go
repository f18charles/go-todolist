package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"go-todolist/internal/handler"
	"go-todolist/internal/repository"
	"go-todolist/internal/router"
	"go-todolist/internal/service"
	"go-todolist/internal/database"
)

func main() {
	db, err := database.NewSQLiteDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := database.InitDB(db); err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}

	repo := repository.NewSQLiteRepo(db)
	svc := service.NewTodoService(repo)
	h := handler.NewChiHandler(svc)
	r := router.NewChiRouter(h)

	log.Println("server running on :8080")
	http.ListenAndServe(":8080", r)
}
