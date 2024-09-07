package domains

// ILabService is the interface that provides the methods for the lab service.
type ILabService interface {
	GetLabsFilter(userID string, programmingID, labId int, isStarted, isFinished *bool) ([]Labs, error)
	GetUserLanguageLabStats(userID string) (programmingLangugageStats []ProgrammingLanguageStats, err error)
	GetUserLabDifficultyStats(userID string) (userLabDifficultyStats UserLabDifficultyStats, err error)
	GetUserLabProgressStats(userID string) (userLabProgressStats UserLabProgressStats, err error)
	CountLabsFilter(userID string, programmingID, labId int, isStarted, isFinished *bool) (counter int, err error)
	CodeTemplateGenerator(programmingName, templatePathObject, content, funcName string, tests []TestLab) (string, error)
	GetLabByID(userID string, programmingID, labID int) (lab *Lab, err error)
}

// ProgrammingLanguageStats represents the statistics for a specific language lab.
type ProgrammingLanguageStats struct {
	id            int
	name          string
	iconPath      string
	totalLabs     int
	completedLabs int
	percentage    float32
}

// Getter and Setter methods for ProgrammingLanguageStats
func (p *ProgrammingLanguageStats) GetTotalLabs() int {
	return p.totalLabs
}

func (p *ProgrammingLanguageStats) SetTotalLabs(totalLabs int) {
	p.totalLabs = totalLabs
}

func (p *ProgrammingLanguageStats) GetID() int {
	return p.id
}

func (p *ProgrammingLanguageStats) SetID(id int) {
	p.id = id
}

func (p *ProgrammingLanguageStats) GetName() string {
	return p.name
}

func (p *ProgrammingLanguageStats) SetName(name string) {
	p.name = name
}

func (p *ProgrammingLanguageStats) GetIconPath() string {
	return p.iconPath
}

func (p *ProgrammingLanguageStats) SetIconPath(iconPath string) {
	p.iconPath = iconPath
}

func (p *ProgrammingLanguageStats) GetCompletedLabs() int {
	return p.completedLabs
}

func (p *ProgrammingLanguageStats) SetCompletedLabs(completedLabs int) {
	p.completedLabs = completedLabs
}

func (p *ProgrammingLanguageStats) GetPercentage() float32 {
	return p.percentage
}

func (p *ProgrammingLanguageStats) SetPercentage(percentage float32) {
	p.percentage = percentage
}

// UserLabDifficultyStats represents the DTO for user general lab statistics
type UserLabDifficultyStats struct {
	easyPercentage   float32
	mediumPercentage float32
	hardPercentage   float32
}

// UserLabProgressStats  for user general lab progresss statistics
type UserLabProgressStats struct {
	progress  float32
	completed float32
}

func (ps *UserLabProgressStats) GetProgress() float32 {
	return ps.progress
}

func (ps *UserLabProgressStats) SetProgress(progress float32) {
	ps.progress = progress
}

func (ps *UserLabProgressStats) GetCompleted() float32 {
	return ps.completed
}

func (ps *UserLabProgressStats) SetCompleted(completed float32) {
	ps.completed = completed
}

// Getter and Setter methods for GeneralStats
func (g *UserLabDifficultyStats) GetEasyPercentage() float32 {
	return g.easyPercentage
}

func (g *UserLabDifficultyStats) SetEasyPercentage(easyPercentage float32) {
	g.easyPercentage = easyPercentage
}

func (g *UserLabDifficultyStats) GetMediumPercentage() float32 {
	return g.mediumPercentage
}

func (g *UserLabDifficultyStats) SetMediumPercentage(mediumPercentage float32) {
	g.mediumPercentage = mediumPercentage
}

func (g *UserLabDifficultyStats) GetHardPercentage() float32 {
	return g.hardPercentage
}

func (g *UserLabDifficultyStats) SetHardPercentage(hardPercentage float32) {
	g.hardPercentage = hardPercentage
}

// LanguageLab represents the details of a programming language.
type LanguageLab struct {
	lang        string
	title       string
	description string
	note        string
	hint        string
}

// Getter and Setter methods for Language
func (l *LanguageLab) GetLang() string {
	return l.lang
}

func (l *LanguageLab) SetLang(lang string) {
	l.lang = lang
}

