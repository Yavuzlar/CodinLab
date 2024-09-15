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
	ID         int            `json:"id"`
	Languages  LabLanguageDTO `json:"language"`
	Template   string         `json:"template,omitempty"`
	IsStarted  bool           `json:"isStarted"`
	IsFinished bool           `json:"isFinished"`
	Difficulty int            `json:"difficulty"`
}
type LabsDTO struct {
	ID         int              `json:"id"`
	Languages  []LabLanguageDTO `json:"languages"`
	IsStarted  bool             `json:"isStarted"`
	IsFinished bool             `json:"isFinished"`
	Difficulty int              `json:"difficulty"`
}

func (m *LabDTOManager) ToLabDTO(lab domains.Lab, languagesDTO LabLanguageDTO, template string) LabDTO {
	return LabDTO{
		ID:         lab.GetID(),
		Languages:  languagesDTO,
		Template:   template,
		IsStarted:  lab.GetIsStarted(),
		IsFinished: lab.GetIsFinished(),
		Difficulty: lab.GetQuest().GetDifficulty(),
	}
}
func (m *LabDTOManager) ToLabsDTO(lab domains.Lab, languagesDTOs LabLanguageDTO) LabDTO {
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
	Progress  int `json:"progress"`
	Completed int `json:"completed"`
}

func (m *LabDTOManager) ToUserLabProgressStatsDto(stats domains.UserLabProgressStats) UserLabProgressStatsDto {
	return UserLabProgressStatsDto{
		Completed: int(stats.GetCompleted()),
		Progress:  int(stats.GetProgress()),
	}
}

type LabLanguageDTO struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Hint        string `json:"hint"`
}

func (m *LabDTOManager) ToLanguageDTO(languageLabs []domains.LanguageLab, language string) LabLanguageDTO {
	var newLanguage LabLanguageDTO

	for _, languageLab := range languageLabs {
		if languageLab.GetLang() == language {
			newLanguage = LabLanguageDTO{
				Lang:        languageLab.GetLang(),
				Title:       languageLab.GetTitle(),
				Description: languageLab.GetDescription(),
				Hint:        languageLab.GetHint(),
				Note:        languageLab.GetNote(),
			}
		}
	}

	return newLanguage
}

type UserProgrammingLanguageLabStatsDTO struct {
	Name          string  `json:"name"`
	IconPath      string  `json:"iconPath"`
	TotalLabs     int     `json:"totalLabs"`
	CompletedLabs int     `json:"completedLabs"`
	Percentage    float32 `json:"percentage"`
}

func (m *LabDTOManager) ToUserProgrammingLanguageStatDTO(stat *domains.ProgrammingLanguageStats) UserProgrammingLanguageLabStatsDTO {
	return UserProgrammingLanguageLabStatsDTO{
		Name:          stat.GetName(),
		IconPath:      stat.GetIconPath(),
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
	UserCode string `json:"userCode"`
}

func (m *LabDTOManager) ToFrontendTemplateDto(frontendTemplate string) LabFrontendTemplateDto {
	return LabFrontendTemplateDto{
		Template: frontendTemplate,
	}
}

type LabFrontendTemplateDto struct {
	Template string `json:"template"`
}
