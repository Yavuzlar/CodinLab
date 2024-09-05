package domains

// IRoadService is the interface that provides the methods for the road service.
type IRoadService interface {
	GetRoadFilter(userID string, programmingID, pathId int, isStarted, isFinished *bool) ([]Road, error)
	GetRoadInformation(programmingID int32) (*Road, error)
	GetUserLanguageRoadStats(userID string) ([]RoadStats, error)
	GetUserRoadProgressStats(userID string) (progressStats *RoadProgressStats, err error)
	CodeTemplateGenerator(programmingName, templatePathObject, content, funcName string, tests []TestRoad) (string, error)
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

// Road represents a collection of roads grouped together.
type Road struct {
	id            int
	name          string
	dockerImage   string
	iconPath      string
	cmd           string
	fileExtension string
	templatePath  string
	paths         []Path
	isStarted     bool
	isFinished    bool
}

// Getter and Setter methods for Roads
func (r *Road) GetID() int {
	return r.id
}

func (r *Road) SetID(id int) {
	r.id = id
}

func (r *Road) GetName() string {
	return r.name
}

func (r *Road) SetName(name string) {
	r.name = name
}

func (r *Road) GetDockerImage() string {
	return r.dockerImage
}

func (r *Road) SetDockerImage(dockerImage string) {
	r.dockerImage = dockerImage
}

func (r *Road) GetIconPath() string {
	return r.iconPath
}

func (r *Road) SetIconPath(iconPath string) {
	r.iconPath = iconPath
}

func (r *Road) GetCmd() string {
	return r.cmd
}

func (r *Road) SetCmd(cmd string) {
	r.cmd = cmd
}

func (r *Road) GetFileExtension() string {
	return r.fileExtension
}

func (r *Road) SetFileExtension(fileExtension string) {
	r.fileExtension = fileExtension
}

func (r *Road) GetTemplatePath() string {
	return r.templatePath
}

func (r *Road) SetTemplatePath(templatePath string) {
	r.templatePath = templatePath
}

func (r *Road) GetPaths() []Path {
	return r.paths
}

func (r *Road) SetPaths(paths []Path) {
	r.paths = paths
}

func (r *Road) GetIsStarted() bool {
	return r.isStarted
}

func (r *Road) SetIsStarted(ok bool) {
	r.isStarted = ok
}

func (r *Road) GetIsFinished() bool {
	return r.isFinished
}

func (r *Road) SetIsFinished(ok bool) {
	r.isFinished = ok
}

// NewRoads creates a new instance of Roads
func NewRoads(id int, name, dockerImage, iconPath, cmd, fileExtension, templatePath string, paths []Path, isStarted, isFinished bool) *Road {
	return &Road{
		id:            id,
		name:          name,
		dockerImage:   dockerImage,
		iconPath:      iconPath,
		cmd:           cmd,
		fileExtension: fileExtension,
		templatePath:  templatePath,
		paths:         paths,
		isStarted:     isStarted,
		isFinished:    isFinished,
	}
}

// RoadStats represents the statistics for a specific language lab.
type RoadStats struct {
	id             int
	name           string
	iconPath       string
	totalRoads     int
	completedRoads int
	percentage     float32
}

func (r *RoadStats) GetID() int {
	return r.id
}

func (r *RoadStats) GetName() string {
	return r.name
}

func (r *RoadStats) GetIconPath() string {
	return r.iconPath
}

func (r *RoadStats) GetTotalRoads() int {
	return r.totalRoads
}

func (r *RoadStats) GetCompletedRoads() int {
	return r.completedRoads
}

func (r *RoadStats) GetPercentage() float32 {
	return r.percentage
}

// NewRoadStats creates a new instance of RoadStats
func NewRoadStats(id int, name, iconPath string, totalRoads, completedRoads int, percentage float32) *RoadStats {
	return &RoadStats{
		id:             id,
		name:           name,
		iconPath:       iconPath,
		totalRoads:     totalRoads,
		completedRoads: completedRoads,
		percentage:     percentage,
	}
}

type RoadProgressStats struct {
	completed float32
	progress  float32
}

func (r *RoadProgressStats) GetCompleted() float32 {
	return r.completed
}

func (r *RoadProgressStats) GetProgress() float32 {
	return r.progress
}

func NewRoadProgressStats(inProgressRoads, completedRoads float32) *RoadProgressStats {
	return &RoadProgressStats{
		completed: completedRoads,
		progress:  inProgressRoads,
	}
}
