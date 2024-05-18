package handles

import (
	"encoding/json"
	"net/http"
	"taco/db"
	"taco/model"
	"taco/services"
)

func UsersHandle(w http.ResponseWriter, _ *http.Request) {
	userList := services.GetUsers(db.Pool)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)
}

func UserHandle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userList := services.GetUser(db.Pool, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)

}
func DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.WriteHeader(http.StatusOK)
	services.DeleteUser(db.Pool, id)
}

func UpdateUserHandle(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	services.UpdateUser(db.Pool, user)
}
