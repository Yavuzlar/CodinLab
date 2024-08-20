package domains

import "context"

// IHomeService is the interface that provides the methods for the home page service.
type IHomeService interface {
	GetUserLevel(ctx context.Context, userID string) (*UserLevel, error)
	GetUserDevelopment(ctx context.Context, userID string) (Development, error)
	GetUserAdvancement(ctx context.Context, userID string) (advancement []Advancement, err error)
	GetInventory(ctx context.Context) (inventory []Inventory, err error)
	GetWelcomeContent() (content []WelcomeContent, err error)
	GetRoadContent() (content []RoadContent, err error)
	GetLabContent() (content []LabContent, err error)
}

func NewUserDevelopment(roadPercentage int32, labPercentage int32) *Development {
	userDevelopment := &Development{
		roadPercentage: roadPercentage,
		labPercentage:  labPercentage,
	}
	return userDevelopment
}

func NewInventory(id int, name string, iconPath string) *Inventory {
	inventory := &Inventory{
		ID:       id,
		Name:     name,
		IconPath: iconPath,
	}
	return inventory
}

func NewAdvancement(id int, name string, iconPath string, roadPercentage int32, labPercentage int32) *Advancement {
	advancement := &Advancement{
		programmingID:  id,
		name:           name,
		iconPath:       iconPath,
		roadPercentage: roadPercentage,
		labPercentage:  labPercentage,
	}
	return advancement
}

type Inventory struct {
	ID       int
	Name     string
	IconPath string
}

type Development struct {
	roadPercentage int32
	labPercentage  int32
}

func (d *Development) RoadPercentage() int32 {
	return d.roadPercentage
}

func (d *Development) LabPercentage() int32 {
	return d.labPercentage
}

type Advancement struct {
	programmingID  int
	name           string
	iconPath       string
	roadPercentage int32
	labPercentage  int32
}

func (a *Advancement) ProgrammingID() int {
	return a.programmingID
}

func (a *Advancement) Name() string {
	return a.name
}

func (a *Advancement) RoadPercentage() int32 {
	return a.roadPercentage
}

func (a *Advancement) LabPercentage() int32 {
	return a.labPercentage
}

func (a *Advancement) IconPath() string {
	return a.iconPath
}
