package dto

import "github.com/Yavuzlar/CodinLab/internal/domains"

// RoadDTOManager handles the conversion of domain road to DTOs
type RoadDTOManager struct{}

// NewRoadDTOManager creates a new instance of RoadDTOManager
func NewRoadDTOManager() RoadDTOManager {
	return RoadDTOManager{}
}

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

func (m *RoadDTOManager) ToLanguageDTO(language *domains.LanguageRoad) LanguageDTO {
	return LanguageDTO{
		Lang:        language.GetLang(),
		Title:       language.GetTitle(),
		Description: language.GetDescription(),
		Content:     language.GetContent(),
		Note:        language.GetNote(),
	}
}

func (m *RoadDTOManager) ToLanguageDTOs(languages []domains.LanguageRoad) []LanguageDTO {
	var languageDTOs []LanguageDTO
	for _, lang := range languages {
		languageDTOs = append(languageDTOs, m.ToLanguageDTO(&lang))
	}
	return languageDTOs
}

type LanguageRoadDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (m *RoadDTOManager) ToLanguageRoadDTO(language *domains.LanguageRoad) LanguageRoadDTO {
	return LanguageRoadDTO{
		Lang:        language.GetLang(),
		Title:       language.GetTitle(),
		Description: language.GetDescription(),
	}
}
func (m *RoadDTOManager) ToLanguageRoadDTOs(language []domains.LanguageRoad) []LanguageRoadDTO {
	var languageDTOs []LanguageRoadDTO
	for _, lang := range language {
		languageDTOs = append(languageDTOs, m.ToLanguageRoadDTO(&lang))
	}
	return languageDTOs
}

type PathDTO struct {
	ID         int           `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Language   []LanguageDTO `json:"languages"`
	Difficulty int           `json:"difficulty"`
	IsStarted  bool          `json:"isStarted"`
	IsFinished bool          `json:"isFinished"`
}

func (m *RoadDTOManager) ToPathDTO(path domains.Path, languages []LanguageDTO) PathDTO {
	return PathDTO{
		ID:         path.GetID(),
		Language:   languages,
		Difficulty: path.GetQuest().GetDifficulty(),
		IsFinished: path.GetIsFinished(),
		IsStarted:  path.GetIsStarted(),
	}
}

type GetRoadPathDTO struct { //this dto is used for get Roads API
	ID         int               `json:"id,omitempty"`
	Name       string            `json:"name,omitempty"`
	Language   []LanguageRoadDTO `json:"languages"`
	Difficulty int               `json:"difficulty"`
	IsStarted  bool              `json:"isStarted"`
	IsFinished bool              `json:"isFinished"`
}

func (m *RoadDTOManager) ToRoadPathDTO(path domains.Path, languages []LanguageRoadDTO) GetRoadPathDTO {
	return GetRoadPathDTO{
		ID:         path.GetID(),
		Language:   languages,
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
