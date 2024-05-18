package private

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	root.Get("/log", h.GetAllLogs)
	root.Get("/log/:id", h.GetLogByID)
}

type LogDTO struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	Title   string    `json:"title"`
	LType   string    `json:"type"`
	Content string    `json:"content"`
}

// @Tags Log
// @Summary Get all logs
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param userID query string false "User ID"
// @Param title query string false "Log Title"
// @Param content query string false "Log Content"
// @Param type query string false "Log Type"
// @Success 200 {object} response.BaseResponse{data=[]LogDTO}
// @Router /private/log [get]
func (h *PrivateHandler) GetAllLogs(c *fiber.Ctx) error {
	userID := c.Query("userID")
	title := c.Query("title")
	content := c.Query("content")
	logType := c.Query("type")

	logs, err := h.services.LogService.GetAllLogs(c.Context(), userID, title, content, logType)
	if err != nil {
		return err
	}

	// Converts to logDto for json tags
	var logDTOs []LogDTO
	for _, log := range logs {
		logDTO := LogDTO{
			ID:      log.ID(),
			UserID:  log.UserID(),
			Title:   log.Title(),
			Content: log.Content(),
		}
		logDTOs = append(logDTOs, logDTO)
	}

	return response.Response(200, "Status OK", logDTOs)
}

// @Tags Log
// @Summary Get log by ID
// @Description Retrieves a single log by its ID.
// @Accept json
// @Produce json
// @Param id path string true "Log ID"
// @Success 200 {object} response.BaseResponse{data=[]LogDTO}
// @Router /private/log/{id} [get]
func (h *PrivateHandler) GetLogByID(c *fiber.Ctx) error {
	id := c.Params("id")

	isExists, err := h.services.LogService.IsExists(c.Context(), id)
	if err != nil {
		return response.Response(500, "error while checking log", err)
	}
	if !isExists {
		return response.Response(400, "log not found", nil)
	}

	log, err := h.services.LogService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	logDTO := LogDTO{
		ID:      log.ID(),
		UserID:  log.UserID(),
		Title:   log.Title(),
		Content: log.Content(),
	}

	return response.Response(200, "Status OK", logDTO)
}
