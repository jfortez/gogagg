package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taco/model"
	"taco/services"
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
