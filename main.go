package main

import (
	"context"
	"database/sql"
	"encoding/json"
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
	router.HandleFunc("/login", http.HandlerFunc(handleLoginView))
	router.HandleFunc("/register", http.HandlerFunc(handleRegisterView))

	router.HandleFunc("POST /login", handleLogin)
	router.HandleFunc("POST /register", handleRegister)

	// API
	api.GetRoutes(router, dbConn.Connection)

	srv := &http.Server{
		Handler:      middlewares(router),
		Addr:         ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer srv.Close()
	fmt.Println("Server listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}

func handleLoginView(w http.ResponseWriter, r *http.Request) {
	templates.Login().Render(r.Context(), w)
}
func handleRegisterView(w http.ResponseWriter, r *http.Request) {
	templates.Register().Render(r.Context(), w)
}

func handleUserView(w http.ResponseWriter, r *http.Request) {
	component := templates.Index()
	component.Render(r.Context(), w)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	var requestUser model.User
	err := json.NewDecoder(r.Body).Decode(&requestUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := r.Context().Value(dbKey).(*sql.DB)

	user, err := services.FindAuthUser(db, requestUser.Email, requestUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "user",
		Value:    strconv.Itoa(user.Id),
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 365),
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Redirect", "/")
	message := map[string]string{
		"message": "User logged in successfully",
	}
	json.NewEncoder(w).Encode(message)

}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	var requestUser struct {
		model.User
		ConfirmPassword string `json:"confirmPassword"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestUser.Password != requestUser.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	db := r.Context().Value(dbKey).(*sql.DB)

	if requestUser.Img == "" {
		requestUser.Img = "https://cdn-icons-png.freepik.com/512/6596/6596121.png"
	}

	err = services.CreateUser(db, model.User{
		Name:        requestUser.Name,
		Age:         requestUser.Age,
		Email:       requestUser.Email,
		Img:         requestUser.Img,
		Password:    requestUser.Password,
		Description: requestUser.Description,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("HX-Redirect", "/")

	message := map[string]string{
		"message": "User created successfully",
	}
	json.NewEncoder(w).Encode(message)
}
