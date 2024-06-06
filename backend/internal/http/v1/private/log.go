package private

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	root.Get("/log", h.GetAllLogs)
}

type LogDTO struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"userId"`
	LanguageID int32     `json:"languageID"`
	LabRoadID  int32     `json:"labRoadID"`
	LType      string    `json:"type"`
	Content    string    `json:"content"`
}

// @Tags Log
// @Summary Get all logs
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param userID query string false "User ID"
// @Param languageID query int32 false "Language ID"
// @Param labRoadID query int32 false "Log Lab or Road ID"
// @Param content query string false "Log Content"
// @Param type query string false "Log Type"
// @Success 200 {object} response.BaseResponse{data=[]LogDTO}
// @Router /private/log [get]
func (h *PrivateHandler) GetAllLogs(c *fiber.Ctx) error {
	userID := c.Query("userID")
	languageID := c.Query("languageID")
	labRoadID := c.Query("labRoadID")
	content := c.Query("content")
	logType := c.Query("type")

	logs, err := h.services.LogService.GetAllLogs(c.Context(), userID, languageID, labRoadID, content, logType)
	if err != nil {
		return err
	}

	// Converts to logDto for json tags
	var logDTOs []LogDTO
	for _, log := range logs {
		logDTO := LogDTO{
			ID:         log.ID(),
			UserID:     log.UserID(),
			LanguageID: log.LanguageID(),
			LabRoadID:  log.LabRoadID(),
			Content:    log.Content(),
		}
		logDTOs = append(logDTOs, logDTO)
	}

	return response.Response(200, "Status OK", logDTOs)
}
