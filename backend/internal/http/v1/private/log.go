package private

import (
	"fmt"
	"time"

	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	root.Get("/log", h.GetAllLogs)
	root.Get("/log/solution/byday", h.GetSolutionsByDay)
	root.Get("/log/solution/hours", h.GetSolutionsHoursByLanguage)
}

type LogDTO struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"userId"`
	LanguageID int32     `json:"languageID"`
	LabRoadID  int32     `json:"labRoadID"`
	LType      string    `json:"type"`
	Content    string    `json:"content"`
}

// lab and road numbers solved day by day
// author: yasir
type SolutionsByDayDTO struct {
	Date      time.Time
	RoadCount int
	LabCount  int
}

// SolutionsHoursByLanguage represents the total hours spent on lab and road solutions for each language.
// author: yasir
type SolutionsHoursByLanguageDTO struct {
	LanguageID int32   `json:"language_id"`
	LabHours   float64 `json:"lab_hours"`
	RoadHours  float64 `json:"road_hours"`
}

// @Tags Log
// @Summary Get all logs
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param userID query string false "User ID"
// @Param languageID query int32 false "Language ID"
// @Param labRoadID query int32 false "Log Lab or Path ID"
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

	logs, err := h.services.LogService.GetAllLogs(c.Context(), userID, languageID, labRoadID, logType, content)
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
			LabRoadID:  log.LabPathID(),
			LType:      log.Type(),
			Content:    log.Content(),
		}
		logDTOs = append(logDTOs, logDTO)
	}

	return response.Response(200, "Status OK", logDTOs)
}

// @Tags Log
// @Summary GetSolutionsByDay
// @Description Retrieves the number of lab and road solutions solved day by day.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]SolutionsByDayDTO}
// @Router /private/log/solution/byday [get]
func (h *PrivateHandler) GetSolutionsByDay(c *fiber.Ctx) error {
	solutionsByDay, err := h.services.LogService.CountSolutionsByDay(c.Context())
	if err != nil {
		return err
	}

	var solutionsByDayDTOs []SolutionsByDayDTO
	for _, solution := range solutionsByDay {
		solutionsByDayDTOs = append(solutionsByDayDTOs, SolutionsByDayDTO{
			Date:      solution.Date,
			RoadCount: solution.RoadCount,
			LabCount:  solution.LabCount,
		})
	}

	return response.Response(200, "Status OK", solutionsByDayDTOs)
}

// @Tags Log
// @Summary GetSolutionsHoursByLanguage
// @Description Retrieves the total hours spent on lab and road solutions for each language in the last week.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]SolutionsHoursByLanguageDTO}
// @Router /private/log/solution/hours [get]
func (h *PrivateHandler) GetSolutionsHoursByLanguage(c *fiber.Ctx) error {
	solutionsHours, err := h.services.LogService.CountSolutionsHoursByLanguageLast7Days(c.Context())
	if err != nil {
		return err
	}

	var solutionsHoursDTOs []SolutionsHoursByLanguageDTO
	for _, solution := range solutionsHours {
		solutionsHoursDTOs = append(solutionsHoursDTOs, SolutionsHoursByLanguageDTO{
			LanguageID: solution.LanguageID,
			LabHours:   solution.LabHours,
			RoadHours:  solution.RoadHours,
		})
	}

	return response.Response(200, "Status OK", solutionsHoursDTOs)
}
