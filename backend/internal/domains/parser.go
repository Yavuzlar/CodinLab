package domains

// IParserService is the interface that provides the methods for the parser service.
type IParserService interface {
	GetLabs() (labs []LabP, err error)
	GetRoads() (roads []RoadP, err error)
	GetInventory() (inventory []InventoryP, err error)
	GetLevels() (userLevel []LevelP, err error)
	GetWelcomeBanner() (content []WelcomeContent, err error)
	GetLabBanner() (content []LabContent, err error)
	GetRoadBanner() (content []RoadContent, err error)
}

// Inventory represents the information related to an item in inventory.
type InventoryP struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	DockerImage   string      `json:"dockerImage"`
	PathDir       string      `json:"pathDir"`
	LabDir        string      `json:"labDir"`
	IconPath      string      `json:"iconPath"`
	Cmd           []string    `json:"cmd"`
	FileExtension string      `json:"fileExtension"`
	Languages     []LanguageP `json:"languages"`
}

// Language represents the details of a programming language.
type LanguageP struct {
	Lang        string `json:"lang"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	Note        string `json:"note,omitempty"`
	Hint        string `json:"hint,omitempty"`
	Content     string `json:"content,omitempty"`
}

// Test represents a test case for a function.
type TestP struct {
	Input  []interface{} `json:"input"`
	Output []interface{} `json:"output"`
}

// Param represents a parameter of a function.
type ParamP struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Returns represents a parameter of a function.
type ReturnP struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type CodeTemplateP struct {
	ProgrammingID int    `json:"programmingID"`
	TemplatePath  string `json:"templatePath"`
}

type QuestP struct {
	Difficulty    int             `json:"difficulty"`
	FuncName      string          `json:"funcName"`
	Tests         []TestP         `json:"tests"`
	Params        []ParamP        `json:"params"`
	Returns       []ReturnP       `json:"returns"`
	QuestImports  []string        `json:"questImports,omitempty"`
	CodeTemplates []CodeTemplateP `json:"codeTemplates"`
}

// Lab represents a specific coding lab exercise.
type LabP struct {
	ID        int         `json:"id"`
	Languages []LanguageP `json:"languages"`
	Quest     QuestP      `json:"quest"`
}

// Path represents a coding learning path.
type PathP struct {
	ID        int         `json:"id"`
	Languages []LanguageP `json:"languages"`
	Quest     QuestP      `json:"quest"`
}

// Road represents a collection of learning paths.
type RoadP struct {
	ID            int
	Name          string
	DockerImage   string
	IconPath      string
	Cmd           []string
	FileExtension string
	Paths         []PathP
}

type LevelP struct {
	Level     int         `json:"level"`
	MinPoints int32       `json:"minPoints"`
	MaxPoints int32       `json:"maxPoints"`
	Languages []LanguageP `json:"languages"`
}

// The structure of the Welcome content on the home page
type WelcomeContent struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// The structure of the Road content on the home page
type RoadContent struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// The structure of the Lab content on the home page
type LabContent struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
