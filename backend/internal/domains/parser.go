package domains

// IParserService is the interface that provides the methods for the parser service.
type IParserService interface {
	GetLabs() (labs []LabsP, err error)
	GetRoads() (roads []RoadP, err error)
	GetInventory() (inventory []InventoryP, err error)
	GetLevels() (userLevel []LevelP, err error)
}

// Inventory represents the information related to an item in inventory.
type InventoryP struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	DockerImage string      `json:"dockerImage"`
	PathDir     string      `json:"pathDir"`
	LabDir      string      `json:"labDir"`
	IconPath    string      `json:"iconPath"`
	Languages   []LanguageP `json:"languages"`
}

// Language represents the details of a programming language.
type LanguageP struct {
	Lang        string `json:"lang"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Note        string `json:"note,omitempty"`
	Hint        string `json:"hint,omitempty"`
}

// Test represents a test case for a function.
type TestP struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

// Param represents a parameter of a function.
type ParamP struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Quest represents a coding challenge or task.
type QuestP struct {
	Difficulty int      `json:"difficulty"`
	FuncName   string   `json:"funcName"`
	Tests      []TestP  `json:"tests"`
	Params     []ParamP `json:"params"`
}

// Lab represents a specific coding lab exercise.
type LabP struct {
	ID        int         `json:"id"`
	Languages []LanguageP `json:"languages"`
	Quest     QuestP      `json:"quest"`
}

// Labs represents a collection of labs grouped together.
type LabsP struct {
	ID          int
	Name        string
	DockerImage string
	IconPath    string
	Labs        []LabP
}

// Path represents a coding learning path.
type PathP struct {
	ID        int         `json:"id"`
	Languages []LanguageP `json:"languages"`
	Quest     QuestP      `json:"quest"`
}

// Road represents a collection of learning paths.
type RoadP struct {
	Name        string
	DockerImage string
	IconPath    string
	Paths       []PathP
}

type LevelP struct {
	Level     int         `json:"level"`
	MinPoints int32       `json:"minPoints"`
	MaxPoints int32       `json:"maxPoints"`
	Languages []LanguageP `json:"languages"`
}
