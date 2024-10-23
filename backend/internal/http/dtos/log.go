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
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
	Level int       `json:"level"`
}

// Represents the total  spent on Lab & Road solutions for each programming language.
type SolutionsByProgrammingDTO struct {
	ProgrammingID   int32  `json:"programmingID"`
	ProgrammingName string `json:"programmingName"`
	LabCount        int    `json:"labCount"`
	RoadCount       int    `json:"roadCount"`
	TotalCount      int    `json:"totalCount"`
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
		Date:  solutionsByDay.GetDate(),
		Count: solutionsByDay.GetCount(),
		Level: solutionsByDay.GetLevel(),
	}
}

func (m *LogDTOManager) ToSolutionsByDayDTOs(solutionsByDays []domains.SolutionsByDay) []SolutionsByDayDTO {
	var solutionByDayDTOs []SolutionsByDayDTO
	for _, solutionByDays := range solutionsByDays {
		solutionByDayDTOs = append(solutionByDayDTOs, m.ToSolutionsByDayDTO(solutionByDays))
	}
	return solutionByDayDTOs
}

func (m *LogDTOManager) ToSolutionsByProgrammingDTO(domain domains.SolutionsByProgramming) SolutionsByProgrammingDTO {
	return SolutionsByProgrammingDTO{
		ProgrammingID:   domain.GetProgrammingID(),
		ProgrammingName: domain.GetProgrammingName(),
		RoadCount:       domain.GetRoadCount(),
		LabCount:        domain.GetLabCount(),
		TotalCount:      domain.GetTotalCount(),
	}
}

func (m *LogDTOManager) ToSolutionsByProgrammingDTOs(domains []domains.SolutionsByProgramming) []SolutionsByProgrammingDTO {
	var solutionByDayDTOs []SolutionsByProgrammingDTO
	for _, solutionByDays := range domains {
		solutionByDayDTOs = append(solutionByDayDTOs, m.ToSolutionsByProgrammingDTO(solutionByDays))
	}
	return solutionByDayDTOs
}
