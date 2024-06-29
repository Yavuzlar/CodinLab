
package domains

// IRoadService is the interface that provides the methods for the road service.
type IRoadService interface {
	GetRoadFilter(userID string, roadId, pathId int, isStarted, isFinished bool) ([]Roads, error)
}

// Language represents the details of a programming language.
type LanguageR struct {
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
	Note        string `json:"note,omitempty"`
}

// Test represents a test case for a function.
type TestR struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

// Param represents a parameter of a function.
type ParamR struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Quest represents a coding challenge or task.
type QuestR struct {
	Difficulty int     `json:"difficulty"`
	FuncName   string  `json:"funcName"`
	Tests      []TestR  `json:"tests"`
	Params     []ParamR `json:"params"`
}

// Path represents a specific coding road exercise.
type Path struct {
	ID         int         `json:"id"`
	Languages  []LanguageR `json:"languages"`
	Quest      QuestR       `json:"quest"`
	IsStarted  bool      `json:"isStarted"`
	IsFinished bool      `json:"isFinished"`
}

// Roads represents a collection of roads grouped together.
type Roads struct {
	ID          int
	Name        string `json:"name"`
	DockerImage string `json:"dockerImage"`
	IconPath    string `json:"iconPath"`
	Paths       []Path `json:"labs"`
}

