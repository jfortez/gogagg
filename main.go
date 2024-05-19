package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/middleware"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./public/index.html"))
	tpl.Execute(w, nil)
}

var dbKey = "db"

func main() {

	conn := db.New()

	defer conn.Close()

	ctx := context.WithValue(context.Background(), dbKey, conn)

	contextMiddleware := middleware.NewContextHandler(ctx)
	middlewares := middleware.Chain(middleware.Logging, contextMiddleware)

	router := http.NewServeMux()

	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handle)

	GetRoutes(router)

	srv := &http.Server{
		Handler:      middlewares(router),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}
