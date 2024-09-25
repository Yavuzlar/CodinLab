package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	root.Get("/log", h.GetAllLogs)
	root.Get("/log/solution/byday", h.GetSolutionsByDay)
	root.Get("/log/solution/hours", h.GetSolutionsHoursByProgramming)
	root.Get("/log/lab", h.AddDummyLabData)
	root.Get("/log/road", h.AddDummyRoadData)
	root.Get("/log/rates", h.LanguageUsageRates)
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
	logType := c.Query("type")
	userID := c.Query("userID")
	content := c.Query("content")
	labRoadID := c.Query("labRoadID")
	programmingID := c.Query("programmingID")

	logs, err := h.services.LogService.GetAllLogs(c.Context(), userID, programmingID, labRoadID, logType, content)
	if err != nil {
		return err
	}
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

// @Tags Log
// @Summary DummyLogData
// @Description Add dummy data for testing
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/log/lab [get]
func (h *PrivateHandler) AddDummyLabData(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Dummy Data for testing
	h.services.LogService.Add(c.Context(), userSession.UserID, "1", "1", domains.TypeLab, domains.ContentStarted)
	h.services.LogService.Add(c.Context(), userSession.UserID, "1", "1", domains.TypeLab, domains.ContentCompleted)

	return response.Response(200, "Dummy Data Added", nil)
}

// @Tags Log
// @Summary DummyLogData
// @Description Add dummy data for testing
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/log/road [get]
func (h *PrivateHandler) AddDummyRoadData(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Dummy Data for testing
	h.services.LogService.Add(c.Context(), userSession.UserID, "1", "1", domains.TypePath, domains.ContentStarted)
	h.services.LogService.Add(c.Context(), userSession.UserID, "2", "1", domains.TypePath, domains.ContentStarted)
	h.services.LogService.Add(c.Context(), userSession.UserID, "2", "1", domains.TypePath, domains.ContentCompleted)

	return response.Response(200, "Dummy Data Added", nil)
}

// @Tags Log
// @Summary Get Language Usage Rates
// @Description Retrieves language usage rates
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Failure 400 {object} response.BaseResponse
// @Router /private/log/rates [get]
func (h *PrivateHandler) LanguageUsageRates(c *fiber.Ctx) error {
	rateLogs, err := h.services.LogService.LanguageUsageRates(c.Context())
	if err != nil {
		return err
	}
	rateLogsDTO := h.dtoManager.LogDTOManager.ToLanguageUsageRatesDTOs(rateLogs)

	return response.Response(200, "STATUS OK", rateLogsDTO)
}
