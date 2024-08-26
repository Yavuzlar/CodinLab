package private

import (
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initHomeRoutes(root fiber.Router) {
	homeRoute := root.Group("/home")
	homeRoute.Get("/inventories", h.GetInventories)
	homeRoute.Get("/level", h.GetUserLevel)
	homeRoute.Get("/development", h.GetUserDevelopment)
	homeRoute.Get("/advancement", h.GetUserAdvancement)
	//homeRoute.Get("/welcome", h.GetWelcomeContent)
	//homeRoute.Get("/road", h.GetRoadContent)
	//homeRoute.Get("/lab", h.GetLabContent)
	// initialize routes
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
	languageLevelDto := h.dtoManager.HomeDTOManager.ToLanguageLevelDTO(userLevel.Languages())
	userLevelDto := h.dtoManager.HomeDTOManager.ToUserLevelDTO(userLevel, languageLevelDto)

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
	inventoryDTOs := h.dtoManager.HomeDTOManager.ToInventoryDTOs(inventoryData)

	return response.Response(200, "GetInventories successful", inventoryDTOs)
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
	userDevelopmentDto := h.dtoManager.HomeDTOManager.ToDevelopmentDTO(userDevelopment)

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
	advancementDTOs := h.dtoManager.HomeDTOManager.ToAdvancementDTOs(userAdvancement)

	return response.Response(200, "GetUserAdvancement successful", advancementDTOs)
}

/* // @Tags Home
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
} */

/* // @Tags Home
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
} */
