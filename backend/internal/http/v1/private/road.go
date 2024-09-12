package private

import (
	"fmt"
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
	roadRoute.Get("/path/:programmingID/:pathID", h.GetPath)
	roadRoute.Get("/general/stats", h.GetUserLanguageRoadStats)
	roadRoute.Get("/path/data", h.AddDummyRoadData)
	roadRoute.Get("/progress/stats", h.GetUserRoadProgressStats)
	roadRoute.Post("/answer/:programmingID/:pathID", h.AnswerRoad)
	roadRoute.Get("/template/:pathID/:programmingID", h.GetRoadTemplate)
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

	tmpContent, err := h.services.CodeService.CodeDockerTemplateGenerator(codeTmp.GetTemplate(), codeTmp.GetCheck(), answerRoadDTO.UserCode, road.GetQuest().GetFuncName(), road.GetQuest().GetTests(), road.GetQuest().GetReturns())
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

	if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypePath, domains.ContentStarted, int32(intProgrammingID), int32(intPathID)); err != nil {
		return response.Response(500, "Docker Image Pull Error", nil)
	}

	return response.Response(200, logs, nil)
}

// @Tags Road
// @Summary Get Road Template
// @Description Get Road Template
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Param pathID path string false "pathID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/template/{pathID}/{programmingID} [get]
func (h *PrivateHandler) GetRoadTemplate(c *fiber.Ctx) error {
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	pathIDInt, err := strconv.Atoi(pathID)
	if err != nil {
		return response.Response(400, "Invalid Lab ID", nil)
	}

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(int32(num))
	if err != nil {
		return response.Response(500, "Programming Language Information Error", err)
	}
	if inventoryInformation == nil {
		return response.Response(404, "Programming Language Not Found", nil)
	}

	path, err := h.services.RoadService.GetRoadByID(userSession.UserID, num, pathIDInt)
	if err != nil {
		return response.Response(404, "Error While Getting Path", err)
	}
	if path == nil {
		return response.Response(404, "Path Not Found", nil)
	}

	codeTmp := path.GetQuest().GetCodeTemplates()[0]

	frontendContent := h.services.CodeService.CodeFrontendTemplateGenerator(inventoryInformation.GetName(), path.GetQuest().GetFuncName(), codeTmp.GetFrontend(), path.GetQuest().GetParams(), path.GetQuest().GetReturns(), path.GetQuest().GetQuestImports())

	return response.Response(200, "Template Successfully Sent", frontendContent)
}
