package domains

// IParserService is the interface that provides the methods for the parser service.
type IParserService interface {
	GetLabs() (labs []Labs, err error)
	GetRoads() (roads []Road, err error)
}

// Inventory represents the information related to an item in inventory.
type Inventory struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	DockerImage string     `json:"dockerImage"`
	PathDir     string     `json:"pathDir"`
	LabDir      string     `json:"labDir"`
	IconPath    string     `json:"iconPath"`
	Languages   []Language `json:"languages"`
}

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
	ID        int        `json:"id"`
	Languages []Language `json:"languages"`
	Quest     Quest      `json:"quest"`
}

// Labs represents a collection of labs grouped together.
type Labs struct {
	Name        string
	DockerImage string
	IconPath    string
	Labs        []Lab
}

// Path represents a coding learning path.
type Path struct {
	ID        int        `json:"id"`
	Languages []Language `json:"languages"`
	Quest     Quest      `json:"quest"`
}

// Road represents a collection of learning paths.
type Road struct {
	Name        string
	DockerImage string
	IconPath    string
	Paths       []Path
}
