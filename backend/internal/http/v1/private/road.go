package private

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	roadRoute := root.Group("/road")
	roadRoute.Post("/start", h.StartRoad)
	roadRoute.Get("/:programmingID", h.GetRoad)
	roadRoute.Get("/path/:programmingID/:pathID", h.GetPath)
	roadRoute.Get("/reset/:programmingID/:pathID", h.ResetPathHistory)
	roadRoute.Get("/general/stats", h.GetUserLanguageRoadStats)
	roadRoute.Get("/path/data", h.AddDummyRoadData)
	roadRoute.Get("/progress/stats", h.GetUserRoadProgressStats)
	roadRoute.Post("/answer/:programmingID/:pathID", h.AnswerRoad)
}

// @Tags Road
// @Summary StartRoad
// @Description Start Road
// @Accept json
// @Produce json
// @Param start body dto.StartDTO true "Start"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/start [post]
func (h *PrivateHandler) StartRoad(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	var startDTO dto.StartDTO

	if err := c.BodyParser(&startDTO); err != nil {
		return err
	}

	programmingID := strconv.Itoa(int(startDTO.ProgrammingID))

	isStarted := false
	_, err := h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, "", &isStarted, nil)
	if err != nil {
		return err
	}
	isExists, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, programmingID, "", domains.TypeRoad, domains.ContentStarted)
	if err != nil {
		return err
	}
	if isExists {
		return response.Response(409, "Road was started already", nil)
	}
	if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, "", domains.TypeRoad, domains.ContentStarted); err != nil {
		return err
	}
	return response.Response(200, "Road Started successfully", nil)
}