func (l *LanguageLab) GetTitle() string {
	return l.title
}

func (l *LanguageLab) SetTitle(title string) {
	l.title = title
}

func (l *LanguageLab) GetDescription() string {
	return l.description
}

func (l *LanguageLab) SetDescription(description string) {
	l.description = description
}

func (l *LanguageLab) GetNote() string {
	return l.note
}

func (l *LanguageLab) SetNote(note string) {
	l.note = note
}

func (l *LanguageLab) GetHint() string {
	return l.hint
}

func (l *LanguageLab) SetHint(hint string) {
	l.hint = hint
}

// TestLab represents a test case for a function.
type TestLab struct {
	input  []string
	output []string
}

// Getter and Setter methods for Test
func (t *TestLab) GetInput() []string {
	return t.input
}

func (t *TestLab) SetInput(input []string) {
	t.input = input
}

func (t *TestLab) GetOutput() []string {
	return t.output
}

func (t *TestLab) SetOutput(output []string) {
	t.output = output
}

// ParamLab represents a parameter of a function.
type ParamLab struct {
	name string
	typ  string
}

// Getter and Setter methods for Param
func (p *ParamLab) GetName() string {
	return p.name
}

func (p *ParamLab) SetName(name string) {
	p.name = name
}

func (p *ParamLab) GetType() string {
	return p.typ
}

func (p *ParamLab) SetType(typ string) {
	p.typ = typ
}

// ReturnLab represents a parameter of a function.
type ReturnLab struct {
	name string
	typ  string
}

// Getter and Setter methods for Return
func (p *ReturnLab) GetName() string {
	return p.name
}

func (p *ReturnLab) SetName(name string) {
	p.name = name
}

func (p *ReturnLab) GetType() string {
	return p.typ
}

func (p *ReturnLab) SetType(typ string) {
	p.typ = typ
}

// QuestLab represents a coding challenge or task.
type QuestLab struct {
	difficulty int
	funcName   string
	tests      []TestLab
	params     []ParamLab
	returns    []ReturnLab
}

// Getter and Setter methods for Quest
func (q *QuestLab) GetDifficulty() int {
	return q.difficulty
}

func (q *QuestLab) SetDifficulty(difficulty int) {
	q.difficulty = difficulty
}

func (q *QuestLab) GetFuncName() string {
	return q.funcName
}

func (q *QuestLab) SetFuncName(funcName string) {
	q.funcName = funcName
}

func (q *QuestLab) GetTests() []TestLab {
	return q.tests
}

func (q *QuestLab) SetTests(tests []TestLab) {
	q.tests = tests
}

func (q *QuestLab) GetParams() []ParamLab {
	return q.params
}

func (q *QuestLab) SetParams(params []ParamLab) {
	q.params = params
}

func (q *QuestLab) GetReturns() []ReturnLab {
	return q.returns
}

func (q *QuestLab) SetReturns(returns []ReturnLab) {
	q.returns = returns
}

// Lab represents a specific coding lab exercise.
type Lab struct {
	id         int
	languages  []LanguageLab
	quest      QuestLab
	isStarted  bool
	isFinished bool
}

// Getter and Setter methods for Lab
func (l *Lab) GetID() int {
	return l.id
}

func (l *Lab) SetID(id int) {
	l.id = id
}

func (l *Lab) GetLanguages() []LanguageLab {
	return l.languages
}

func (l *Lab) SetLanguages(languages []LanguageLab) {
	l.languages = languages
}

func (l *Lab) GetQuest() *QuestLab {
	return &l.quest
}

func (l *Lab) SetQuest(quest QuestLab) {
	l.quest = quest
}

func (l *Lab) GetIsStarted() bool {
	return l.isStarted
}

func (l *Lab) SetIsStarted(isStarted bool) {
	l.isStarted = isStarted
}

func (l *Lab) GetIsFinished() bool {
	return l.isFinished
}

func (l *Lab) SetIsFinished(isFinished bool) {
	l.isFinished = isFinished
}

// Labs represents a collection of labs grouped together.
type Labs struct {
	id            int // programming ID
	name          string
	dockerImage   string
	iconPath      string
	cmd           []string
	fileExtension string
	templatePath  string
	labs          []Lab
}

