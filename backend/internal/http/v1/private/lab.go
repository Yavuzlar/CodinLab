package private

import (
	"fmt"
	"strings"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// UserLanguageLabStatsDto represents the DTO for user language lab statistics

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:programmingID", h.GetLabs)
	root.Get("/lab/:labID", h.GetLabByID)
	root.Get("/lab/reset/:programmingID/:labID", h.ResetLabHistory)
	root.Get("/labs/general/stats", h.GetUserLanguageLabStats)
	root.Get("/labs/difficulty/stats", h.GetUserLabDifficultyStats)
	root.Get("/labs/progress/stats", h.GetUserLabProgressStats)
	root.Post("/lab/answer/:programmingID/:labID", h.AnswerLab)
}

// @Tags Lab
// @Summary GetUserProgrammingLanguageLabStats
// @Description Get User Programming Language Lab Statistics
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/general/stats [get]
func (h *PrivateHandler) GetUserLanguageLabStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	stat, err := h.services.LabService.GetUserLanguageLabStats(userSession.UserID)
	if err != nil {
		return err
	}
	labStatsDTO := h.dtoManager.LabDTOManager.ToUserProgrammingLanguageStatsDTO(stat)

	return response.Response(200, "GetUserLanguageLabStats successful", labStatsDTO)
}

// @Tags Lab
// @Summary GetUserLabDifficultyStats
// @Description Get User Lab Difficulty Statistics
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/difficulty/stats [get]
func (h *PrivateHandler) GetUserLabDifficultyStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	stats, err := h.services.LabService.GetUserLabDifficultyStats(userSession.UserID)
	if err != nil {
		return err
	}
	dto := h.dtoManager.LabDTOManager.ToUserLabDifficultyStatsDto(stats)

	return response.Response(200, "GetUserLabDifficultyStats successful", dto)
}

// @Tags Lab
// @Summary GetUserLabProgressStats
// @Description Get User Lab Progress Statistics
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/progress/stats [get]
func (h *PrivateHandler) GetUserLabProgressStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	stats, err := h.services.LabService.GetUserLabProgressStats(userSession.UserID)
	if err != nil {
		return err
	}
	dto := h.dtoManager.LabDTOManager.ToUserLabProgressStatsDto(*stats)

	return response.Response(200, "GetUserLabDifficultyStats successful", dto)
}

// @Tags Lab
// @Summary GetLabs
// @Description Get Labs
// @Accept json
// @Produce json
// @Param Language header string false "Language"
// @Param programmingID path string true "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{programmingID} [get]
func (h *PrivateHandler) GetLabs(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	programmingID := c.Params("programmingID")
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
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

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, programmingID, "", nil, nil)
	if err != nil {
		return err
	}

	var labDTOs []dto.LabForAllDTO
	for _, labCollection := range labData {
		languageDTO := h.dtoManager.LabDTOManager.ToLanguageForAllDTO(labCollection.GetLanguages(), language)
		labDTOs = append(labDTOs, h.dtoManager.LabDTOManager.ToLabForAllDTO(labCollection, languageDTO))
	}
	labDTOs = h.dtoManager.LabDTOManager.FilterLabForAllDTOs(labDTOs)
	labsDTO := h.dtoManager.LabDTOManager.ToLabsForAllDTO(labDTOs, isExists)

	return response.Response(200, "GetLabs successful", labsDTO)
}

// @Tags Lab
// @Summary GetLabByID
// @Description Get Lab By Programming Language ID & Lab ID
// @Accept json
// @Produce json
// @Param Language header string false "Language"
// @Param labID path string true "Lab ID"
// @Param programmingID query string false "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/{labID} [get]
func (h *PrivateHandler) GetLabByID(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	labID := c.Params("labID")
	programmingID := c.Query("programmingID")
	language := h.services.UtilService.GetLanguageHeader(c.Get("Language"))

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
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

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, programmingID, labID, nil, nil)
	if err != nil {
		return err
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, labID, domains.TypeLab, inventoryInformation.GetFileExtension())
	if err != nil {
		return err
	}

	var labDTOList []dto.LabDTO
	for _, labCollection := range labData {
		languageDTO := h.dtoManager.LabDTOManager.ToLanguageDTO(labCollection.GetLanguages(), language)
		labDTOList = append(labDTOList, h.dtoManager.LabDTOManager.ToLabDTO(labCollection, languageDTO, frontendTemplate))
	}
	if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, labID, domains.TypeLab, domains.ContentStarted); err != nil {
		return err
	}

	return response.Response(200, "GetLab successful", labDTOList)
}

