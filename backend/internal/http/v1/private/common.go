package private

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initCommonRoutes(root fiber.Router) {
	commonRoutes := root.Group("/common")
	commonRoutes.Get("/stop/:containerID", h.StopAPI)
}

// @Tags Common
// @Summary Stops Quest Tests
// @Description Stops Quest Tests
// @Accept json
// @Produce json
// @Param containerID path string true "Container ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/common/stop/{containerID} [get]
func (h *PrivateHandler) StopAPI(c *fiber.Ctx) error {
	containerID := c.Params("containerID")
	if err := h.services.CodeService.StopContainer(c.Context(), containerID); err != nil {
		return err
	}

	return response.Response(200, "Test Stopped successfully", nil)
}
