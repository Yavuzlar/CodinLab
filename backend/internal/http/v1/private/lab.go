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
)

// UserLanguageLabStatsDto represents the DTO for user language lab statistics

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs", h.GetLabs)
	root.Get("/lab/:labID", h.GetLabByID)
	root.Get("/labs/general/stats", h.GetUserLanguageLabStats)
	root.Get("/labs/difficulty/stats", h.GetUserLabDifficultyStats)
	root.Get("/labs/progress/stats", h.GetUserLabProgressStats)
	root.Post("/lab/answer/:programmingID/:labID", h.AnswerLab)
	root.Get("/lab/template/:programmingID/:labID", h.GetLabTemplate)
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
// @Param programmingID query string false "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/ [get]
func (h *PrivateHandler) GetLabs(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	programmingID := c.Query("programmingID")

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming Language ID", nil)
	}

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, 0, intProgrammingID, nil, nil)
	if err != nil {
		return err
	}

	var labDTOList []dto.LabDTO
	for _, labCollection := range labData {
		languageDTOs := h.dtoManager.LabDTOManager.ToLanguageDTOs(labCollection.GetLanguages())
		labDTOList = append(labDTOList, h.dtoManager.LabDTOManager.ToLabDTO(labCollection, languageDTOs))
	}
	if len(labDTOList) == 0 {
		return response.Response(404, "Labs not found", nil)
	}

	return response.Response(200, "GetLabs successful", labDTOList)
}

// @Tags Lab
// @Summary GetLabByID
// @Description Get Lab By Programming Language ID & Lab ID
// @Accept json
// @Produce json
// @Param labID path string true "Lab ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/{labID} [get]
func (h *PrivateHandler) GetLabByID(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	labID := c.Params("labID")

	intLabID, err := strconv.Atoi(labID)
	if err != nil {
		return response.Response(400, "Invalid Labs ID", nil)
	}

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intLabID, 0, nil, nil)
	if err != nil {
		return err
	}
	if len(labData) == 0 {
		return response.Response(404, "Lab not found", nil)
	}

	//FIXME: FRONTEND TEMPLATE
	// Code Service bir service yaz o sana query ile verilen dilin frontend template'ini  döndürsün. Eğer verilmediyse 0. indexi yollasın zaten boş olamaz.
	var labDTOList []dto.LabDTO
	for _, labCollection := range labData {
		languageDTOs := h.dtoManager.LabDTOManager.ToLanguageDTOs(labCollection.GetLanguages())
		labDTOList = append(labDTOList, h.dtoManager.LabDTOManager.ToLabDTO(labCollection, languageDTOs))
	}
	if len(labDTOList) == 0 {
		return response.Response(404, "Labs not found", nil)
	}

	return response.Response(200, "GetLab successful", labDTOList)
}

// @Tags Lab
// @Summary Answer
// @Description This is for answering quests.
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Param labID path string false "labID"
// @Param answerLabDTO body dto.AnswerLabDTO true "Answer Lab DTO"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/answer/{programmingID}/{labID} [post]
func (h *PrivateHandler) AnswerLab(c *fiber.Ctx) error {
	var answerLabDTO dto.AnswerLabDTO
	if err := c.BodyParser(&answerLabDTO); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "Invalid Format", err)
	}

	labID := c.Params("labID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	labIDInt, err := strconv.Atoi(labID)
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

	lab, err := h.services.LabService.GetLabByID(userSession.UserID, labIDInt)
	if err != nil {
		return response.Response(404, "Error While Getting Lab", err)
	}
	if lab == nil {
		return response.Response(404, "Lab Not Found", nil)
	}

	tmpPath, err := h.services.CodeService.UploadUserCode(c.Context(), userSession.UserID, num, labIDInt, domains.TypeLab, inventoryInformation.GetFileExtension(), answerLabDTO.UserCode)
	if err != nil {
		return err
	}

	var codeTemplate domains.CodeTemplate
	for _, codeTmp := range lab.GetQuest().GetCodeTemplates() {
		if codeTmp.GetProgrammingID() == inventoryInformation.GetID() {
			codeTemplate = codeTmp
		}
	}

	tmpContent, err := h.services.CodeService.CodeDockerTemplateGenerator(codeTemplate.GetTemplate(), codeTemplate.GetCheck(), codeTemplate.GetSuccess(), answerLabDTO.UserCode, lab.GetQuest().GetFuncName(), lab.GetQuest().GetTests(), lab.GetQuest().GetReturns())
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
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeLab, domains.ContentCompleted, int32(num), int32(labIDInt)); err != nil {
			return response.Response(500, "Docker Image Pull Error", nil)
		}
	}

	return response.Response(200, logs, nil)
}

// @Tags Lab
// @Summary Get Lab Template
// @Description Get Lab Template
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Param labID path string false "labID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/template/{programmingID}/{labID} [get]
func (h *PrivateHandler) GetLabTemplate(c *fiber.Ctx) error {
	labID := c.Params("labID")
	programmingID := c.Params("programmingID")
	userSession := session_store.GetSessionData(c)

	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	labIDInt, err := strconv.Atoi(labID)
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

	lab, err := h.services.LabService.GetLabByID(userSession.UserID, labIDInt)
	if err != nil {
		return response.Response(404, "Error While Getting Lab", err)
	}
	if lab == nil {
		return response.Response(404, "Lab Not Found", nil)
	}

	var codeTemplate domains.CodeTemplate
	for _, codeTmp := range lab.GetQuest().GetCodeTemplates() {
		if codeTmp.GetProgrammingID() == inventoryInformation.GetID() {
			codeTemplate = codeTmp
		}
	}
	frontendContent := h.services.CodeService.CodeFrontendTemplateGenerator(inventoryInformation.GetName(), lab.GetQuest().GetFuncName(), codeTemplate.GetFrontend(), lab.GetQuest().GetParams(), lab.GetQuest().GetReturns(), codeTemplate.GetQuestImports())

	return response.Response(200, "Template Successfully Sent", frontendContent)
}