// @Tags Lab
// @Summary Answer
// @Description This is for answering quests.
// @Accept json
// @Produce json
// @Param labID path string false "labID"
// @Param programmingID path string true "programmingID"
// @Param answerLabDTO body dto.AnswerLabDTO true "Answer Lab DTO"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/answer/{programmingID}/{labID} [post]
func (h *PrivateHandler) AnswerLab(c *fiber.Ctx) error {
	labID := c.Params("labID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	var answerLabDTO dto.AnswerLabDTO
	if err := c.BodyParser(&answerLabDTO); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "Invalid Format", err)
	}
	if err := h.services.UtilService.Validator().ValidateStruct(answerLabDTO); err != nil {
		return err
	}

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	lab, err := h.services.LabService.GetLabByID(userSession.UserID, labID)
	if err != nil {
		return err
	}

	tmpPath, err := h.services.CodeService.UploadUserCode(c.Context(), userSession.UserID, programmingID, labID, domains.TypeLab, inventoryInformation.GetFileExtension(), answerLabDTO.UserCode)
	if err != nil {
		return err
	}

	var codeTemplate domains.CodeTemplate
	for _, codeTmp := range lab.GetQuest().GetCodeTemplates() {
		if codeTmp.GetProgrammingID() == inventoryInformation.GetID() {
			codeTemplate = codeTmp
		}
	}

	tmpContent, err := h.services.CodeService.CodeDockerTemplateGenerator(codeTemplate.GetTemplatePath(), lab.GetQuest().GetFuncName(), answerLabDTO.UserCode, lab.GetQuest().GetTests())
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

	logs, err := h.services.CodeService.RunContainerWithTar(c.Context(), inventoryInformation.GetDockerImage(), tmpPath, fmt.Sprintf("main.%v", inventoryInformation.GetFileExtension()), inventoryInformation.GetCmd(), conn)
	if err != nil {
		return err
	}
	if strings.Contains(logs, "Test Passed") {
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, programmingID, labID, domains.TypeLab, domains.ContentCompleted); err != nil {
			return err
		}
	}

	return response.Response(200, logs, nil)
}

// @Tags Lab
// @Summary ResetLabHistory
// @Description Reset Lab By Programming Language ID & Lab ID
// @Accept json
// @Produce json
// @Param labID path string true "Lab ID"
// @Param programmingID path string true "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/reset/{programmingID}/{labID} [get]
func (h *PrivateHandler) ResetLabHistory(c *fiber.Ctx) error {
	labID := c.Params("labID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	programmingInformation, err := h.services.LabRoadService.GetInventoryInformation(programmingID)
	if err != nil {
		return err
	}

	_, err = h.services.LabService.GetLabsFilter(userSession.UserID, programmingID, labID, nil, nil)
	if err != nil {
		return err
	}

	err = h.services.CodeService.DeleteFrontendTemplateHistory(userSession.UserID, programmingID, labID, domains.TypeLab, programmingInformation.GetFileExtension())
	if err != nil {
		return err
	}

	frontendTemplate, err := h.services.CodeService.GetFrontendTemplate(userSession.UserID, programmingID, labID, domains.TypeLab, programmingInformation.GetFileExtension())
	if err != nil {
		return err
	}
	frontendTemplateDto := h.dtoManager.LabDTOManager.ToFrontendTemplateDto(frontendTemplate)

	return response.Response(200, "ResetLabHistory successful", frontendTemplateDto)
}
