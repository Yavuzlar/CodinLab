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

func (h *PrivateHandler) initLabRoutes(root fiber.Router) {
	root.Get("/labs/:id", h.GetLabsByID)
	root.Get("/lab/:langId/:labId", h.GetLabByID)
	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
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
// @Summary GetLabById
// @Description Get Lab By Lang ID & Lab ID
// @Accept json
// @Produce json
// @Param langId path string true "Lang ID"
// @Param labId path string true "Lab ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/lab/{langId}/{labId} [get]
func (h *PrivateHandler) GetLabByID(c *fiber.Ctx) error {
	langId := c.Params("langId")
	labId := c.Params("labId")

	intLangId, err := strconv.Atoi(langId)
	if err != nil {
		return response.Response(400, "Invalid Lang ID", nil)
	}
	intLabId, err := strconv.Atoi(labId)
	if err != nil {
		return response.Response(400, "Invalid Labs ID", nil)
	}

	userSession := session_store.GetSessionData(c)

	labData, err := h.services.LabService.GetLabsFilter(userSession.UserID, intLangId, intLabId, "", "")
	if err != nil {
		return err
	}

	if len(labData) == 0 {
		return response.Response(404, "Lab not found", labData)
	}

	return response.Response(200, "GetLab successful", labData)
}
