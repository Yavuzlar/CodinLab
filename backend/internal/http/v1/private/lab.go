package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

// Language Dto for transfer
type LanguageDto struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note,omitempty"`
	Hint        string `json:"hint,omitempty"`
}

// Lab Dto for transfer
type LabDto struct {
	ID         int           `json:"id"`
	Languages  []LanguageDto `json:"languages"`
	IsStarted  string        `json:"isStarted"`
	IsFinished string        `json:"isFinished"`
	Difficulty int           `json:"difficulty"`
}

// Labs Dto for transfer
type LabsDto struct {
	ID       int
	Name     string   `json:"name"`
	IconPath string   `json:"iconPath"`
	Labs     []LabDto `json:"labs"`
}

// UserLanguageLabStatsDto represents the DTO for user language lab statistics
type UserLanguageLabStatsDto struct {
	TotalLabs     int     `json:"totalLabs"`
	CompletedLabs int     `json:"completedLabs"`
	Percentage    float64 `json:"Percentage"`
}

// UserGeneralLabStatsDto represents the DTO for user general lab statistics
type UserGeneralLabStatsDto struct {
	TotalLabs        int     `json:"totalLabs"`
	TotalPercentage  float64 `json:"TotalPercentage"`
	EasyLabs         int     `json:"easyLabs"`
	EasyPercentage   float64 `json:"easyPercentage"`
	MediumLabs       int     `json:"mediumlabs"`
	MediumPercentage float64 `json:"mediumPercentage"`
	HardLabs         int     `json:"hardLabs"`
	HardPercentage   float64 `json:"hardPercentage"`
}

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:ID", h.GetLabsByID)
	root.Get("/lab/:programmingID/:labID", h.GetLabByID)
	root.Get(("/labs/stats/:language/:userID"), h.GetUserLanguageLabStats)
	root.Get(("/labs/stats/:userID"), h.GetUserGeneralLabStats)

	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
}

// @Tags Lab
// @Summary GetUserLanguageLabStats
// @Description Get user language lab statistics
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

	dto := UserLanguageLabStatsDto{
		TotalLabs:     stats.TotalLabs,
		CompletedLabs: stats.CompletedLabs,
		Percentage:    stats.Percentage,
	}
	return response.Response(200, "GetUserLanguageLabStats successful", dto)
}

// @Tags Lab
// @Summary GetUserGeneralLabStats
// @Description Get user general lab statistics
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

	dto := UserGeneralLabStatsDto{
		TotalLabs:        stats.TotalLabs,
		TotalPercentage:  stats.TotalPercentage,
		EasyLabs:         stats.EasyLabs,
		EasyPercentage:   stats.EasyPercentage,
		MediumLabs:       stats.MediumLabs,
		MediumPercentage: stats.MediumPercentage,
		HardLabs:         stats.HardLabs,
		HardPercentage:   stats.HardPercentage,
	}

	return response.Response(200, "GetUserGeneralLabStats successful", dto)
}

// @Tags Lab
// @Summary GetLabsById
// @Description Get Labs By Lang ID
// @Accept json
// @Produce json
// @Param id path string true "Labs ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/labs/{id} [get]
func (h *PrivateHandler) GetLabsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}
	userSession := session_store.GetSessionData(c)

	filteredLabs, err := h.services.LabService.GetLabsFilter(userSession.UserID, intId, 0, "", "")
	if err != nil {
		return err
	}

	var labsDtoList []LabsDto
	for _, labCollection := range filteredLabs {
		var labDtoList []LabDto
		for _, lab := range labCollection.Labs {
			var languageDtoList []LanguageDto
			for _, language := range lab.Languages {
				languageDto := LanguageDto{
					Lang:        language.Lang,
					Title:       language.Title,
					Description: language.Description,
					Note:        language.Note,
					Hint:        language.Hint,
				}
				languageDtoList = append(languageDtoList, languageDto)
			}
			labDto := LabDto{
				ID:         lab.ID,
				Languages:  languageDtoList,
				IsStarted:  lab.IsStarted,
				IsFinished: lab.IsFinished,
				Difficulty: lab.Quest.Difficulty,
			}
			labDtoList = append(labDtoList, labDto)
		}
		labsDto := LabsDto{
			ID:       labCollection.ID,
			Name:     labCollection.Name,
			IconPath: labCollection.IconPath,
			Labs:     labDtoList,
		}
		labsDtoList = append(labsDtoList, labsDto)
	}

	if len(labsDtoList) == 0 {
		return response.Response(404, "Labs not found", labsDtoList)
	}

	return response.Response(200, "GetLabs successful", labsDtoList)
}

// @Tags Lab
// @Summary GetLabByID
// @Description Get Lab By Lang ID & Lab ID
// @Accept json
// @Produce json
// @Param programmingID path string true "Lang ID"
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
