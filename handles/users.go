package handles

import (
	"encoding/json"
	"net/http"

	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
)

func UsersHandle(w http.ResponseWriter, r *http.Request) {

	connection, ok := r.Context().Value("db").(*db.DataBase)
	if !ok {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}

	userList := services.GetUsers(connection.Connection)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)
}

func UserHandle(w http.ResponseWriter, r *http.Request) {
	connection, ok := r.Context().Value("db").(*db.DataBase)
	if !ok {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}
	id := r.PathValue("id")
	userList := services.GetUser(connection.Connection, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)

}
func DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	connection, ok := r.Context().Value("db").(*db.DataBase)
	if !ok {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	services.DeleteUser(connection.Connection, id)
}

func UpdateUserHandle(w http.ResponseWriter, r *http.Request) {

	connection, ok := r.Context().Value("db").(*db.DataBase)
	if !ok {
		http.Error(w, "Database connection not found", http.StatusInternalServerError)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	services.UpdateUser(connection.Connection, user)
}
