package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:id", h.GetLabsByID)
	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
}

// @Tags Lab
// @Summary GetLab
// @Description Get Lab By ID
// @Accept json
// @Produce json
// @Param id path string true "Lab ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{id} [get]
func (h *PrivateHandler) GetLabsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, num, 0, "", "")
	if err != nil {
		return err
	}

	if len(labData) == 0 {
		return response.Response(404, "Labs not found", labData)
	}

	return response.Response(200, "GetLabs successful", labData)
}
