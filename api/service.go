package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jfortez/gogagg/api/middleware"
	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
	"github.com/jfortez/gogagg/web/templates"
)

type Service struct {
	storage *sql.DB
	wsHub   *services.Hub
	address string
	session *Session
}

const duration = time.Hour * 24

func NewService(address string, storage *db.Storage, wsHub *services.Hub) *Service {
	session := NewSession(duration, []byte(os.Getenv("SECRET_KEY")))
	return &Service{
		address: address,
		storage: storage.DB,
		wsHub:   wsHub,
		session: session,
	}
}

func (s *Service) Run() {
	router := http.NewServeMux()

	apiRest := NewAPIRoutes(router, s.storage)

	// STATIC

	dir := http.Dir("./web/static")
	fs := http.FileServer(dir)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// WEB
	router.HandleFunc("/", s.handleProtectedRoute(s.handleChatView))

	router.HandleFunc("/login", s.handlePublicRoute(handleLoginView))
	router.HandleFunc("/register", s.handlePublicRoute(handleRegisterView))
	router.HandleFunc("/ws", s.handleWs)

	router.HandleFunc("POST /sendMessage", s.handleSendMessage)
	router.HandleFunc("POST /message", s.handleMessage)
	router.HandleFunc("POST /login", s.handleLogin)
	router.HandleFunc("POST /register", s.handleRegister)

	// API
	apiRest.Run()

	srv := &http.Server{
		Handler:      middleware.Logging(router),
		Addr:         s.address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server Listening on port %s\n", s.address)
	log.Fatal(srv.ListenAndServe())
}

func (s *Service) handleMessage(w http.ResponseWriter, r *http.Request) {

	vals := model.RequestedMessages{}
	err := json.NewDecoder(r.Body).Decode(&vals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loggedUser, err := s.session.getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages, err := services.GetMessages(s.storage, vals.UserId, loggedUser.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUser := model.CurrentChatUser{
		UserId:              vals.UserId,
		UserName:            vals.UserName,
		LastInteractionTime: vals.UpdatedAt,
		Avatar:              vals.UserAvatar,
	}

	fmt.Println(currentUser)

	chatContent := templates.ChatContent(messages, currentUser, loggedUser)
	chatContent.Render(r.Context(), w)

}

func (s *Service) handleSendMessage(w http.ResponseWriter, r *http.Request) {

	requestedMessage := model.RequestMessage{}

	err := json.NewDecoder(r.Body).Decode(&requestedMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loggedUser, err := s.session.getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newMessage := model.CreateMessage{
		Content:    requestedMessage.Content,
		Status:     "delivered",
		FromUserId: loggedUser.Id,
		ToUserId:   requestedMessage.UserId,
	}
	err = services.SendMessage(s.storage, newMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newMessageItem := model.ChatMessage{
		UserId:           loggedUser.Id,
		UserName:         loggedUser.Name,
		MessageContent:   newMessage.Content,
		MessageStatus:    newMessage.Status,
		Avatar:           loggedUser.Avatar,
		MessageCreatedAt: time.Now(),
		MessageUpdatedAt: time.Now(),
	}

	chatItem := templates.MessageItem(newMessageItem, loggedUser)

	chatItem.Render(r.Context(), w)

}
func (s *Service) handleProtectedRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		err = s.session.VerifyToken(token.Value)

		if err != nil {

			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		handler(w, r)

	}
}

func (s *Service) handlePublicRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		if err != nil {
			handler(w, r)
			return
		}

		err = s.session.VerifyToken(token.Value)

		if err != nil {

			s.session.removeSession(w, r)
			handler(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)

	}
}
func handleLoginView(w http.ResponseWriter, r *http.Request) {
	templates.Login().Render(r.Context(), w)
}
func handleRegisterView(w http.ResponseWriter, r *http.Request) {
	templates.Register().Render(r.Context(), w)
}

func (s *Service) handleChatView(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	userLogged, err := s.session.getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatList, err := services.GetMessageListByCurrentUser(s.storage, userLogged.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	component := templates.Index(chatList)
	component.Render(r.Context(), w)
}

func (s *Service) handleLogin(w http.ResponseWriter, r *http.Request) {

	var requestUser model.User
	err := json.NewDecoder(r.Body).Decode(&requestUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.FindAuthUser(s.storage, requestUser.Email, requestUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.session.setSession(w, r, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	message := map[string]string{
		"message": "User logged in successfully",
	}
	json.NewEncoder(w).Encode(message)

}

func (s *Service) handleRegister(w http.ResponseWriter, r *http.Request) {
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

	if requestUser.Avatar == "" {
		requestUser.Avatar = "https://cdn-icons-png.freepik.com/512/6596/6596121.png"
	}

	err = services.CreateUser(s.storage, model.User{
		Name:        requestUser.Name,
		Age:         requestUser.Age,
		Email:       requestUser.Email,
		Avatar:      requestUser.Avatar,
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

func (s *Service) handleWs(w http.ResponseWriter, r *http.Request) {
	services.ServeWs(s.wsHub, w, r)
}
