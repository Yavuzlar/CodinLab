package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLabRoadRoutes(root fiber.Router) {
	root.Get("/start/:programmingID", h.Start)
}

// @Tags LabRoadCommon
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/start/{programmingID} [get]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	plInformation, err := h.services.LabRoadService.GetInventoryInformation(int32(num))
	if err != nil {
		return response.Response(500, "Get Programming Language error", nil)
	}
	if plInformation == nil {
		return response.Response(404, "Programming Language Not Found", nil)
	}

	// Recive user session from session_store
	userSession := session_store.GetSessionData(c)

	isExsits, err := h.services.CodeService.IsImageExists(c.Context(), plInformation.GetDockerImage())
	if err != nil {
		return response.Response(500, "Docker Image Check Error", nil)
	}

	if !isExsits {
		if err := h.services.CodeService.Pull(c.Context(), plInformation.GetDockerImage()); err != nil {
			return response.Response(500, "Docker Image Pull Error", nil)
		}
	}

	// if the road has started. Log will not be created
	// Log a road start event for the user
	ok, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}

	if !ok {
		if num == 0 {
			return response.Response(200, "Invalid Programming ID", nil)
		}
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0); err != nil {
			return response.Response(500, "Error adding log", nil)
		}
	}

	isExist, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}
	if !isExist {
		return response.Response(500, "Programming Language could not started", nil)
	}

	return response.Response(200, "Progamming Language Started Successfully", nil)
}
