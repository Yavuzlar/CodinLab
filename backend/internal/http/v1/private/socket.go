package private

import (
	"log"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
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
	user := c.Locals("user")
	if user == nil {
		return
	}
	session_data, ok := user.(session_store.SessionData)
	if !ok {
		return
	}
	userID := session_data.UserID
	//userID := "b05ca195-c0a9-4ac9-905d-2409962b26bd" // This is for test

	newClient, err := domains.NewClient(userID, c)
	if err != nil {
		return
	}

	// To add the client, we lock it with mutex and add it to the map.
	h.mu.Lock()
	h.clients[newClient] = true
	h.mu.Unlock()

	// We define defer to remove the client from the map when the connection is closed.
	defer func() {
		h.mu.Lock()
		delete(h.clients, newClient)
		h.mu.Unlock()

		c.Close()
		log.Printf("WebSocket Connection Closed: %s", userID)
	}()

	log.Printf("New WebSocket Connection: %s", userID)
	for {
		_, messages, err := c.ReadMessage()
		if err != nil {
			break
		}
		if err = h.services.CodeService.SaveUserHistory(c, messages, userID); err != nil {
			c.WriteJSON(domains.Response{
				Type: "close",
				Data: struct {
					Status  int    `json:"status"`
					Message string `json:"message"`
				}{
					Status:  400,
					Message: err.Error(),
				},
			})
		}

	}
}
