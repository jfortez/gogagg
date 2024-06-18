package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
}

func NewService(address string, storage *db.Storage, wsHub *services.Hub) *Service {
	return &Service{
		address: address,
		storage: storage.DB,
		wsHub:   wsHub,
	}
}

func (s *Service) generateToken() string {
	// generate a random token
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(token)
}

func (s *Service) Run() {
	router := http.NewServeMux()

	apiRest := NewAPIRoutes(router, s.storage)

	// STATIC

	dir := http.Dir("./web/static")
	fs := http.FileServer(dir)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// WEB
	router.HandleFunc("/", handleProtectedRoute(s.handleChatView))

	router.HandleFunc("/login", handlePublicRoute(handleLoginView))
	router.HandleFunc("/register", handlePublicRoute(handleRegisterView))
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

	messages, err := services.GetMessages(s.storage, vals.UserId, 1)
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

	chatContent := templates.ChatContent(messages, currentUser)
	chatContent.Render(r.Context(), w)

}

func (s *Service) handleSendMessage(w http.ResponseWriter, r *http.Request) {
	requestedMessage := model.RequestMessage{}

	err := json.NewDecoder(r.Body).Decode(&requestedMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMessage := model.CreateMessage{
		Content:    requestedMessage.Content,
		Status:     "delivered",
		FromUserId: 1,
		ToUserId:   requestedMessage.UserId,
	}
	err = services.SendMessage(s.storage, newMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatItem := templates.MessageItem(model.ChatMessage{UserId: newMessage.FromUserId, UserName: "John Doe", MessageContent: newMessage.Content, MessageStatus: newMessage.Status})

	chatItem.Render(r.Context(), w)

}
func handleProtectedRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// validate token
		if token.Value != "" {
			handler(w, r)
		}

	}
}

func handlePublicRoute(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")

		if err != nil {
			handler(w, r)
			return
		}

		// validate token
		if token.Value != "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
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

	chatList, err := services.GetMessageListByCurrentUser(s.storage, 1)
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

	fmt.Println(user)
	cookie := &http.Cookie{
		// Name:     "user",
		// Value:    strconv.Itoa(user.Id),
		Name:     "token",
		Value:    s.generateToken(),
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 365),
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
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
