package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	root.Post("/road/start", h.Start)
}

// @Tags Road
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/start [post]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	// Need Road Service For -> Road title & Docker Image For Log

	// Recive user session from session_store
	userSession := session_store.GetSessionData(c)

	// Need Spesific imageReference for this. If it's wrong it wont work
	isExsits, err := h.services.DockerService.IsImageExists(c.Context(), "golang:latest")
	if err != nil {
		return err
	}

	if !isExsits {
		// Downloads Spesific Image. This golang fetched from road.json
		if err := h.services.DockerService.Pull(c.Context(), "golang:latest"); err != nil {
			return err
		}

	}

	// if the road has started. Log will not be created
	// Log a road start event for the user
	if err := h.services.LogService.Add(c.Context(), userSession.UserID, "", domains.TypeRoad, domains.ContentStarted); err != nil {
		return err
	}

	// if the "Road Started Successfully" message recived. The frontend will redirect the user to the spesific road.
	return response.Response(200, "Road Started Successfully", nil)
}
