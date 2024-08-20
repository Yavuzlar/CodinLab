package private

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	root.Get("/log", h.GetAllLogs)
	root.Get("/log/solution/byday", h.GetSolutionsByDay)
	root.Get("/log/solution/hours", h.GetSolutionsHoursByProgramming)
}

// @Tags Log
// @Summary Get all logs
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param userID query string false "User ID"
// @Param programmingID query int32 false "Programming ID"
// @Param labRoadID query int32 false "Log Lab or Path ID"
// @Param content query string false "Log Content"
// @Param type query string false "Log Type"
// @Success 200 {object} response.BaseResponse{data=[]dto.LogDTO}
// @Router /private/log [get]
func (h *PrivateHandler) GetAllLogs(c *fiber.Ctx) error {
	userID := c.Query("userID")
	programmingID := c.Query("programmingID")
	labRoadID := c.Query("labRoadID")
	content := c.Query("content")
	logType := c.Query("type")

	logs, err := h.services.LogService.GetAllLogs(c.Context(), userID, programmingID, labRoadID, logType, content)
	if err != nil {
		return err
	}

	// Converts to logDto for json tags
	logDTOs := h.dtoManager.LogDTOManager.ToLogDTOs(logs)

	return response.Response(200, "Status OK", logDTOs)
}

// @Tags Log
// @Summary GetSolutionsByDay
// @Description Retrieves the number of lab and road solutions solved day by day.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]dto.SolutionsByDayDTO}
// @Router /private/log/solution/byday [get]
func (h *PrivateHandler) GetSolutionsByDay(c *fiber.Ctx) error {
	solutionsByDay, err := h.services.LogService.CountSolutionsByDay(c.Context())
	if err != nil {
		return err
	}
	solutionsByDayDTOs := h.dtoManager.LogDTOManager.ToSolutionsByDayDTOs(solutionsByDay)

	return response.Response(200, "Status OK", solutionsByDayDTOs)
}

// @Tags Log
// @Summary GetSolutionsHoursByProgramming
// @Description Retrieves the total hours spent on lab and road solutions for each programming language in the last week.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]dto.SolutionsHoursByProgrammingDTO}
// @Router /private/log/solution/hours [get]
func (h *PrivateHandler) GetSolutionsHoursByProgramming(c *fiber.Ctx) error {
	solutionsHours, err := h.services.LogService.CountSolutionsHoursByProgrammingLast7Days(c.Context())
	if err != nil {
		return err
	}
	solutionsHoursDTOs := h.dtoManager.LogDTOManager.ToSolutionsHoursByProgrammingDTOs(solutionsHours)

	return response.Response(200, "Status OK", solutionsHoursDTOs)
}
