package domains

import (
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type Client struct {
	userID     uuid.UUID
	connection *websocket.Conn
}

type Response struct {
	Type string
	Data any
}

func NewClient(
	userID string,
	connection *websocket.Conn,
) (*Client, error) {
	var newClient Client

	if err := newClient.SetUserID(userID); err != nil {
		return nil, err
	}
	newClient.SetConnection(connection)

	return &newClient, nil
}

// Getters & Setters
func (c *Client) SetUserID(userID string) error {
	UUIDuserID, err := uuid.Parse(userID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "Invalid user id", err)
	}
	c.userID = UUIDuserID

	return nil
}

func (c *Client) GetUserID() uuid.UUID {
	return c.userID
}

func (c *Client) SetConnection(conn *websocket.Conn) {
	c.connection = conn
}

func (c *Client) GetConnection() *websocket.Conn {
	return c.connection
}
