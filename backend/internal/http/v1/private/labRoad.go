package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLabRoadRoutes(root fiber.Router) {
	root.Get("/template/:labID/:programmingID", h.GetTemplate)
	root.Get("/start/:programmingID", h.Start)
}

// @Tags LabRoadCommon
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/start/{programmingID} [get]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	programmingID := c.Params("programmingID")
	num, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid Programming ID", nil)
	}

	plInformation, err := h.services.LabRoadService.GetInventoryInformation(num)
	if err != nil {
		return response.Response(500, "Get Programming Language error", nil)
	}
	if plInformation == nil {
		return response.Response(404, "Programming Language Not Found", nil)
	}

	// Recive user session from session_store
	userSession := session_store.GetSessionData(c)

	isExsits, err := h.services.CodeService.IsImageExists(c.Context(), plInformation.GetDockerImage())
	if err != nil {
		return response.Response(500, "Docker Image Check Error", nil)
	}

	if !isExsits {
		if err := h.services.CodeService.Pull(c.Context(), plInformation.GetDockerImage()); err != nil {
			return response.Response(500, "Docker Image Pull Error", nil)
		}
	}

	// if the road has started. Log will not be created
	// Log a road start event for the user
	ok, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}

	if !ok {
		if num == 0 {
			return response.Response(200, "Invalid Programming ID", nil)
		}
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0); err != nil {
			return response.Response(500, "Error adding log", nil)
		}
	}

	isExist, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeProgrammingLanguage, domains.ContentStarted, int32(num), 0)
	if err != nil {
		return response.Response(500, "Log Check Error", nil)
	}
	if !isExist {
		return response.Response(500, "Programming Language could not started", nil)
	}

	return response.Response(200, "Progamming Language Started Successfully", nil)
}

// @Tags LabRoadCommon
// @Summary Get Template
// @Description Get Template
// @Accept json
// @Produce json
// @Param programmingID path string true "programmingID"
// @Param labID path string true "labID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/template/{labID}/{programmingID} [get]
func (h *PrivateHandler) GetTemplate(c *fiber.Ctx) error {
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

	inventoryInformation, err := h.services.LabRoadService.GetInventoryInformation(num)
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
		if codeTmp.ProgrammingID == inventoryInformation.GetID() {
			codeTemplate = codeTmp
		}
	}
	frontendContent := h.services.CodeService.CodeFrontendTemplateGenerator(inventoryInformation.GetName(), lab.GetQuest().GetFuncName(), codeTemplate.Frontend, lab.GetQuest().GetParams(), lab.GetQuest().GetReturns())

	return response.Response(200, "Template Successfully Sent", frontendContent)
}
