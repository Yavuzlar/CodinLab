package domains

// ILabService is the interface that provides the methods for the lab service.
type ILabService interface {
	GetLabsFilter(userID string, labsId, labId int, isStarted, isFinished bool) ([]Labs, error)
	UserLanguageLabStats(userID string, language string) (LanguageStats, error)
	UserGeneralLabStats(userID string) (GeneralStats, error)
}

// LanguageStats represents the statistics for a specific language lab.
type LanguageStats struct {
	TotalLabs     int     `json:"totalLabs"`
	CompletedLabs int     `json:"completedLabs"`
	Percentage    float64 `json:"Percentage"`
}

// UserGeneralLabStatsDto represents the DTO for user general lab statistics
type GeneralStats struct {
	TotalLabs        int     `json:"totalLabs"`
	TotalPercentage  float64 `json:"TotalPercentage"`
	EasyLabs         int     `json:"easyLabs"`
	EasyPercentage   float64 `json:"easyPercentage"`
	MediumLabs       int     `json:"mediumlabs"`
	MediumPercentage float64 `json:"mediumPercentage"`
	HardLabs         int     `json:"hardLabs"`
	HardPercentage   float64 `json:"hardPercentage"`
}

// ILabService is the interface that provides the methods for the lab service.

// Language represents the details of a programming language.
type Language struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Note        string `json:"note,omitempty"`
	Hint        string `json:"hint,omitempty"`
}

// Test represents a test case for a function.
type Test struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

// Param represents a parameter of a function.
type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Quest represents a coding challenge or task.
type Quest struct {
	Difficulty int     `json:"difficulty"`
	FuncName   string  `json:"funcName"`
	Tests      []Test  `json:"tests"`
	Params     []Param `json:"params"`
}

// Lab represents a specific coding lab exercise.
type Lab struct {
	ID         int        `json:"id"`
	Languages  []Language `json:"languages"`
	Quest      Quest      `json:"quest"`
	IsStarted  bool       `json:"isStarted"`
	IsFinished bool       `json:"isFinished"`
}

// Labs represents a collection of labs grouped together.
type Labs struct {
	ID          int
	Name        string `json:"name"`
	DockerImage string `json:"dockerImage"`
	IconPath    string `json:"iconPath"`
	Labs        []Lab  `json:"labs"`
}
