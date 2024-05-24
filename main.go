package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jfortez/gogagg/api"
	"github.com/jfortez/gogagg/api/middleware"
	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
	"github.com/jfortez/gogagg/web/templates"
	"github.com/joho/godotenv"
)

const dbKey = "db"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ADDRESS := os.Getenv("ADDRESS")
	dbConn := db.New()
	defer dbConn.Close()
	dbConn.InitDB()

	ctx := context.WithValue(context.Background(), dbKey, dbConn.Connection)

	contextMiddleware := middleware.NewContextHandler(ctx)
	middlewares := middleware.Chain(middleware.Logging, contextMiddleware)

	router := http.NewServeMux()

	// STATIC
	dir := http.Dir("./web/static")
	fs := http.FileServer(dir)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// WEB
	router.HandleFunc("/", http.HandlerFunc(handleUserView))
	router.HandleFunc("/todos", handleTodosView)

	router.HandleFunc("POST /create", handleCreate)
	router.HandleFunc("DELETE /remove/{id}", handleRemove)

	// API
	api.GetRoutes(router, dbConn.Connection)

	srv := &http.Server{
		Handler:      middlewares(router),
		Addr:         ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}

func handleTodosView(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todos"))
}

func handleUserView(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers(r.Context().Value(dbKey).(*sql.DB))
	component := templates.Hello(users)
	component.Render(r.Context(), w)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	Name := r.PostFormValue("name")
	Age := r.PostFormValue("age")
	Email := r.PostFormValue("email")
	Image := "https://cdn-icons-png.freepik.com/512/6596/6596121.png"

	AgeInt, _ := strconv.Atoi(Age)

	currentUser := model.User{
		Name:  Name,
		Age:   AgeInt,
		Email: Email,
		Img:   Image,
	}

	db := r.Context().Value(dbKey).(*sql.DB)
	services.CreateUser(db, currentUser)

}

func handleRemove(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	db := r.Context().Value(dbKey).(*sql.DB)

	services.DeleteUser(db, id)

	w.WriteHeader(http.StatusOK)

}
