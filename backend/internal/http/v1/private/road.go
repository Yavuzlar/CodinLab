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
	roadRoute.Get("/:programmingID", h.GetRoad)
	root.Get("/path/:programmingID/:pathID", h.GetPath)
	roadRoute.Get("/general/stats", h.GetUserLanguageRoadStats)
	roadRoute.Get("/path/data", h.AddDummyRoadData)
	roadRoute.Get("/progress/stats", h.GetUserRoadProgressStats)
}

// @Tags Road
// @Summary GetRoads
// @Description Get Road with Paths
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{data=dto.GetRoadDTO}
// @Router /private/road/{programmingID} [get]
func (h *PrivateHandler) GetRoad(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Declare and assign roadID variable
	programmingID := c.Params("programmingID")
	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	isExist, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}
	if !isExist {
		return response.Response(500, "Programming Language could not started", nil)
	}

	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, num, 0, nil, nil)
	if err != nil {
		return err
	}
	if len(roads) == 0 {
		return response.Response(404, "Road not found", nil)
	}

	var pathDTOs []dto.GetRoadPathDTO
	var roadDTO dto.GetRoadDTO
	for _, road := range roads {
		for _, path := range road.GetPaths() {
			languageDTOs := h.dtoManager.RoadDTOManager.ToLanguageRoadDTOs(path.GetLanguages())
			pathDTOs = append(pathDTOs, h.dtoManager.RoadDTOManager.ToRoadPathDTO(path, languageDTOs))
		}
		roadDTO = h.dtoManager.RoadDTOManager.ToGetRoadDTO(road, pathDTOs)
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
// @Router /private/road/path/{programmingID}/{pathID} [get]
func (h *PrivateHandler) GetPath(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	pathID := c.Params("pathID")

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	intPathID, err := strconv.Atoi(pathID)
	if err != nil {
		return response.Response(400, "Invalid Path ID", nil)
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

// @Tags Road
// @Summary GetUserLanguageRoadStats
// @Description Gets users language road stats
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/general/stats [get]
func (h *PrivateHandler) GetUserLanguageRoadStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	roadStats, err := h.services.RoadService.GetUserLanguageRoadStats(userSession.UserID)
	if err != nil {
		return err
	}
	roadDTO := h.dtoManager.RoadDTOManager.ToRoadStatsDTO(roadStats)

	return response.Response(200, "Get User Language Road Stats", roadDTO)
}

// @Tags Road
// @Summary DummyLogData
// @Description Add dummy data for testing
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/path/data [get]
func (h *PrivateHandler) AddDummyRoadData(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Dummy Data for testing
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentStarted, 1, 2)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentStarted, 2, 1)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentStarted, 2, 2)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentStarted, 1, 1)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentCompleted, 1, 1)

	return response.Response(200, "Dummy Data Added", nil)
}

// @Tags Road
// @Summary GetUserRoadProgressStats
// @Description Gets users road progress stats
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/progress/stats [get]
func (h *PrivateHandler) GetUserRoadProgressStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	roadStats, err := h.services.RoadService.GetUserRoadProgressStats(userSession.UserID)
	if err != nil {
		return err
	}
	roadDTO := h.dtoManager.RoadDTOManager.ToRoadProgressDTO(*roadStats)

	return response.Response(200, "Get User Road Progress Stats", roadDTO)
}
