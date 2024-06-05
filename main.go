package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jfortez/gogagg/api"
	"github.com/jfortez/gogagg/api/middleware"
	"github.com/jfortez/gogagg/db"
	"github.com/jfortez/gogagg/model"
	"github.com/jfortez/gogagg/services"
	"github.com/jfortez/gogagg/web/templates"
	_ "github.com/joho/godotenv/autoload"
)

const dbKey = "db"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
func main() {

	// ADDRESS := os.Getenv("ADDRESS")
	dbConn := db.New()
	defer dbConn.Close()
	go dbConn.InitDB()

	hub := newHub()
	go hub.run()

	ctx := context.WithValue(context.Background(), dbKey, dbConn.Connection)

	contextMiddleware := middleware.NewContextHandler(ctx)
	middlewares := middleware.Chain(middleware.Logging, contextMiddleware)

	router := http.NewServeMux()

	// STATIC
	dir := http.Dir("./web/static")
	fs := http.FileServer(dir)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// WEB
	router.HandleFunc("/", http.HandlerFunc(handleChatView))
	// router.HandleFunc("/message/{fromUserId}", http.HandlerFunc(handleMessageView))
	router.HandleFunc("/login", http.HandlerFunc(handleLoginView))
	router.HandleFunc("/register", http.HandlerFunc(handleRegisterView))
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	router.HandleFunc("POST /login", handleLogin)
	router.HandleFunc("POST /register", handleRegister)

	// API
	api.GetRoutes(router, dbConn.Connection)

	srv := &http.Server{
		Handler:      middlewares(router),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// defer srv.Close()

	fmt.Println("Server Listening on localhost:8000")
	log.Fatal(srv.ListenAndServe())

}

func handleLoginView(w http.ResponseWriter, r *http.Request) {
	templates.Login().Render(r.Context(), w)
}
func handleRegisterView(w http.ResponseWriter, r *http.Request) {
	templates.Register().Render(r.Context(), w)
}

func handleChatView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	db := r.Context().Value(dbKey).(*sql.DB)
	chatList, err := services.GetMessageListByCurrentUser(db, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	component := templates.Index(chatList)
	component.Render(r.Context(), w)
}

// func handleMessageView(w http.ResponseWriter, r *http.Request) {
// 	id,err := strconv.Atoi(r.PathValue("fromUserId"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	db := r.Context().Value(dbKey).(*sql.DB)
// 	message, err := services.GetMessages(db, id, 1)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	component := templates.Index(message)
// 	component.Render(r.Context(), w)
// }

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
