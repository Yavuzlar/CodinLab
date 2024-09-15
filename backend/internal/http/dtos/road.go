package dto

import "github.com/Yavuzlar/CodinLab/internal/domains"

// RoadDTOManager handles the conversion of domain road to DTOs
type RoadDTOManager struct{}

// NewRoadDTOManager creates a new instance of RoadDTOManager
func NewRoadDTOManager() RoadDTOManager {
	return RoadDTOManager{}
}

type LanguageRoadDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Content     string `json:"content"`
}

func (m *RoadDTOManager) ToLanguageRoadDTO(languageRoads []domains.LanguageRoad, langauge string) LanguageRoadDTO {
	var languageRoadDto LanguageRoadDTO

	for _, languageRoad := range languageRoads {
		if languageRoad.GetLang() == langauge {
			languageRoadDto = LanguageRoadDTO{
				Lang:        languageRoad.GetLang(),
				Title:       languageRoad.GetTitle(),
				Description: languageRoad.GetDescription(),
				Note:        languageRoad.GetNote(),
				Content:     languageRoad.GetContent(),
			}
		}
	}

	return languageRoadDto
}

type PathDTO struct {
	ID         int             `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Language   LanguageRoadDTO `json:"language"`
	Template   string          `json:"template,omitempty"`
	Difficulty int             `json:"difficulty"`
	IsStarted  bool            `json:"isStarted"`
	IsFinished bool            `json:"isFinished"`
}

func (m *RoadDTOManager) ToPathDTO(path domains.Path, language LanguageRoadDTO, template string) PathDTO {
	return PathDTO{
		ID:         path.GetID(),
		Language:   language,
		Template:   template,
		Difficulty: path.GetQuest().GetDifficulty(),
		IsFinished: path.GetIsFinished(),
		IsStarted:  path.GetIsStarted(),
	}
}

type GetRoadPathDTO struct {
	ID         int             `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Language   LanguageRoadDTO `json:"languages"`
	Difficulty int             `json:"difficulty"`
	IsStarted  bool            `json:"isStarted"`
	IsFinished bool            `json:"isFinished"`
}

func (m *RoadDTOManager) ToRoadPathDTO(path domains.Path, language LanguageRoadDTO) GetRoadPathDTO {
	return GetRoadPathDTO{
		ID:         path.GetID(),
		Language:   language,
		Difficulty: path.GetQuest().GetDifficulty(),
		IsFinished: path.GetIsFinished(),
		IsStarted:  path.GetIsStarted(),
	}
}

type RoadDTO struct {
	Name       string    `json:"name"`
	IconPath   string    `json:"iconPath"`
	Paths      []PathDTO `json:"paths"`
	IsStarted  bool      `json:"isStarted"`
	IsFinished bool      `json:"isFinished"`
}

func (m *RoadDTOManager) ToRoadDTO(road domains.Road, paths []PathDTO) RoadDTO {
	return RoadDTO{
		Name:       road.GetName(),
		IconPath:   road.GetIconPath(),
		Paths:      paths,
		IsStarted:  road.GetIsStarted(),
		IsFinished: road.GetIsFinished(),
	}
}

type GetRoadDTO struct {
	Name       string           `json:"name"`
	IconPath   string           `json:"iconPath"`
	IsStarted  bool             `json:"isStarted"`
	IsFinished bool             `json:"isFinished"`
	Paths      []GetRoadPathDTO `json:"paths"`
}

func (m *RoadDTOManager) ToGetRoadDTO(road domains.Road, paths []GetRoadPathDTO) GetRoadDTO { //for get Roads API
	return GetRoadDTO{
		Name:       road.GetName(),
		IconPath:   road.GetIconPath(),
		IsStarted:  road.GetIsStarted(),
		IsFinished: road.GetIsFinished(),
		Paths:      paths,
	}
}

type RoadStatsDTO struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	IconPath       string  `json:"iconPath"`
	TotalRoads     int     `json:"totalRoads"`
	CompletedRoads int     `json:"completedRoads"`
	Percentage     float32 `json:"percentage"`
}

func (m *RoadDTOManager) ToRoadStatsDTO(stats []domains.RoadStats) []RoadStatsDTO {
	var roadStatsDTO []RoadStatsDTO
	for _, stat := range stats {
		roadStats := RoadStatsDTO{
			ID:             stat.GetID(),
			Name:           stat.GetName(),
			IconPath:       stat.GetIconPath(),
			TotalRoads:     stat.GetTotalRoads(),
			CompletedRoads: stat.GetCompletedRoads(),
			Percentage:     stat.GetPercentage(),
		}
		roadStatsDTO = append(roadStatsDTO, roadStats)
	}

	return roadStatsDTO
}

type RoadProgressDTO struct {
	Progress  float32 `json:"progress"`
	Completed float32 `json:"completed"`
}

func (m *RoadDTOManager) ToRoadProgressDTO(stats domains.RoadProgressStats) RoadProgressDTO {
	return RoadProgressDTO{
		Completed: stats.GetCompleted(),
		Progress:  stats.GetProgress(),
	}
}

type AnswerRoadDTO struct {
	/* ProgrammingID int    `json:"programmindID"`
	PathID        int    `json:"pathID"` */
	UserCode string `json:"userCode"`
}

func (m *RoadDTOManager) ToFrontendTemplateDto(frontendTemplate string) RoadFrontendTemplateDto {
	return RoadFrontendTemplateDto{
		Template: frontendTemplate,
	}
}

type RoadFrontendTemplateDto struct {
	Template string `json:"template"`
}
