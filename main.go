package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/middleware"
	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
)

const dbKey = "db"

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

	router.HandleFunc("POST /create", handleCreate)
	router.HandleFunc("DELETE /remove/{id}", handleRemove)

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

func handle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./public/index.html"))
	userList := services.GetUsers(r.Context().Value(dbKey).(*db.DataBase).Connection)

	tpl.Execute(w, userList)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	Name := r.PostFormValue("name")
	Age := r.PostFormValue("age")
	Email := r.PostFormValue("email")
	Image := "https://cdn-icons-png.freepik.com/512/6596/6596121.png"

	tpl := template.Must(template.ParseFiles("./public/index.html"))

	AgeInt, _ := strconv.Atoi(Age)

	currentUser := model.User{
		Name:  Name,
		Age:   AgeInt,
		Email: Email,
		Img:   Image,
	}

	db := r.Context().Value(dbKey).(*db.DataBase).Connection
	services.CreateUser(db, currentUser)

	tpl.ExecuteTemplate(w, "user-element", currentUser)
}

func handleRemove(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	db := r.Context().Value(dbKey).(*db.DataBase).Connection

	services.DeleteUser(db, id)

	w.WriteHeader(http.StatusOK)

}
