package private

import (
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
	userSession := session_store.GetSessionData(c)
	programmingID := c.Params("programmingID")

	programmingInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	isExsits, err := h.services.CodeService.IsImageExists(c.Context(), programmingInformation.GetDockerImage())
	if err != nil {
		return err
	}

	if !isExsits {
		if err := h.services.CodeService.Pull(c.Context(), programmingInformation.GetDockerImage()); err != nil {
			return err
		}
	}

	ok, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, programmingID, "", domains.TypeProgrammingLanguage, domains.ContentStarted)
	if err != nil {
		return err
	}

	if !ok {
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, "", domains.TypeProgrammingLanguage, domains.ContentStarted); err != nil {
			return err
		}
	}

	isExist, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, programmingID, "", domains.TypeProgrammingLanguage, domains.ContentStarted)
	if err != nil {
		return err
	}
	if !isExist {
		return response.Response(500, "Programming Language could not started", nil)
	}

	return response.Response(200, "Progamming Language Started Successfully", nil)
}
