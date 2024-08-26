package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	roadRoute := root.Group("/road")
	roadRoute.Post("/start", h.Start)
	roadRoute.Get("/:programmingID", h.GetRoad)
	roadRoute.Get("/:programmingID/:pathID", h.GetPath)
}

// @Tags Road
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Param start body dto.StartDTO true "Start"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/start [post]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	var start dto.StartDTO
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
	ok, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeRoad, domains.ContentStarted, start.ProgrammingID, 0)
	if err != nil {
		return err
	}

	if !ok {
		if start.ProgrammingID == 0 {
			return response.Response(200, "Invalid Programming ID", nil)
		}
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeRoad, domains.ContentStarted, start.ProgrammingID, 0); err != nil {
			return err
		}
	}

	// if the "Road Started Successfully" message recived. The frontend will redirect the user to the spesific road.
	return response.Response(200, "Road Started Successfully", nil)
}

// @Tags Road
// @Summary GetRoads
// @Description Get Road with Paths
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{data=dto.RoadDTO}
// @Router /private/road/{programmingID} [get]
func (h *PrivateHandler) GetRoad(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Declare and assign roadID variable
	programmingID := c.Params("programmingID")
	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, num, 0, nil, nil)
	if err != nil {
		return err
	}
	if len(roads) == 0 {
		return response.Response(404, "Road not found", nil)
	}
	var pathDTOs []dto.PathDTO
	var roadDTO dto.RoadDTO
	for _, road := range roads {
		for _, path := range road.GetPaths() {
			languageDTOs := h.dtoManager.RoadDTOManager.ToLanguageDTOs(path.GetLanguages())
			pathDTOs = append(pathDTOs, h.dtoManager.RoadDTOManager.ToPathDTO(path, languageDTOs))
		}
		roadDTO = h.dtoManager.RoadDTOManager.ToRoadDTO(road, pathDTOs)
	}
	return response.Response(200, "GetRoads successful", roadDTO)
}

// @Tags Road
// @Summary GetPathByID
// @Description Get Path By ID
// @Accept json
// @Produce json
// @Param programmingID path string true "Programming ID"
// @Param pathID path string true "Path ID"
// @Success 200 {object} response.BaseResponse{data=dto.PathDTO}
// @Router /private/road/{programmingID}/{pathID} [get]
func (h *PrivateHandler) GetPath(c *fiber.Ctx) error {

	programmingID := c.Params("programmingID")
	pathID := c.Params("pathID")

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	intPathID, err := strconv.Atoi(pathID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	session := session_store.GetSessionData(c)
	userID := session.UserID

	roadData, err := h.services.RoadService.GetRoadFilter(userID, intProgrammingID, intPathID, nil, nil)
	if err != nil {
		return err
	}
	if len(roadData) == 0 {
		return response.Response(404, "Path not found", nil)
	}

	var roadDTO []dto.RoadDTO

	for _, road := range roadData {
		var pathsDTO []dto.PathDTO
		for _, path := range road.GetPaths() {
			languageDTOs := h.dtoManager.RoadDTOManager.ToLanguageDTOs(path.GetLanguages())
			pathsDTO = append(pathsDTO, h.dtoManager.RoadDTOManager.ToPathDTO(path, languageDTOs))
		}
		roadDTO = append(roadDTO, h.dtoManager.RoadDTOManager.ToRoadDTO(road, pathsDTO))
	}

	return response.Response(200, "Path Retrieved Successfully", roadDTO)
}
