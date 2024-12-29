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
	ID              int            `json:"id"`
	ProgrammingName string         `json:"programmingName"`
	Languages       LabLanguageDTO `json:"language"`
	Template        string         `json:"template,omitempty"`
	IsStarted       bool           `json:"isStarted"`
	IsFinished      bool           `json:"isFinished"`
	Difficulty      int            `json:"difficulty"`
	FileExtention   string         `json:"fileExtention"`
	MonacoEditor    string         `json:"monacoEditor"`
}

type LabForAllDTO struct {
	ID         int                  `json:"id"`
	Languages  LabLanguageForAllDTO `json:"language"`
	IsFinished bool                 `json:"isFinished"`
	Difficulty int                  `json:"difficulty"`
}

type LabsDTO struct {
	Labs          []LabDTO `json:"labs"`
	IsImageExists bool     `json:"isImageExists"`
}

type LabsForAllDTO struct {
	Labs          []LabForAllDTO `json:"labs"`
	IsImageExists bool           `json:"isImageExists"`
	IconPath      string         `json:"iconPath"`
}

func (m *LabDTOManager) ToLabsDTO(labs []LabDTO, isImageExists bool) LabsDTO {
	return LabsDTO{
		Labs:          labs,
		IsImageExists: isImageExists,
	}
}

func (m *LabDTOManager) ToLabsForAllDTO(labs []LabForAllDTO, isImageExists bool, iconPath string) LabsForAllDTO {
	return LabsForAllDTO{
		Labs:          labs,
		IsImageExists: isImageExists,
		IconPath:      iconPath,
	}
}

// This is for get lab by id
func (m *LabDTOManager) ToLabDTO(lab domains.Lab, languagesDTO LabLanguageDTO, template string, programmingName, fileExtention, monacoEditor string) LabDTO {
	return LabDTO{
		ID:              lab.GetID(),
		ProgrammingName: programmingName,
		Languages:       languagesDTO,
		Template:        template,
		IsStarted:       lab.GetIsStarted(),
		IsFinished:      lab.GetIsFinished(),
		Difficulty:      lab.GetQuest().GetDifficulty(),
		FileExtention:   fileExtention,
		MonacoEditor:    monacoEditor,
	}
}

func (m *LabDTOManager) ToLabForAllDTO(lab domains.Lab, languagesDTO LabLanguageForAllDTO) LabForAllDTO {
	return LabForAllDTO{
		ID:         lab.GetID(),
		Languages:  languagesDTO,
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

type LabLanguageForAllDTO struct {
	Lang  string `json:"lang"`
	Title string `json:"title"`
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

func (m *LabDTOManager) ToLanguageForAllDTO(languageLabs []domains.LanguageLab, language string) LabLanguageForAllDTO {
	var newLanguage LabLanguageForAllDTO

	for _, languageLab := range languageLabs {
		if languageLab.GetLang() == language {
			newLanguage = LabLanguageForAllDTO{
				Lang:  languageLab.GetLang(),
				Title: languageLab.GetTitle(),
			}
		}
	}

	return newLanguage
}

func (m *LabDTOManager) FilterLabDTOs(labDTOs []LabDTO) []LabDTO {
	idMap := make(map[int]bool)
	var newLabDTOs []LabDTO

	for _, labDTO := range labDTOs {
		if !idMap[labDTO.ID] {
			newLabDTOs = append(newLabDTOs, labDTO)
			idMap[labDTO.ID] = true
		}
	}

	return newLabDTOs
}

func (m *LabDTOManager) FilterLabForAllDTOs(labDTOs []LabForAllDTO) []LabForAllDTO {
	idMap := make(map[int]bool)
	var newLabDTOs []LabForAllDTO

	for _, labDTO := range labDTOs {
		if !idMap[labDTO.ID] {
			newLabDTOs = append(newLabDTOs, labDTO)
			idMap[labDTO.ID] = true
		}
	}

	return newLabDTOs
}

type UserProgrammingLanguageLabStatsDTO struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	IconPath      string  `json:"iconPath"`
	TotalLabs     int     `json:"totalLabs"`
	CompletedLabs int     `json:"completedLabs"`
	Percentage    float32 `json:"percentage"`
}

func (m *LabDTOManager) ToUserProgrammingLanguageStatDTO(stat *domains.ProgrammingLanguageStats) UserProgrammingLanguageLabStatsDTO {
	return UserProgrammingLanguageLabStatsDTO{
		ID:            stat.GetID(),
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
