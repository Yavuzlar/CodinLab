package private

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/internal/http/response"
	"github.com/Yavuzlar/CodinLab/internal/http/session_store"
	"github.com/gofiber/fiber/v2"
)

type StartDTO struct {
	ProgrammingID int32 `json:"programmingID" validate:"required"`
}

type LanguageDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Note        string `json:"note"`
}

type PathDTO struct {
	ID         int           `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Language   []LanguageDTO `json:"languages"`
	Difficulty int           `json:"difficulty"`
	IsStarted  bool          `json:"isStarted"`
	IsFinished bool          `json:"isFinished"`
}

type RoadDTO struct {
	Name     string    `json:"name"`
	IconPath string    `json:"iconPath"`
	Paths    []PathDTO `json:"paths"`
}

func (h *PrivateHandler) initRoadRoutes(root fiber.Router) {
	roadRoute := root.Group("/road")
	roadRoute.Post("/start", h.Start)
	roadRoute.Get("/:roadID", h.GetAllRoads)
	roadRoute.Get("/:roadID/:pathID", h.GetPath)
}

// @Tags Road
// @Summary Start
// @Description Start
// @Accept json
// @Produce json
// @Param start body StartDTO true "Start"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/road/start [post]
func (h *PrivateHandler) Start(c *fiber.Ctx) error {
	var start StartDTO
	if err := c.BodyParser(&start); err != nil {
		return err
	}

	// NEED ROAD SERVICE FOR BOTTOM
	// We have to get the road that will start according to the id
	// In that road struct we will change isStarted to true.
	// And we need road's docker image for DocerService

	// Need Road Service For -> Road title & Docker Image For Log
	// With road service we will get road by road id and recive docker image.

	// Recive user session from session_store
	userSession := session_store.GetSessionData(c)

	// Need Spesific imageReference for this. If it's wrong it wont work
	isExsits, err := h.services.DockerService.IsImageExists(c.Context(), "golang:latest")
	if err != nil {
		return err
	}

	if !isExsits {
		// Downloads Spesific Image. This golang fetched from road.json
		if err := h.services.DockerService.Pull(c.Context(), "golang:latest"); err != nil {
			return err
		}
	}

	// if the road has started. Log will not be created
	// Log a road start event for the user
	ok, err := h.services.LogService.IsExists(c.Context(), userSession.UserID, domains.TypeRoad, domains.ContentStarted, start.ProgrammingID, 0)
	if err != nil {
		return err
	}

	if !ok {
		if start.ProgrammingID == 0 {
			return response.Response(200, "Invalid Programming ID", nil)
		}
		if err := h.services.LogService.Add(c.Context(), userSession.UserID, domains.TypeRoad, domains.ContentStarted, start.ProgrammingID, 0); err != nil {
			return err
		}
	}

	// if the "Road Started Successfully" message recived. The frontend will redirect the user to the spesific road.
	return response.Response(200, "Road Started Successfully", nil)
}

// @Tags Road
// @Summary GetRoads
// @Description Get All Roads
// @Accept json
// @Produce json
// @Param roadID path string true "roadID"
// @Success 200 {object} response.BaseResponse{data=RoadDTO}
// @Router /private/road/{roadID} [get]
func (h *PrivateHandler) GetAllRoads(c *fiber.Ctx) error {
	// NEED ROAD SERVICE FOR BOTTOM
	// We have to get all roads and send them to the frontend.

	// Recive user session from session_store

	userSession := session_store.GetSessionData(c)

	// Declare and assign roadID variable
	roadID := c.Params("roadID")
	num, err := strconv.Atoi(roadID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	// Get all roads
	roads, err := h.services.RoadService.GetRoadFilter(userSession.UserID, num, 0, false, false)
	if err != nil {
		return err
	}

	var pathDTOs []PathDTO
	var roadDTO *RoadDTO
	for _, road := range roads {
		for _, path := range road.Paths {
			var languageDTOs []LanguageDTO
			for _, lang := range path.Languages {
				languageDTO := LanguageDTO{
					Lang:        lang.Lang,
					Title:       lang.Title,
					Description: lang.Description,
					Content:     lang.Content,
					Note:        lang.Note,
				}
				languageDTOs = append(languageDTOs, languageDTO)
			}
			pathDTO := PathDTO{
				ID:         path.ID,
				Language:   languageDTOs,
				Difficulty: path.Quest.Difficulty,
				IsStarted:  path.IsStarted,
				IsFinished: path.IsFinished,
			}
			pathDTOs = append(pathDTOs, pathDTO)
		}
		roadDTO = &RoadDTO{
			Name:     road.Name,
			IconPath: road.IconPath,
			Paths:    pathDTOs,
		}
	}

	return response.Response(200, "GetRoads successful", roadDTO)
}

// @Tags Road
// @Summary GetPathByID
// @Description Get Path By ID
// @Accept json
// @Produce json
// @Param roadID path string true "Road ID"
// @Param pathID path string true "Path ID"
// @Success 200 {object} response.BaseResponse{data=PathDTO}
// @Router /private/road/{roadID}/{pathID} [get]
func (h *PrivateHandler) GetPath(c *fiber.Ctx) error {
	programmingID := c.Params("roadID")
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	pathID := c.Params("pathID")

	intPathID, err := strconv.Atoi(pathID)
	if err != nil {
		return response.Response(400, "Invalid ID", nil)
	}

	session := session_store.GetSessionData(c)
	userID := session.UserID

	roadData, err := h.services.RoadService.GetRoadFilter(userID, intProgrammingID, intPathID, false, false)
	if err != nil {
		return err
	}
	if len(roadData) == 0 {
		return response.Response(404, "Path not found", nil)
	}
	roads := &roadData[0]
	var pathDTO PathDTO
	for _, path := range roads.Paths {
		var langugaeDTOs []LanguageDTO
		for _, lang := range path.Languages {
			langugaeDTO := LanguageDTO{
				Lang:        lang.Lang,
				Title:       lang.Title,
				Description: lang.Description,
				Content:     lang.Content,
				Note:        lang.Note,
			}
			langugaeDTOs = append(langugaeDTOs, langugaeDTO)
		}
		pathDTO = PathDTO{
			Name:       roads.Name,
			Language:   langugaeDTOs,
			Difficulty: path.Quest.Difficulty,
			IsStarted:  path.IsStarted,
			IsFinished: path.IsFinished,
		}
	}

	return response.Response(200, "Path Retrieved Successfully", pathDTO)
}
