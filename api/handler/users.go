package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
)

type userHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{db: db}
}

func (h *userHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	if h.db == nil {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}

	userList := services.GetUsers(h.db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)
}

func (h *userHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	services.CreateUser(h.db, user)

	w.WriteHeader(http.StatusOK)

}

func (h *userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

	if h.db == nil {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}
	id := r.PathValue("id")
	userList := services.GetUser(h.db, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)

}
func (h *userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if h.db == nil {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	services.DeleteUser(h.db, id)
}

func (h *userHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	if h.db == nil {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	services.UpdateUser(h.db, user)
}
