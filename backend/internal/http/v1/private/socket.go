package private

import (
	"log"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (h *PrivateHandler) initSocketRoutes(root fiber.Router) {
	socketRoute := root.Group("/socket")
	socketRoute.Get("/ws", h.handleWebSocketConnection)
}

// WebSocket Connection
func (h *PrivateHandler) handleWebSocketConnection(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return websocket.New(h.handleWebSocket)(c)
	}

	return fiber.ErrUpgradeRequired
}

// @Tags WebSocket
// @Summary Establish a WebSocket Connection
// @Description Initiates a WebSocket connection for real-time communication. Clients can send and receive messages.
// @Success 101 {object} string "Connection Established"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /private/socket/ws [get]
func (h *PrivateHandler) handleWebSocket(c *websocket.Conn) {
	/* 	user := c.Locals("user")
	   	if user == nil {
	   		return
	   	}
	   	session_data, ok := user.(session_store.SessionData)
	   	if !ok {
	   		return
	   	}
	   	userID := session_data.UserID */
	userID := "b05ca195-c0a9-4ac9-905d-2409962b26bd" // This is for test

	newClient, err := domains.NewClient(userID, c)
	if err != nil {
		return
	}

	// Client'ı eklemek için mutex ile kilitleyip map'e ekliyoruz
	h.mu.Lock()
	h.clients[newClient] = true
	h.mu.Unlock()

	// Bağlantı kapandığında client'ı map'ten çıkarmak için defer ile tanımlıyoruz
	defer func() {
		h.mu.Lock()
		delete(h.clients, newClient)
		h.mu.Unlock()

		c.Close()
		log.Printf("WebSocket Connection Closed: %s", userID)
	}()

	log.Printf("New WebSocket Connection: %s", userID)
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}
