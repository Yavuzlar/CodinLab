package private

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

type UserLevelDTO struct {
	Level           int                 `json:"level"`
	TotalPoints     int32               `json:"total_points"`
	LevelPercentage int32               `json:"level_percentage"`
	Languages       []domains.LanguageL `json:"languages"`
}

type InventoryDto struct {
	Id       int
	Name     string
	IconPath string
}

type DevelopmentDto struct {
	RoadPercentage int32
	LabPercentage  int32
}

type AdvancementDTO struct {
	LangID         int
	Name           string
	IconPath       string
	RoadPercentage int32
	LabPercentage  int32
}

func (h *PrivateHandler) initHomeRoutes(root fiber.Router) {
	root.Get("/home/inventories", h.GetInventories)
	root.Get("/home/level", h.GetUserLevel)
	root.Get("/home/development", h.GetUserDevelopment)
	root.Get("/home/advancement", h.GetUserAdvancement)
	root.Get("/home/welcome", h.GetWelcomeContent)
	root.Get("/home/road", h.GetRoadContent)
	root.Get("/home/lab", h.GetLabContent)
	// initialize routes
	// Buraya yeni route'lar eklenecek lütfen Swagger'da belirtmeyi unutmayın
}

// @Tags Home
// @Summary GetUserLevel
// @Description Get User Level
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/level [get]
func (h *PrivateHandler) GetUserLevel(c *fiber.Ctx) error {

	userSession := session_store.GetSessionData(c)
	userLevel, err := h.services.HomeService.GetUserLevel(c.Context(), userSession.UserID)
	if err != nil {
		return err
	}
	userLevelDto := UserLevelDTO{
		Level:           userLevel.Level(),
		TotalPoints:     userLevel.TotalPoints(),
		LevelPercentage: userLevel.LevelPercentage(),
		Languages:       userLevel.Languages(),
	}
	return response.Response(200, "GetUserLevel successful", userLevelDto)
}

// @Tags Home
// @Summary GetInventories
// @Description Get Inventories
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/inventories [get]
func (h *PrivateHandler) GetInventories(c *fiber.Ctx) error {

	inventoryData, err := h.services.HomeService.GetInventory(c.Context())
	if err != nil {
		return err
	}

	if len(inventoryData) == 0 {
		return response.Response(404, "Inventories not found", nil)
	}

	return response.Response(200, "GetInventories successful", inventoryData)
}

// @Tags Home
// @Summary GetUserDevelopment
// @Description Get User Development
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/development [get]
func (h *PrivateHandler) GetUserDevelopment(c *fiber.Ctx) error {

	userSession := session_store.GetSessionData(c)

	userDevelopment, err := h.services.HomeService.GetUserDevelopment(c.Context(), userSession.UserID)

	if err != nil {
		return err
	}

	userDevelopmentDto := DevelopmentDto{
		RoadPercentage: userDevelopment.RoadPercentage(),
		LabPercentage:  userDevelopment.LabPercentage(),
	}

	return response.Response(200, "GetUserDevelopment successful", userDevelopmentDto)
}

// @Tags Home
// @Summary GetUserAdvancement
// @Description Get User Advancement
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/advancement [get]
func (h *PrivateHandler) GetUserAdvancement(c *fiber.Ctx) error {
	userSession := session_store.GetSessionData(c)
	userAdvancement, err := h.services.HomeService.GetUserAdvancement(c.Context(), userSession.UserID)

	if err != nil {
		return err
	}

	var advancementDTOs []AdvancementDTO
	for _, advancement := range userAdvancement {
		advancementDTO := AdvancementDTO{
			LangID:         advancement.LangID(),
			Name:           advancement.Name(),
			IconPath:       advancement.IconPath(),
			RoadPercentage: advancement.RoadPercentage(),
			LabPercentage:  advancement.LabPercentage(),
		}
		advancementDTOs = append(advancementDTOs, advancementDTO)
	}

	return response.Response(200, "GetUserAdvancement successful", advancementDTOs)
}

// @Tags Home
// @Summary GetWelcomeContent
// @Description Get Welcome Content
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/welcome [get]
func (h *PrivateHandler) GetWelcomeContent(c *fiber.Ctx) error {

	content, err := h.services.HomeService.GetWelcomeContent()

	if err != nil {
		return err
	}

	return response.Response(200, "GetWelcomeContent successful", content)
}

// @Tags Home
// @Summary GetLabContent
// @Description Get Lab Content
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/lab [get]
func (h *PrivateHandler) GetLabContent(c *fiber.Ctx) error {

	content, err := h.services.HomeService.GetLabContent()

	if err != nil {
		return err
	}

	return response.Response(200, "GetLabContent successful", content)
}

// @Tags Home
// @Summary GetRoadContent
// @Description Get Road Content
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/home/road [get]
func (h *PrivateHandler) GetRoadContent(c *fiber.Ctx) error {

	content, err := h.services.HomeService.GetRoadContent()

	if err != nil {
		return err
	}

	return response.Response(200, "GetRoadContent successful", content)
}
