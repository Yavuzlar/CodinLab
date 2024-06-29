package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

type StartDTO struct {
	LanguageID int32 `json:"languageID" validate:"required"`
}

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	roadRoute := root.Group("/road")

	roadRoute.Post("/start", h.Start)
	roadRoute.Get("/:roadId", h.GetAllRoads)
}

// @Tags Road
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Param start body StartDTO true "Start"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/start [post]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	var start StartDTO
	if err := c.BodyParser(&start); err != nil {
		return err
	}

	// NEED ROAD SERVICE FOR BOTTOM
	// We have to get the road that will start according to the id
	// In that road struct we will change isStarted to true.
	// And we need road's docker image for DocerService

	// Need Road Service For -> Road title & Docker Image For Log
	// With road service we will get road by road id and recive docker image.

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
	if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeRoad, domains.ContentStarted, start.LanguageID, 0); err != nil {
		return err
	}

	// if the "Road Started Successfully" message recived. The frontend will redirect the user to the spesific road.
	return response.Response(200, "Road Started Successfully", nil)
}

// @Tags Road
// @Summary GetRoads
// @Description Get All Roads
// @Accept json
// @Produce json
// @Param roadId path string true "roadId"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/{roadId} [get]
func (h *PrivateHandler) GetAllRoads(c *fiber.Ctx) error {
	// NEED ROAD SERVICE FOR BOTTOM
	// We have to get all roads and send them to the frontend.

	// Recive user session from session_store

	userSession := session_store.GetSessionData(c)

	// Declare and assign roadId variable
	roadId := c.Params("roadId")
	num, err := strconv.Atoi(roadId)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	// Get all roads
	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, num, 0, false, false)
	if err != nil {
		return err
	}

	return response.Response(200, "GetRoads successful", roads)
}

