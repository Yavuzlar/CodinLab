package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

// UserLanguageLabStatsDto represents the DTO for user language lab statistics

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:programmingID", h.GetLabsByID)
	root.Get("/lab/:programmingID/:labID", h.GetLabByID)
	root.Get("/labs/general/stats", h.GetUserLanguageLabStats)
	root.Get("/labs/difficulty/stats", h.GetUserLabDifficultyStats)
	root.Get("/labs/progress/stats", h.GetUserLabProgressStats)
	root.Get("/lab/data", h.AddDummyLabData)
	root.Post("/lab/answer", h.AnswerLab)
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

	stats, err := h.services.LabService.GetUserLanguageLabStats(userSession.UserID)
	if err != nil {
		return err
	}
	labStatsDTO := h.dtoManager.LabDTOManager.ToUserProgrammingLanguageStatsDTO(stats)
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
	dto := h.dtoManager.LabDTOManager.ToUserLabProgressStatsDto(stats)

	return response.Response(200, "GetUserLabDifficultyStats successful", dto)
}

// @Tags Lab
// @Summary GetLabs
// @Description Get Labs By Programming Language ID
// @Accept json
// @Produce json
// @Param programmingID path string true "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{programmingID} [get]
func (h *PrivateHandler) GetLabsByID(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	indID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

	isExist, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(indID), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}
	if !isExist {
		return response.Response(500, "Programming Language could not started", nil)
	}

	filteredLabs, err := h.services.LabService.GetLabsFilter(userSession.UserID, indID, 0, nil, nil)
	if err != nil {
		return err
	}
	var labsDtoList []dto.LabsDTO
	for _, labCollection := range filteredLabs {
		var labDTOList []dto.LabDTO
		for _, lab := range labCollection.GetLabs() {
			languageDTOs := h.dtoManager.LabDTOManager.ToLanguageDTOs(lab.GetLanguages())
			labDTOList = append(labDTOList, h.dtoManager.LabDTOManager.ToLabDTO(lab, languageDTOs))

		}
		labsDtoList = append(labsDtoList, h.dtoManager.LabDTOManager.ToLabsDTO(labCollection, labDTOList))
	}
	if len(labsDtoList) == 0 {
		return response.Response(404, "Labs not found", labsDtoList)
	}
	return response.Response(200, "GetLabs successful", labsDtoList)
}

// @Tags Lab
// @Summary GetLabByID
// @Description Get Lab By Programming Language ID & Lab ID
// @Accept json
// @Produce json
// @Param programmingID path string true "Programming Language ID"
// @Param labID path string true "Lab ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/{programmingID}/{labID} [get]
func (h *PrivateHandler) GetLabByID(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	labID := c.Params("labID")

	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Lang ID", nil)
	}

	intLabID, err := strconv.Atoi(labID)
	if err != nil {
		return response.Response(400, "Invalid Labs ID", nil)
	}

	userSession := session_store.GetSessionData(c)
	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intProgrammingID, intLabID, nil, nil)
	if err != nil {
		return err
	}
	if len(labData) == 0 {
		return response.Response(404, "Lab not found", nil)
	}

	var labsDtoList []dto.LabsDTO
	for _, labCollection := range labData {
		var labDTOList []dto.LabDTO
		for _, lab := range labCollection.GetLabs() {
			languageDTOs := h.dtoManager.LabDTOManager.ToLanguageDTOs(lab.GetLanguages())
			labDTOList = append(labDTOList, h.dtoManager.LabDTOManager.ToLabDTO(lab, languageDTOs))

		}
		labsDtoList = append(labsDtoList, h.dtoManager.LabDTOManager.ToLabsDTO(labCollection, labDTOList))
	}

	return response.Response(200, "GetLab successful", labsDtoList)
}

// @Tags Lab
// @Summary DummyLogData
// @Description Add dummy data for testing
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/data [get]
func (h *PrivateHandler) AddDummyLabData(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	// Dummy Data for testing
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeLab, domains.ContentStarted, 1, 2)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeLab, domains.ContentStarted, 2, 2)
	h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeLab, domains.ContentCompleted, 1, 2)

	return response.Response(200, "Dummy Data Added", nil)
}

// @Tags Lab
// @Summary Answer
// @Description This is for answering quests.
// @Accept json
// @Produce json
// @Param answerLabDTO body dto.AnswerLabDTO true "Answer Lab DTO"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/answer [post]
func (h *PrivateHandler) AnswerLab(c *fiber.Ctx) error {
	var answerLabDTO dto.AnswerLabDTO
	if err := c.BodyParser(&answerLabDTO); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "Invalid Format", err)
	}

	userSession := session_store.GetSessionData(c)
	roadInformation, err := h.services.RoadService.GetRoadInformation(int32(answerLabDTO.ProgrammingID))
	if err != nil {
		return response.Response(500, "Road Information Error", nil)
	}
	if roadInformation == nil {
		return response.Response(404, "Road Not Found", nil)
	}

	lab, err := h.services.LabService.GetLabByID(userSession.UserID, answerLabDTO.ProgrammingID, answerLabDTO.LabID)
	if err != nil {
		return response.Response(404, "Error While Getting Lab", nil)
	}
	if lab == nil {
		return response.Response(404, "Lab Not Found", nil)
	}

	tmpPath, err := h.services.CodeService.UploadUserCode(c.Context(), userSession.UserID, answerLabDTO.ProgrammingID, answerLabDTO.LabID, domains.TypeLab, roadInformation.GetFileExtension(), answerLabDTO.UserCode)
	if err != nil {
		return err
	}

	tmpContent, err := h.services.LabService.CodeTemplateGenerator(roadInformation.GetName(), roadInformation.GetTemplatePath(), answerLabDTO.UserCode, lab.GetQuest().GetFuncName(), lab.GetQuest().GetTests())
	if err != nil {
		return err
	}

	if err := h.services.CodeService.CreateFileAndWrite(tmpPath, tmpContent); err != nil {
		return err
	}

	logs, err := h.services.CodeService.RunContainerWithTar(c.Context(), roadInformation.GetDockerImage(), tmpPath, roadInformation.GetCmd())
	if err != nil {
		return err
	}

	if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeLab, domains.ContentStarted, int32(answerLabDTO.ProgrammingID), int32(answerLabDTO.LabID)); err != nil {
		return response.Response(500, "Docker Image Pull Error", nil)
	}

	return response.Response(200, logs, nil)
}