// Getter and Setter methods for Labs
func (l *Labs) GetID() int {
	return l.id
}

func (l *Labs) SetID(id int) {
	l.id = id
}

func (l *Labs) GetName() string {
	return l.name
}

func (l *Labs) SetName(name string) {
	l.name = name
}

func (l *Labs) GetDockerImage() string {
	return l.dockerImage
}

func (l *Labs) SetDockerImage(dockerImage string) {
	l.dockerImage = dockerImage
}

func (l *Labs) GetIconPath() string {
	return l.iconPath
}

func (l *Labs) SetIconPath(iconPath string) {
	l.iconPath = iconPath
}

func (l *Labs) GetCmd() []string {
	return l.cmd
}

func (l *Labs) SetCmd(cmd []string) {
	l.cmd = cmd
}

func (l *Labs) GetFileExtension() string {
	return l.fileExtension
}

func (l *Labs) SetFileExtension(fileExtension string) {
	l.fileExtension = fileExtension
}

func (l *Labs) GetTemplatePath() string {
	return l.templatePath
}

func (l *Labs) SetTemplatePath(templatePath string) {
	l.templatePath = templatePath
}

func (l *Labs) GetLabs() []Lab {
	return l.labs
}

func (l *Labs) SetLabs(labs []Lab) {
	l.labs = labs
}

// NewProgrammingLanguageStats creates a new instance of ProgrammingLanguageStats
func NewProgrammingLanguageStats(id int, name, iconPath string, totalLabs, completedLabs int, percentage float32) *ProgrammingLanguageStats {
	return &ProgrammingLanguageStats{
		id:            id,
		name:          name,
		iconPath:      iconPath,
		totalLabs:     totalLabs,
		completedLabs: completedLabs,
		percentage:    percentage,
	}
}

// NewGeneralStats creates a new instance of GeneralStats
func NewserLabLevelStats(easyPercentage, mediumPercentage, hardPercentage float32) *UserLabDifficultyStats {
	return &UserLabDifficultyStats{
		easyPercentage:   easyPercentage,
		mediumPercentage: mediumPercentage,
		hardPercentage:   hardPercentage,
	}
}

// NewsUserLabProgressStats creates a new instance of GeneralStats
func NewsUserLabProgressStats(progress, completed float32) *UserLabProgressStats {
	return &UserLabProgressStats{
		progress:  progress,
		completed: completed,
	}
}

// NewLanguageLab creates a new instance of Language
func NewLanguageLab(lang, title, description, note, hint string) *LanguageLab {
	return &LanguageLab{
		lang:        lang,
		title:       title,
		description: description,
		note:        note,
		hint:        hint,
	}
}

// NewTestLab creates a new instance of Test
func NewTestLab(input, output []string) *TestLab {
	return &TestLab{
		input:  input,
		output: output,
	}
}

// NewParamLab creates a new instance of Param
func NewParamLab(name, typ string) *ParamLab {
	return &ParamLab{
		name: name,
		typ:  typ,
	}
}

// NewReturnLab creates a new instance of Return
func NewReturnLab(name, typ string) *ReturnLab {
	return &ReturnLab{
		name: name,
		typ:  typ,
	}
}

// NewQuestLab creates a new instance of Quest
func NewQuestLab(difficulty int, funcName string, tests []TestLab, params []ParamLab, returns []ReturnLab) *QuestLab {
	return &QuestLab{
		difficulty: difficulty,
		funcName:   funcName,
		tests:      tests,
		params:     params,
		returns:    returns,
	}
}

// NewLab creates a new instance of Lab
func NewLab(id int, languages []LanguageLab, quest QuestLab, isStarted, isFinished bool) *Lab {
	return &Lab{
		id:         id,
		languages:  languages,
		quest:      quest,
		isStarted:  isStarted,
		isFinished: isFinished,
	}
}

// NewLabs creates a new instance of Labs
func NewLabs(id int, name, dockerImage, iconPath, fileExtension, templatePath string, cmd []string, labs []Lab) *Labs {
	return &Labs{
		id:            id,
		name:          name,
		dockerImage:   dockerImage,
		iconPath:      iconPath,
		cmd:           cmd,
		fileExtension: fileExtension,
		templatePath:  templatePath,
		labs:          labs,
	}
}
