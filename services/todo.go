package services

import (
	"encoding/json"
	"io"
	"net/http"
	"taco/model"
)

func GetTodo[K model.Todo | []model.Todo](url string) (jsonData K) {
	result, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		panic(err)
	}
	return
}
