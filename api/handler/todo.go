package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
)

const url = "https://jsonplaceholder.typicode.com/todos/"

func TodosHandle(w http.ResponseWriter, _ *http.Request) {

	jsonData := services.GetTodo[[]model.Todo](url)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func TodoHandle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)

	jsonData := services.GetTodo[model.Todo](url + id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}
