package dto

import (
	"time"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/google/uuid"
)

type LogDTOManager struct{}

func NewLogDTOManager() LogDTOManager {
	return LogDTOManager{}
}

type LogDTO struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"userId"`
	ProgrammingID int32     `json:"programmingID"`
	LabPathID     int32     `json:"labPathID"`
	LogType       string    `json:"logType"`
	Content       string    `json:"content"`
}

// Lab & Road numbers solved day by day
type SolutionsByDayDTO struct {
	Date      time.Time `json:"date"`
	RoadCount int       `json:"roadCount"`
	LabCount  int       `json:"labCount"`
}

// Represents the total hours spent on Lab & Road solutions for each programming language.
type SolutionsHoursByProgrammingDTO struct {
	ProgrammingID int32   `json:"programmingID"`
	LabHours      float64 `json:"labHours"`
	RoadHours     float64 `json:"roadHours"`
}

type LanguageUsageRatesDTO struct {
	IconPath        string  `json:"iconPath"`
	Name            string  `json:"name"`
	UsagePercentage float32 `json:"usagePercentage"`
}

func (m *LogDTOManager) ToLanguageUsageRatesDTO(log domains.LanguageUsageRates) LanguageUsageRatesDTO {
	return LanguageUsageRatesDTO{
		IconPath:        log.GetIconPath(),
		Name:            log.GetName(),
		UsagePercentage: log.GetUsagePercentage(),
	}
}

func (m *LogDTOManager) ToLanguageUsageRatesDTOs(logs []domains.LanguageUsageRates) (languageUsageRatesDTOs []LanguageUsageRatesDTO) {
	for _, log := range logs {
		languageUsageRatesDTOs = append(languageUsageRatesDTOs, m.ToLanguageUsageRatesDTO(log))
	}
	return
}

func (m *LogDTOManager) ToLogDTO(log domains.Log) LogDTO {
	return LogDTO{
		ID:            log.ID(),
		UserID:        log.UserID(),
		ProgrammingID: log.ProgrammingID(),
		LabPathID:     log.LabPathID(),
		LogType:       log.Type(),
		Content:       log.Content(),
	}
}

func (m *LogDTOManager) ToLogDTOs(logs []domains.Log) []LogDTO {
	var logDTOs []LogDTO
	for _, log := range logs {
		logDTOs = append(logDTOs, m.ToLogDTO(log))
	}
	return logDTOs
}

func (m *LogDTOManager) ToSolutionsByDayDTO(solutionsByDay domains.SolutionsByDay) SolutionsByDayDTO {
	return SolutionsByDayDTO{
		Date:      solutionsByDay.GetDate(),
		RoadCount: solutionsByDay.GetRoadCount(),
		LabCount:  solutionsByDay.GetLabCount(),
	}
}

func (m *LogDTOManager) ToSolutionsByDayDTOs(solutionsByDays []domains.SolutionsByDay) []SolutionsByDayDTO {
	var solutionByDayDTOs []SolutionsByDayDTO
	for _, solutionByDays := range solutionsByDays {
		solutionByDayDTOs = append(solutionByDayDTOs, m.ToSolutionsByDayDTO(solutionByDays))
	}
	return solutionByDayDTOs
}

func (m *LogDTOManager) ToSolutionsHoursByProgrammingDTO(domain domains.SolutionsHoursByProgramming) SolutionsHoursByProgrammingDTO {
	return SolutionsHoursByProgrammingDTO{
		ProgrammingID: domain.GetProgrammingID(),
		RoadHours:     domain.GetRoadHours(),
		LabHours:      domain.GetLabHours(),
	}
}

func (m *LogDTOManager) ToSolutionsHoursByProgrammingDTOs(domains []domains.SolutionsHoursByProgramming) []SolutionsHoursByProgrammingDTO {
	var solutionByDayDTOs []SolutionsHoursByProgrammingDTO
	for _, solutionByDays := range domains {
		solutionByDayDTOs = append(solutionByDayDTOs, m.ToSolutionsHoursByProgrammingDTO(solutionByDays))
	}
	return solutionByDayDTOs
}
