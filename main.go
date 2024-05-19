package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/jfortez/gogagg/db"
)

type contextHandler struct {
	ctx context.Context
	h   http.Handler
}

func (c contextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.h.ServeHTTP(w, r.WithContext(c.ctx))
}

func handle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./public/index.html"))
	tpl.Execute(w, nil)
}

var dbKey = "db"

func main() {

	conn := db.New()

	defer conn.Close()

	ctx := context.WithValue(context.Background(), dbKey, conn)

	router := http.NewServeMux()

	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handle)

	GetRoutes(router)

	srv := &http.Server{
		Handler:      Logging(contextHandler{ctx, router}),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}
