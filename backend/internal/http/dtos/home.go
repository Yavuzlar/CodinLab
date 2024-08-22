package dto

import "github.com/Yavuzlar/CodinLab/internal/domains"

type HomeDTOManager struct{}

// NewHomeDTOManager creates a new instance of HomeDTOManager
func NewHomeDTOManager() HomeDTOManager {
	return HomeDTOManager{}
}

type UserLevelDTO struct {
	Level           int                     `json:"level"`
	TotalPoints     int32                   `json:"totalPoints"`
	LevelPercentage int32                   `json:"levelPercentage"`
	Languages       []domains.LanguageLevel `json:"languages"`
}

func (m *HomeDTOManager) ToUserLevelDTO(userLevel *domains.UserLevel) UserLevelDTO {
	return UserLevelDTO{
		Level:           userLevel.Level(),
		TotalPoints:     userLevel.TotalPoints(),
		LevelPercentage: userLevel.LevelPercentage(),
		Languages:       userLevel.Languages(),
	}
}

type InventoryDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IconPath string `json:"iconPath"`
}

func (m *HomeDTOManager) ToInventoryDTO(inventory *domains.Inventory) InventoryDTO {
	return InventoryDTO{
		ID:       inventory.ID(),
		Name:     inventory.Name(),
		IconPath: inventory.IconPath(),
	}
}

func (m *HomeDTOManager) ToInventoryDTOs(inventories []domains.Inventory) []InventoryDTO {
	var inventoryDTOs []InventoryDTO
	for _, inventory := range inventories {
		inventoryDTOs = append(inventoryDTOs, m.ToInventoryDTO(&inventory))
	}

	return inventoryDTOs
}

type DevelopmentDTO struct {
	RoadPercentage int32 `json:"roadPercentage"`
	LabPercentage  int32 `json:"labPercantage"`
}

func (m *HomeDTOManager) ToDevelopmentDTO(dev domains.Development) DevelopmentDTO {
	return DevelopmentDTO{
		RoadPercentage: dev.RoadPercentage(),
		LabPercentage:  dev.LabPercentage(),
	}
}

type AdvancementDTO struct {
	ProgrammingID  int    `json:"programmingID"`
	Name           string `json:"name"`
	IconPath       string `json:"iconPath"`
	RoadPercentage int32  `json:"roadPercentage"`
	LabPercentage  int32  `json:"labPercentage"`
}

func (m *HomeDTOManager) ToAdvancementDTO(advancement domains.Advancement) AdvancementDTO {
	return AdvancementDTO{
		ProgrammingID:  advancement.ProgrammingID(),
		Name:           advancement.Name(),
		IconPath:       advancement.IconPath(),
		RoadPercentage: advancement.RoadPercentage(),
		LabPercentage:  advancement.LabPercentage(),
	}
}

func (m *HomeDTOManager) ToAdvancementDTOs(advancements []domains.Advancement) []AdvancementDTO {
	var advancementDTOs []AdvancementDTO
	for _, advancement := range advancements {
		advancementDTOs = append(advancementDTOs, m.ToAdvancementDTO(advancement))
	}
	return advancementDTOs
}