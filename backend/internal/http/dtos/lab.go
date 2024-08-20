package dto

import "github.com/Yavuzlar/CodinLab/internal/domains"

type LabDTOManager struct{}

// NewLabDTOManager creates a new instance of LabDTOManager
func NewLabDTOManager() LabDTOManager {
	return LabDTOManager{}
}

type LabDTO struct {
	ID         int              `json:"id"`
	Languages  []LabLanguageDTO `json:"languages"`
	IsStarted  bool             `json:"isStarted"`
	IsFinished bool             `json:"isFinished"`
	Difficulty int              `json:"difficulty"`
}

func (m *LabDTOManager) ToLabDTO(lab domains.Lab, languagesDTOs []LabLanguageDTO) LabDTO {
	return LabDTO{
		ID:         lab.ID,
		Languages:  languagesDTOs,
		IsStarted:  lab.IsStarted,
		IsFinished: lab.IsFinished,
		Difficulty: lab.Quest.Difficulty,
	}
}

type LabsDTO struct {
	ID       int
	Name     string   `json:"name"`
	IconPath string   `json:"iconPath"`
	Labs     []LabDTO `json:"labs"`
}

func (m *LabDTOManager) ToLabsDTO(labs domains.Labs, labDTOs []LabDTO) LabsDTO {
	return LabsDTO{
		ID:       labs.ID,
		Name:     labs.Name,
		IconPath: labs.IconPath,
		Labs:     labDTOs,
	}
}

type UserGeneralLabStatsDTO struct {
	TotalLabs        int     `json:"totalLabs"`
	TotalPercentage  float64 `json:"TotalPercentage"`
	EasyLabs         int     `json:"easyLabs"`
	EasyPercentage   float64 `json:"easyPercentage"`
	MediumLabs       int     `json:"mediumlabs"`
	MediumPercentage float64 `json:"mediumPercentage"`
	HardLabs         int     `json:"hardLabs"`
	HardPercentage   float64 `json:"hardPercentage"`
}

func (m *LabDTOManager) ToUserGeneralLabStatsDTO(stats domains.GeneralStats) UserGeneralLabStatsDTO {
	return UserGeneralLabStatsDTO{
		TotalLabs:        stats.TotalLabs,
		TotalPercentage:  stats.TotalPercentage,
		EasyLabs:         stats.EasyLabs,
		EasyPercentage:   stats.EasyPercentage,
		MediumLabs:       stats.MediumLabs,
		MediumPercentage: stats.MediumPercentage,
		HardLabs:         stats.HardLabs,
		HardPercentage:   stats.HardPercentage,
	}
}

type LabLanguageDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Hint        string `json:"hint"`
}

func (m *LabDTOManager) ToLanguageDTO(language domains.Language) LabLanguageDTO {
	return LabLanguageDTO{
		Lang:        language.Lang,
		Title:       language.Title,
		Description: language.Description,
		Hint:        language.Hint,
		Note:        language.Note,
	}
}

func (m *LabDTOManager) ToLanguageDTOs(languages []domains.Language) []LabLanguageDTO {
	var languageDTOs []LabLanguageDTO
	for _, lang := range languages {
		languageDTOs = append(languageDTOs, m.ToLanguageDTO(lang))
	}
	return languageDTOs
}

type UserProgrammingLanguageLabStatsDTO struct {
	TotalLabs     int     `json:"totalLabs"`
	CompletedLabs int     `json:"completedLabs"`
	Percentage    float64 `json:"percentage"`
}

func (m *LabDTOManager) ToUserProgrammingLanguageLabStatsDTO(stats domains.ProgrammingLanguageStats) UserProgrammingLanguageLabStatsDTO {
	return UserProgrammingLanguageLabStatsDTO{
		TotalLabs:     stats.TotalLabs,
		CompletedLabs: stats.CompletedLabs,
		Percentage:    stats.Percentage,
	}
}