// @Tags Road
// @Summary GetRoads
// @Description Get Road with Paths
// @Accept json
// @Produce json
// @Param Language header string false "Language"
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{data=dto.RoadDTO}
// @Router /private/road/{programmingID} [get]
func (h *PrivateHandler) GetRoad(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	inventoryInformation, err := h.services.CommonService.GetInventoryInformation(programmingID, language)
	if err != nil {
		return err
	}

	isExists, err := h.services.CodeService.IsImageExists(c.Context(), inventoryInformation.GetDockerImage())
	if err != nil {
		return err
	}
	var conn *websocket.Conn
	for c, ok := range h.clients {
		if c.GetUserID().String() == userSession.UserID && ok {
			conn = c.GetConnection()
			break
		}
	}

	if !isExists {
		go func() {
			if err := h.services.CodeService.Pull(c.Context(), inventoryInformation.GetDockerImage(), inventoryInformation.GetName(), conn); err != nil {
				fmt.Printf("Error pulling Docker image: %v\n", err)
			}
		}()
	}

	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, "", nil, nil)
	if err != nil {
		return err
	}

	var pathDTOs []dto.PathDTO
	var roadDTO dto.RoadDTO
	for _, road := range roads {
		for _, path := range road.GetPaths() {
			languageDTO := h.dtoManager.RoadDTOManager.ToLanguageRoadDTO(path.GetLanguages(), language)
			pathDTOs = append(pathDTOs, h.dtoManager.RoadDTOManager.ToPathDTO(path, languageDTO, ""))
		}
		roadDTO = h.dtoManager.RoadDTOManager.ToRoadDTO(road, pathDTOs, isExists, *inventoryInformation.GetLanguage(), inventoryInformation.GetFileExtension())
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

	inventoryInformation, err := h.services.CommonService.GetInventoryInformation(programmingID, language)
	if err != nil {
		return err
	}

	isExists, err := h.services.CodeService.IsImageExists(c.Context(), inventoryInformation.GetDockerImage())
	if err != nil {
		return err
	}
	if !isExists {
		return response.Response(400, fmt.Sprintf("%s image does not exist. Please visit the homepage to download it.", inventoryInformation.GetDockerImage()), nil)
	}

	roadData, err := h.services.RoadService.GetRoadFilter(userSession.UserID, programmingID, pathID, nil, nil)
	if err != nil {
		return err
	}
	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, pathID, domains.TypePath, inventoryInformation.GetFileExtension(), true)
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
		roadDTO = append(roadDTO, h.dtoManager.RoadDTOManager.ToRoadDTO(road, pathsDTO, isExists, *inventoryInformation.GetLanguage(), inventoryInformation.GetFileExtension()))
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
// @Param Language header string false "Language"
// @Param pathID path string true "Path ID"
// @Param answerRoadDTO body dto.AnswerRoadDTO true "Answer Road DTO"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/answer/{programmingID}/{pathID} [post]
func (h *PrivateHandler) AnswerRoad(c *fiber.Ctx) error {
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	var answerRoadDTO dto.AnswerRoadDTO
	if err := c.BodyParser(&answerRoadDTO); err != nil {
		return err
	}
	if err := h.services.UtilService.Validator().ValidateStruct(answerRoadDTO); err != nil {
		return err
	}

	programmingInformation, err := h.services.CommonService.GetInventoryInformation(programmingID, language)
	if err != nil {
		return err
	}

	road, err := h.services.RoadService.GetPathByID(userSession.UserID, programmingID, pathID)
	if err != nil {
		return err
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, pathID, domains.TypePath, programmingInformation.GetFileExtension(), false)
	if err != nil {
		return err
	}

	if frontendTemplate == answerRoadDTO.UserCode {
		return service_errors.NewServiceErrorWithMessageAndError(400, "SAME_CODE_ERROR", err)
	}

	tmpPath, err := h.services.CodeService.UploadUserCode(userSession.UserID, programmingID, pathID, domains.TypePath, programmingInformation.GetFileExtension(), answerRoadDTO.UserCode)
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

	var conn *websocket.Conn
	for c, ok := range h.clients {
		if c.GetUserID().String() == userSession.UserID && ok {
			conn = c.GetConnection()
			break
		}
	}
	// TODO: Belki Getirebilirsin
	// if conn == nil {
	// 	return response.Response(500, "This user was not found in socket.", nil)
	// }

	var logs string
	if strings.EqualFold(road.GetQuest().GetFuncName(), "main") {
		err = h.services.CodeService.CreateBashFile(programmingInformation.GetCmd(), road.GetQuest().GetTests(), userSession.UserID, programmingInformation.GetPathDir())
		if err != nil {
			return err
		}
		logs, err = h.services.CodeService.RunContainerWithTar(c.Context(), programmingInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", programmingInformation.GetFileExtension()), programmingInformation.GetShCmd(), conn)
		if err != nil {
			return err
		}
	} else {
		logs, err = h.services.CodeService.RunContainerWithTar(c.Context(), programmingInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", programmingInformation.GetFileExtension()), programmingInformation.GetCmd(), conn)
		if err != nil {
			return err
		}
	}

	if strings.Contains(logs, "Test Passed") {
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, pathID, domains.TypePath, domains.ContentCompleted); err != nil {
			return err
		}
	}

	parsedLog, err := h.services.CodeService.ParseCodeLog(logs)
	if err != nil {
		return err
	}

	//if there is no error and the function does not produce output, it gives EMPTY_OUTPUT_ERROR
	if parsedLog.Output == "" && parsedLog.ErrorMessage == "" {
		return service_errors.NewServiceErrorWithMessageAndError(400, "EMPTY_OUTPUT_ERROR", err)
	}

	return response.Response(200, "RoadAnswer Successfull", parsedLog)
}

// @Tags Road
// @Summary ResetPathHistory
// @Description Reset Path By Programming Language ID & Path ID
// @Accept json
// @Produce json
// @Param pathID path string true "Path ID"
// @Param programmingID path string true "Programming Language ID"
// @Param Language header string false "Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/reset/{programmingID}/{pathID} [get]
func (h *PrivateHandler) ResetPathHistory(c *fiber.Ctx) error {
	pathID := c.Params("pathID")
	programmingID := c.Params("programmingID")
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	userSession := session_store.GetSessionData(c)

	inventoryInformation, err := h.services.CommonService.GetInventoryInformation(programmingID, language)
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

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, pathID, domains.TypePath, inventoryInformation.GetFileExtension(), true)
	if err != nil {
		return err
	}
	frontendTemplateDto := h.dtoManager.RoadDTOManager.ToFrontendTemplateDto(frontendTemplate)

	return response.Response(200, "ResetPathHistory successful", frontendTemplateDto)
}
