package private

import (
	"strconv"

	dto "github.com/Yavuzlar/CodinLab/internal/http/dtos"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

// UserLanguageLabStatsDto represents the DTO for user language lab statistics

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:ID", h.GetLabsByID)
	root.Get("/lab/:programmingID/:labID", h.GetLabByID)
	root.Get("/labs/stats/language/:language", h.GetUserLanguageLabStats)
	root.Get("/labs/stats/general", h.GetUserGeneralLabStats)
}

// @Tags Lab
// @Summary GetUserProgrammingLanguageLabStats
// @Description Get User Programming Language Lab Statistics
// @Accept json
// @Produce json
// @Param language path string true "Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/stats/language/{language} [get]
func (h *PrivateHandler) GetUserLanguageLabStats(c *fiber.Ctx) error {
	language := c.Params("language")
	userSession := session_store.GetSessionData(c)

	stats, err := h.services.LabService.UserLanguageLabStats(userSession.UserID, language)
	if err != nil {
		return err
	}
	labStatsDTO := h.dtoManager.LabDTOManager.ToUserProgrammingLanguageLabStatsDTO(stats)

	return response.Response(200, "GetUserLanguageLabStats successful", labStatsDTO)
}

// @Tags Lab
// @Summary GetUserGeneralLabStats
// @Description Get User General Lab Statistics
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/stats/general [get]
func (h *PrivateHandler) GetUserGeneralLabStats(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)

	stats, err := h.services.LabService.UserGeneralLabStats(userSession.UserID)
	if err != nil {
		return err
	}
	dto := h.dtoManager.LabDTOManager.ToUserGeneralLabStatsDTO(stats)

	return response.Response(200, "GetUserGeneralLabStats successful", dto)
}

// @Tags Lab
// @Summary GetLabsById
// @Description Get Labs By Programming Language ID
// @Accept json
// @Produce json
// @Param ID path string true "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{ID} [get]
func (h *PrivateHandler) GetLabsByID(c *fiber.Ctx) error {
	id := c.Params("ID")
	indID, err := strconv.Atoi(id)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

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
	filteredLabs, err := h.services.LabService.GetLabsFilter(userSession.UserID, intProgrammingID, intLabID, nil, nil)
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

	return response.Response(200, "GetLab successful", labsDtoList)
}
