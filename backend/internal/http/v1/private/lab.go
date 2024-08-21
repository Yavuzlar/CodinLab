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
	root.Get("/labs/stats/:language/:userID", h.GetUserLanguageLabStats)
	root.Get("/labs/stats/:userID", h.GetUserGeneralLabStats)
}

// @Tags Lab
// @Summary GetUserProgrammingLanguageLabStats
// @Description Get User Programming Language Lab Statistics
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param language path string true "Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/stats/{language}/{userID} [get]
func (h *PrivateHandler) GetUserLanguageLabStats(c *fiber.Ctx) error {
	userID := c.Params("userID")
	language := c.Params("language")

	stats, err := h.services.LabService.UserLanguageLabStats(userID, language)
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
// @Param userID path string true "User ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/stats/{userID} [get]
func (h *PrivateHandler) GetUserGeneralLabStats(c *fiber.Ctx) error {
	userID := c.Params("userID")
	stats, err := h.services.LabService.UserGeneralLabStats(userID)
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
	ID := c.Params("ID")
	indID, err := strconv.Atoi(ID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

	filteredLabs, err := h.services.LabService.GetLabsFilter(userSession.UserID, indID, 0, "", "") //buraya bakilsin (false tan dolayi veriler gelmeyebilir) !!11!1!!11
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
	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intProgrammingID, intLabID, "", "")
	if err != nil {
		return err
	}
	if len(labData) == 0 {
		return response.Response(404, "Lab not found", labData)
	}

	return response.Response(200, "GetLab successful", labData)
}
