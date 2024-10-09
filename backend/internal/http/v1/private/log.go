package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLogRoutes(root fiber.Router) {
	logRoutes := root.Group("/log")
	logRoutes.Get("/", h.GetAllLogs)
	logRoutes.Get("/solution/byday/:year", h.GetSolutionsByDay)
	logRoutes.Get("/solution/week", h.GetSolutionsByProgramming)
	logRoutes.Get("/lab", h.AddDummyLabData)
	logRoutes.Get("/road", h.AddDummyRoadData)
	logRoutes.Get("/rates", h.LanguageUsageRates)
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
// @Param year path string true "Desired Year"
// @Success 200 {object} response.BaseResponse{data=[]dto.SolutionsByDayDTO}
// @Router /private/log/solution/byday/{year} [get]
func (h *PrivateHandler) GetSolutionsByDay(c *fiber.Ctx) error {
	year := c.Params("year")
	solutionsByDay, err := h.services.LogService.CountSolutionsByDay(c.Context(), year)
	if err != nil {
		return err
	}
	solutionsByDayDTOs := h.dtoManager.LogDTOManager.ToSolutionsByDayDTOs(solutionsByDay)

	return response.Response(200, "Status OK", solutionsByDayDTOs)
}

// @Tags Log
// @Summary GetSolutionsByProgramming
// @Description Retrieves the total counts for lab and road solutions for each programming language in the last week.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{data=[]dto.SolutionsByProgrammingDTO}
// @Router /private/log/solution/week [get]
func (h *PrivateHandler) GetSolutionsByProgramming(c *fiber.Ctx) error {
	solutions, err := h.services.LogService.CountSolutionsByProgrammingLast7Days(c.Context())
	if err != nil {
		return err
	}
	solutionsDTOs := h.dtoManager.LogDTOManager.ToSolutionsByProgrammingDTOs(solutions)

	return response.Response(200, "Status OK", solutionsDTOs)
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
