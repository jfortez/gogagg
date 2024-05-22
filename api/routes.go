package api

import (
	"net/http"

	"github.com/jfortez/gogagg/api/handler"
)

func GetRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /todos", handler.TodosHandle)
	router.HandleFunc("GET /todos/{id}", handler.TodoHandle)

	router.HandleFunc("GET /users", handler.UsersHandle)
	router.HandleFunc("GET /users/{id}", handler.UserHandle)
	router.HandleFunc("POST /users", handler.UpdateUserHandle)
	router.HandleFunc("DELETE /users/{id}", handler.DeleteUserHandle)
}
