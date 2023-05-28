package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// User represents a connected user.
type User struct {
	UUID     string
	Conn     *websocket.Conn
	Messages chan string
}

// ConnectionManager handles WebSocket connections.
type ConnectionManager struct {
	Users      map[string]*User
	Register   chan *User
	Unregister chan *User
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

/*// ServeHTTP handles the HTTP requests and upgrades the connection to WebSocket.
func (c *ConnectionManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	uuid := r.PostFormValue("uuid")
	if uuid == "" {
		log.Println("No UUID provided.")
		conn.Close()
		return
	}

	user := &User{
		UUID:     uuid,
		Conn:     conn,
		Messages: make(chan string),
	}

	c.Register <- user

	// Start goroutine to listen for incoming messages from the user's WebSocket connection
	go c.handleIncomingMessages(user)

	// Start goroutine to send messages to the user's WebSocket connection
	go c.handleOutgoingMessages(user)
}*/

// WebSocketHandler is the handler function for the WebSocket route.
func (c *ConnectionManager) WebSocketHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	uuid := ctx.PostForm("uuid")
	if uuid == "" {
		log.Println("No UUID provided.")
		conn.Close()
		return
	}

	user := &User{
		UUID:     uuid,
		Conn:     conn,
		Messages: make(chan string),
	}

	c.Register <- user

	// Start goroutine to listen for incoming messages from the user's WebSocket connection
	go c.handleIncomingMessages(user)

	// Start goroutine to send messages to the user's WebSocket connection
	go c.handleOutgoingMessages(user)
}

// handleIncomingMessages listens for incoming messages from a user's WebSocket connection.
func (c *ConnectionManager) handleIncomingMessages(user *User) {
	defer func() {
		c.Unregister <- user
		user.Conn.Close()
	}()

	for {
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			break
		}

		// Process the received message as needed
		fmt.Printf("Received message from user %s: %s\n", user.UUID, string(msg))
	}
}

// handleOutgoingMessages sends messages to a user's WebSocket connection.
func (c *ConnectionManager) handleOutgoingMessages(user *User) {
	for message := range user.Messages {
		err := user.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			break
		}
	}
}

var WsConnectionManager *ConnectionManager

func InitChatServices() {
	WsConnectionManager = &ConnectionManager{
		Users:      make(map[string]*User),
		Register:   make(chan *User),
		Unregister: make(chan *User),
	}

	go WsConnectionManager.StartUserManagement()

	// todo 需要在外面转发至此
	//http.Handle("/ws", WsConnectionManager)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("Failed to start server:", err)
	//}
}

// StartUserManagement starts the user management goroutine.
func (c *ConnectionManager) StartUserManagement() {

	for {
		select {
		case user := <-c.Register:
			c.Users[user.UUID] = user
			log.Printf("User %s connected.\n", user.UUID)

		case user := <-c.Unregister:
			delete(c.Users, user.UUID)
			close(user.Messages)
			log.Printf("User %s disconnected.\n", user.UUID)

		}
	}
}

func IsUserOnline(uuid string) bool {
	_, ok := WsConnectionManager.Users[uuid]
	return ok
}
