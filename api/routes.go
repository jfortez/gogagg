package api

import (
	"database/sql"
	"net/http"

	"github.com/jfortez/gogagg/api/handler"
)

func GetRoutes(router *http.ServeMux, db *sql.DB) {

	userHandles := handler.NewUserHandler(db)

	router.HandleFunc("GET /api/v1/todos", handler.TodosHandle)
	router.HandleFunc("GET /api/v1/todos/{id}", handler.TodoHandle)

	router.HandleFunc("GET /api/v1/users", userHandles.GetUsersHandler)
	router.HandleFunc("GET /api/v1/users/create", userHandles.CreateUserHandler)
	router.HandleFunc("GET /api/v1/users/{id}", userHandles.GetUserHandler)
	router.HandleFunc("POST /api/v1/users", userHandles.DeleteUserHandler)
	router.HandleFunc("DELETE /api/v1/users/{id}", userHandles.UpdateUserHandler)
}
