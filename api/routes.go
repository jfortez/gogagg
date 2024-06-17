package api

import (
	"database/sql"
	"net/http"

	"github.com/jfortez/gogagg/api/handler"
)

type APIRoutes struct {
	router  *http.ServeMux
	storage *sql.DB
}

func NewAPIRoutes(router *http.ServeMux, db *sql.DB) *APIRoutes {
	return &APIRoutes{router: router, storage: db}
}

func (r *APIRoutes) Run() {

	userHandles := handler.NewUserHandler(r.storage)
	router := r.router

	router.HandleFunc("GET /api/v1/todos", handler.TodosHandle)
	router.HandleFunc("GET /api/v1/todos/{id}", handler.TodoHandle)

	router.HandleFunc("GET /api/v1/users", userHandles.GetUsersHandler)
	router.HandleFunc("GET /api/v1/users/create", userHandles.CreateUserHandler)
	router.HandleFunc("GET /api/v1/users/{id}", userHandles.GetUserHandler)
	router.HandleFunc("POST /api/v1/users", userHandles.DeleteUserHandler)
	router.HandleFunc("DELETE /api/v1/users/{id}", userHandles.UpdateUserHandler)
}
