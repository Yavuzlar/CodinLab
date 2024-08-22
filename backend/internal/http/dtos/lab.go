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
		ID:         lab.GetID(),
		Languages:  languagesDTOs,
		IsStarted:  lab.GetIsStarted(),
		IsFinished: lab.GetIsFinished(),
		Difficulty: lab.GetQuest().GetDifficulty(),
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
		ID:       labs.GetID(),
		Name:     labs.GetName(),
		IconPath: labs.GetIconPath(),
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
		TotalLabs:        stats.GetTotalLabs(),
		TotalPercentage:  stats.GetTotalPercentage(),
		EasyLabs:         stats.GetEasyLabs(),
		EasyPercentage:   stats.GetEasyPercentage(),
		MediumLabs:       stats.GetMediumLabs(),
		MediumPercentage: stats.GetMediumPercentage(),
		HardLabs:         stats.GetHardLabs(),
		HardPercentage:   stats.GetHardPercentage(),
	}
}

type LabLanguageDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Hint        string `json:"hint"`
}

func (m *LabDTOManager) ToLanguageDTO(language domains.LanguageLab) LabLanguageDTO {
	return LabLanguageDTO{
		Lang:        language.GetLang(),
		Title:       language.GetTitle(),
		Description: language.GetDescription(),
		Hint:        language.GetHint(),
		Note:        language.GetNote(),
	}
}

func (m *LabDTOManager) ToLanguageDTOs(languages []domains.LanguageLab) []LabLanguageDTO {
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
		TotalLabs:     stats.GetTotalLabs(),
		CompletedLabs: stats.GetCompletedLabs(),
		Percentage:    stats.GetPercentage(),
	}
}
