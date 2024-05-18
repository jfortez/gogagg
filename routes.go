package main

import (
	"net/http"
	"taco/handles"
)

func GetRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /todos", handles.TodosHandle)
	router.HandleFunc("GET /todos/{id}", handles.TodoHandle)

	router.HandleFunc("GET /users", handles.UsersHandle)
	router.HandleFunc("GET /users/{id}", handles.UserHandle)
	router.HandleFunc("POST /users", handles.UpdateUserHandle)
	router.HandleFunc("DELETE /users/{id}", handles.DeleteUserHandle)
}
