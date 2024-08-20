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

func (m *RoadDTOManager) ToLanguageDTO(language *domains.LanguageR) LanguageDTO {
	return LanguageDTO{
		Lang:        language.Lang,
		Title:       language.Title,
		Description: language.Description,
		Content:     language.Content,
		Note:        language.Note,
	}
}

func (m *RoadDTOManager) ToLanguageDTOs(languages []domains.LanguageR) []LanguageDTO {
	var languageDTOs []LanguageDTO
	for _, lang := range languages {
		languageDTOs = append(languageDTOs, m.ToLanguageDTO(&lang))
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
		ID:         path.ID,
		Language:   languages,
		Difficulty: path.Quest.Difficulty,
		IsFinished: path.IsFinished,
		IsStarted:  path.IsStarted,
	}
}

type RoadDTO struct {
	Name     string    `json:"name"`
	IconPath string    `json:"iconPath"`
	Paths    []PathDTO `json:"paths"`
}

func (m *RoadDTOManager) ToRoadDTO(road domains.Roads, paths []PathDTO) RoadDTO {
	return RoadDTO{
		Name:     road.Name,
		IconPath: road.IconPath,
		Paths:    paths,
	}
}
