package private

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	roadRoute := root.Group("/road")
	roadRoute.Get("/:programmingID", h.GetRoad)
	roadRoute.Get("/path/:programmingID/:pathID", h.GetPath)
	roadRoute.Get("/general/stats", h.GetUserLanguageRoadStats)
	roadRoute.Get("/path/data", h.AddDummyRoadData)
	roadRoute.Get("/progress/stats", h.GetUserRoadProgressStats)
	roadRoute.Post("/answer/:programmingID/:pathID", h.AnswerRoad)
}

// @Tags Road
// @Summary GetRoads
// @Description Get Road with Paths
// @Accept json
// @Produce json
// @Param Language header string false "Language"
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{data=dto.GetRoadDTO}
// @Router /private/road/{programmingID} [get]
func (h *PrivateHandler) GetRoad(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	language := c.Get("Language")
	if language == "" {
		language = "en"
	}
	programmingID := c.Params("programmingID")
	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
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
			languageDTO := h.dtoManager.RoadDTOManager.ToLanguageRoadDTO(path.GetLanguages(), language)
			pathDTOs = append(pathDTOs, h.dtoManager.RoadDTOManager.ToRoadPathDTO(path, languageDTO))
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
// @Param Language header string false "Language"
// @Param programmingID path string true "Programming ID"
// @Param pathID path string true "Path ID"
// @Success 200 {object} response.BaseResponse{data=dto.PathDTO}
// @Router /private/road/path/{programmingID}/{pathID} [get]
func (h *PrivateHandler) GetPath(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	programmingID := c.Params("programmingID")
	pathID := c.Params("pathID")
	language := c.Get("Language")
	if language == "" {
		language = "en"
	}

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

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(int32(intProgrammingID))
	if err != nil {
		return response.Response(500, "Programming Language Information Error", err)
	}
	if inventoryInformation == nil {
		return response.Response(404, "Programming Language Not Found", nil)
	}

	roadData, err := h.services.RoadService.GetRoadFilter(userID, intProgrammingID, intPathID, nil, nil)
	if err != nil {
		return err
	}
	if len(roadData) == 0 {
		return response.Response(404, "Path not found", nil)
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, domains.TypePath, intProgrammingID, intPathID, inventoryInformation.GetFileExtension())
	if err != nil {
		return err
	}

	var roadDTO []dto.RoadDTO
	for _, road := range roadData {
		var pathsDTO []dto.PathDTO
		for _, path := range road.GetPaths() {
			languageDTO := h.dtoManager.RoadDTOManager.ToLanguageRoadDTO(path.GetLanguages(), language)
			pathsDTO = append(pathsDTO, h.dtoManager.RoadDTOManager.ToPathDTO(path, languageDTO, frontendTemplate))
		}
		roadDTO = append(roadDTO, h.dtoManager.RoadDTOManager.ToRoadDTO(road, pathsDTO))
	}

	if err := h.services.LogService.Add(c.Context(), userID, domains.TypePath, domains.ContentStarted, int32(intProgrammingID), int32(intPathID)); err != nil {
		return err
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

// @Tags Road
// @Summary Answer
// @Description This is for answering quests.
// @Accept json
// @Produce json
// @Param programmingID path string true "Programming ID"
// @Param pathID path string true "Path ID"
// @Param answerRoadDTO body dto.AnswerRoadDTO true "Answer Road DTO"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/answer/{programmingID}/{pathID} [post]
func (h *PrivateHandler) AnswerRoad(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	var answerRoadDTO dto.AnswerRoadDTO
	if err := c.BodyParser(&answerRoadDTO); err != nil {
		return err
	}
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

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(int32(intProgrammingID))
	if err != nil {
		return response.Response(500, "Programming Language Information Error", nil)
	}
	if inventoryInformation == nil {
		return response.Response(404, "Programming Language Not Found", nil)
	}
	tmpPath, err := h.services.CodeService.UploadUserCode(c.Context(), userSession.UserID, intProgrammingID, intPathID, domains.TypePath, inventoryInformation.GetFileExtension(), answerRoadDTO.UserCode)
	if err != nil {
		return err
	}
	road, err := h.services.RoadService.GetRoadByID(userSession.UserID, intProgrammingID, intPathID)
	if err != nil {
		return response.Response(500, "Path Not Found", nil)
	}
	codeTmp := road.GetQuest().GetCodeTemplates()[0]

	tmpContent, err := h.services.CodeService.CodeDockerTemplateGenerator(codeTmp.GetTemplatePath(), road.GetQuest().GetFuncName(), answerRoadDTO.UserCode, road.GetQuest().GetTests())
	if err != nil {
		return err
	}

	if err := h.services.CodeService.CreateFileAndWrite(tmpPath, tmpContent); err != nil {
		return err
	}
	logs, err := h.services.CodeService.RunContainerWithTar(c.Context(), inventoryInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", inventoryInformation.GetFileExtension()), inventoryInformation.GetCmd())
	if err != nil {
		return err
	}
	if strings.Contains(logs, "Test Passed") {
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentCompleted, int32(intProgrammingID), int32(intPathID)); err != nil {
			return response.Response(500, "Docker Image Pull Error", nil)
		}
	}

	return response.Response(200, logs, nil)
}
