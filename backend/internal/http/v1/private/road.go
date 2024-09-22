package private

import (
	"fmt"
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
	roadRoute.Get("/reset/:programmingID/:pathID", h.ResetPathHistory)
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
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, "", nil, nil)
	if err != nil {
		return err
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
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	roadData, err := h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, pathID, nil, nil)
	if err != nil {
		return err
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, pathID, domains.TypePath, inventoryInformation.GetFileExtension())
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

	if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, pathID, domains.TypePath, domains.ContentStarted); err != nil {
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
	roadStatsDTO := h.dtoManager.RoadDTOManager.ToRoadStatsDTO(roadStats)

	return response.Response(200, "Get User Language Road Stats", roadStatsDTO)
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
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	var answerRoadDTO dto.AnswerRoadDTO
	if err := c.BodyParser(&answerRoadDTO); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(answerRoadDTO); err != nil {
		return err
	}

	programmingInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	tmpPath, err := h.services.CodeService.UploadUserCode(c.Context(), userSession.UserID, programmingID, pathID, domains.TypePath, programmingInformation.GetFileExtension(), answerRoadDTO.UserCode)
	if err != nil {
		return err
	}
	road, err := h.services.RoadService.GetRoadByID(userSession.UserID, programmingID, pathID)
	if err != nil {
		return err
	}
	codeTmp := road.GetQuest().GetCodeTemplates()[0] // Çünkü road hangi dil için ise onun template'i kullanılıcak başka gerek yok.

	tmpContent, err := h.services.CodeService.CodeDockerTemplateGenerator(codeTmp.GetTemplatePath(), road.GetQuest().GetFuncName(), answerRoadDTO.UserCode, road.GetQuest().GetTests())
	if err != nil {
		return err
	}
	if err := h.services.CodeService.CreateFileAndWrite(tmpPath, tmpContent); err != nil {
		return err
	}

	var cmd []string
	var logs string
	if strings.EqualFold(road.GetQuest().GetFuncName(), "main") {
		cmd, err = h.services.CodeService.ChangeCMD(programmingInformation.GetCmd(), road.GetQuest().GetTests(), userSession.UserID)
		if err != nil {
			return err
		}
		logs, err = h.services.CodeService.RunContainerWithTar(c.Context(), programmingInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", programmingInformation.GetFileExtension()), cmd)
		if err != nil {
			return err
		}
	} else {
		logs, err = h.services.CodeService.RunContainerWithTar(c.Context(), programmingInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", programmingInformation.GetFileExtension()), programmingInformation.GetCmd())
		if err != nil {
			return err
		}
	}

	if strings.Contains(logs, "Test Passed") {
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, pathID, domains.TypePath, domains.ContentCompleted); err != nil {
			return err
		}
	}

	return response.Response(200, logs, nil)
}

// @Tags Road
// @Summary ResetPathHistory
// @Description Reset Path By Programming Language ID & Path ID
// @Accept json
// @Produce json
// @Param pathID path string true "Path ID"
// @Param programmingID path string true "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/reset/{programmingID}/{pathID} [get]
func (h *PrivateHandler) ResetPathHistory(c *fiber.Ctx) error {
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	_, err = h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, pathID, nil, nil)
	if err != nil {
		return err
	}

	err = h.services.CodeService.DeleteFrontendTemplateHistory(userSession.UserID, programmingID, pathID, domains.TypePath, inventoryInformation.GetFileExtension())
	if err != nil {
		return err
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, pathID, domains.TypePath, inventoryInformation.GetFileExtension())
	if err != nil {
		return err
	}
	frontendTemplateDto := h.dtoManager.RoadDTOManager.ToFrontendTemplateDto(frontendTemplate)

	return response.Response(200, "ResetPathHistory successful", frontendTemplateDto)
}
