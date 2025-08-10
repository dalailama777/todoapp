package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"log"

	_ "todo-app/docs"

	_ "github.com/mattn/go-sqlite3"

	. "todo-app/internal/handler"
	. "todo-app/internal/middleware"
	. "todo-app/internal/repository"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Task представляет задачу

// @title Todo App API
// @version 1.0
// @description API для управления задачами
// @host localhost:8080
// @BasePath /

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	err = InitTable()
	if err != nil {
		log.Fatal(err)
	}
	r.HandleFunc("/tasks", GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks", CreateTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", GetTaskByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
