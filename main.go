package main

import (
	"fmt"
	"log"
	"net/http"
	"taco/db"
	"text/template"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./public/index.html"))
	tpl.Execute(w, nil)
}

func main() {

	db.InitDB()

	router := http.NewServeMux()

	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handle)

	GetRoutes(router)

	srv := &http.Server{
		Handler:      Logging(router),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}
