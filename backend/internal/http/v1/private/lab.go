package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:id", h.GetLabsByID)
	root.Get("/lab/:langId/:labId", h.GetLabByID)
	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
}

// @Tags Lab
// @Summary GetLabsById
// @Description Get Labs By ID
// @Accept json
// @Produce json
// @Param id path string true "Labs ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{id} [get]
func (h *PrivateHandler) GetLabsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intId, 0, "", "")
	if err != nil {
		return err
	}

	if len(labData) == 0 {
		return response.Response(404, "Labs not found", labData)
	}

	return response.Response(200, "GetLabs successful", labData)
}

// @Tags Lab
// @Summary GetLabById
// @Description Get Lab By ID
// @Accept json
// @Produce json
// @Param langId path string true "Lang ID"
// @Param labId path string true "Lab ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/{langId}/{labId} [get]
func (h *PrivateHandler) GetLabByID(c *fiber.Ctx) error {
	langId := c.Params("langId")
	labId := c.Params("labId")

	intLangId, err := strconv.Atoi(langId)
	if err != nil {
		return response.Response(400, "Invalid Lang ID", nil)
	}
	intLabId, err := strconv.Atoi(labId)
	if err != nil {
		return response.Response(400, "Invalid Labs ID", nil)
	}

	userSession := session_store.GetSessionData(c)

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intLangId, intLabId, "", "")
	if err != nil {
		return err
	}

	if len(labData) == 0 {
		return response.Response(404, "Lab not found", labData)
	}

	return response.Response(200, "GetLab successful", labData)
}
