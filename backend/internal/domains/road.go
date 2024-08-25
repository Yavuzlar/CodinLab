package domains

// IRoadService is the interface that provides the methods for the road service.
type IRoadService interface {
	GetRoadFilter(userID string, programmingID, pathId int, isStarted, isFinished *bool) ([]Roads, error)
}

// LanguageRoad represents the details of a programming language.
type LanguageRoad struct {
	lang        string
	title       string
	description string
	content     string
	note        string
}

// Getter and Setter methods for LanguageRoad
func (l *LanguageRoad) GetLang() string {
	return l.lang
}

func (l *LanguageRoad) SetLang(lang string) {
	l.lang = lang
}

func (l *LanguageRoad) GetTitle() string {
	return l.title
}

func (l *LanguageRoad) SetTitle(title string) {
	l.title = title
}

func (l *LanguageRoad) GetDescription() string {
	return l.description
}

func (l *LanguageRoad) SetDescription(description string) {
	l.description = description
}

func (l *LanguageRoad) GetContent() string {
	return l.content
}

func (l *LanguageRoad) SetContent(content string) {
	l.content = content
}

func (l *LanguageRoad) GetNote() string {
	return l.note
}

func (l *LanguageRoad) SetNote(note string) {
	l.note = note
}

// NewLanguageRoad creates a new instance of LanguageRoad
func NewLanguageRoad(lang, title, description, content, note string) *LanguageRoad {
	return &LanguageRoad{
		lang:        lang,
		title:       title,
		description: description,
		content:     content,
		note:        note,
	}
}

// TestRoad represents a test case for a function.
type TestRoad struct {
	input  []string
	output []string
}

// Getter and Setter methods for TestRoad
func (t *TestRoad) GetInput() []string {
	return t.input
}

func (t *TestRoad) SetInput(input []string) {
	t.input = input
}

func (t *TestRoad) GetOutput() []string {
	return t.output
}

func (t *TestRoad) SetOutput(output []string) {
	t.output = output
}

// NewTestRoad creates a new instance of TestRoad
func NewTestRoad(input, output []string) *TestRoad {
	return &TestRoad{
		input:  input,
		output: output,
	}
}

// ParamRoad represents a parameter of a function.
type ParamRoad struct {
	name string
	typ  string
}

// Getter and Setter methods for ParamRoad
func (p *ParamRoad) GetName() string {
	return p.name
}

func (p *ParamRoad) SetName(name string) {
	p.name = name
}

func (p *ParamRoad) GetType() string {
	return p.typ
}

func (p *ParamRoad) SetType(typ string) {
	p.typ = typ
}

// NewParamRoad creates a new instance of ParamRoad
func NewParamRoad(name, typ string) *ParamRoad {
	return &ParamRoad{
		name: name,
		typ:  typ,
	}
}

// QuestRoad represents a coding challenge or task.
type QuestRoad struct {
	difficulty int
	funcName   string
	tests      []TestRoad
	params     []ParamRoad
}

// Getter and Setter methods for QuestRoad
func (q *QuestRoad) GetDifficulty() int {
	return q.difficulty
}

func (q *QuestRoad) SetDifficulty(difficulty int) {
	q.difficulty = difficulty
}

func (q *QuestRoad) GetFuncName() string {
	return q.funcName
}

func (q *QuestRoad) SetFuncName(funcName string) {
	q.funcName = funcName
}

func (q *QuestRoad) GetTests() []TestRoad {
	return q.tests
}

func (q *QuestRoad) SetTests(tests []TestRoad) {
	q.tests = tests
}

func (q *QuestRoad) GetParams() []ParamRoad {
	return q.params
}

func (q *QuestRoad) SetParams(params []ParamRoad) {
	q.params = params
}

// NewQuestRoad creates a new instance of QuestRoad
func NewQuestRoad(difficulty int, funcName string, tests []TestRoad, params []ParamRoad) *QuestRoad {
	return &QuestRoad{
		difficulty: difficulty,
		funcName:   funcName,
		tests:      tests,
		params:     params,
	}
}

// Path represents a specific coding road exercise.
type Path struct {
	id         int
	languages  []LanguageRoad
	quest      QuestRoad
	isStarted  bool
	isFinished bool
}

// Getter and Setter methods for Path
func (p *Path) GetID() int {
	return p.id
}

func (p *Path) SetID(id int) {
	p.id = id
}

func (p *Path) GetLanguages() []LanguageRoad {
	return p.languages
}

func (p *Path) SetLanguages(languages []LanguageRoad) {
	p.languages = languages
}

func (p *Path) GetQuest() *QuestRoad {
	return &p.quest
}

func (p *Path) SetQuest(quest QuestRoad) {
	p.quest = quest
}

func (p *Path) GetIsStarted() bool {
	return p.isStarted
}

func (p *Path) SetIsStarted(isStarted bool) {
	p.isStarted = isStarted
}

func (p *Path) GetIsFinished() bool {
	return p.isFinished
}

func (p *Path) SetIsFinished(isFinished bool) {
	p.isFinished = isFinished
}

// NewPath creates a new instance of Path
func NewPath(id int, languages []LanguageRoad, quest QuestRoad, isStarted, isFinished bool) *Path {
	return &Path{
		id:         id,
		languages:  languages,
		quest:      quest,
		isStarted:  isStarted,
		isFinished: isFinished,
	}
}

// Roads represents a collection of roads grouped together.
type Roads struct {
	id          int
	name        string
	dockerImage string
	iconPath    string
	paths       []Path
}

// Getter and Setter methods for Roads
func (r *Roads) GetID() int {
	return r.id
}

func (r *Roads) SetID(id int) {
	r.id = id
}

func (r *Roads) GetName() string {
	return r.name
}

func (r *Roads) SetName(name string) {
	r.name = name
}

func (r *Roads) GetDockerImage() string {
	return r.dockerImage
}

func (r *Roads) SetDockerImage(dockerImage string) {
	r.dockerImage = dockerImage
}

func (r *Roads) GetIconPath() string {
	return r.iconPath
}

func (r *Roads) SetIconPath(iconPath string) {
	r.iconPath = iconPath
}

func (r *Roads) GetPaths() []Path {
	return r.paths
}

func (r *Roads) SetPaths(paths []Path) {
	r.paths = paths
}

// NewRoads creates a new instance of Roads
func NewRoads(id int, name, dockerImage, iconPath string, paths []Path) *Roads {
	return &Roads{
		id:          id,
		name:        name,
		dockerImage: dockerImage,
		iconPath:    iconPath,
		paths:       paths,
	}
}
