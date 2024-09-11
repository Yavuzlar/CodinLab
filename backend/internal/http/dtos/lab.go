package dto

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
)

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

type UserLabDifficultyStatsDTO struct {
	EasyPercentage   float32 `json:"easyPercentage"`
	MediumPercentage float32 `json:"mediumPercentage"`
	HardPercentage   float32 `json:"hardPercentage"`
}

func (m *LabDTOManager) ToUserLabDifficultyStatsDto(stats domains.UserLabDifficultyStats) UserLabDifficultyStatsDTO {
	return UserLabDifficultyStatsDTO{
		EasyPercentage:   stats.GetEasyPercentage(),
		MediumPercentage: stats.GetMediumPercentage(),
		HardPercentage:   stats.GetHardPercentage(),
	}
}

type UserLabProgressStatsDto struct {
	Progress  float32 `json:"progress"`
	Completed float32 `json:"completed"`
}

func (m *LabDTOManager) ToUserLabProgressStatsDto(stats domains.UserLabProgressStats) UserLabProgressStatsDto {
	return UserLabProgressStatsDto{
		Completed: stats.GetCompleted(),
		Progress:  stats.GetProgress(),
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
	Percentage    float32 `json:"percentage"`
}

func (m *LabDTOManager) ToUserProgrammingLanguageStatDTO(stat *domains.ProgrammingLanguageStats) UserProgrammingLanguageLabStatsDTO {
	return UserProgrammingLanguageLabStatsDTO{
		TotalLabs:     stat.GetTotalLabs(),
		CompletedLabs: stat.GetCompletedLabs(),
		Percentage:    stat.GetPercentage(),
	}
}

func (m *LabDTOManager) ToUserProgrammingLanguageStatsDTO(stats []domains.ProgrammingLanguageStats) []UserProgrammingLanguageLabStatsDTO {
	var userProgrammingLanguageLabStatsDTO []UserProgrammingLanguageLabStatsDTO
	for _, stat := range stats {
		userProgrammingLanguageLabStatsDTO = append(userProgrammingLanguageLabStatsDTO, m.ToUserProgrammingLanguageStatDTO(&stat))
	}
	return userProgrammingLanguageLabStatsDTO
}

type AnswerLabDTO struct {
	ProgrammingID int    `json:"programmindID"`
	LabID         int    `json:"labID"`
	UserCode      string `json:"userCode"`
}
