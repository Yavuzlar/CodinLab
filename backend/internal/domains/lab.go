package domains

// ILabService is the interface that provides the methods for the lab service.
type ILabService interface {
	GetLabsFilter(userID string, labsId, labId int, isStarted, isFinished *bool) ([]Labs, error)
	UserLanguageLabStats(userID string, language string) (ProgrammingLanguageStats, error)
	UserGeneralLabStats(userID string) (GeneralStats, error)
}

// ProgrammingLanguageStats represents the statistics for a specific language lab.
type ProgrammingLanguageStats struct {
	totalLabs     int
	completedLabs int
	percentage    float64
}

// Getter and Setter methods for ProgrammingLanguageStats
func (p *ProgrammingLanguageStats) GetTotalLabs() int {
	return p.totalLabs
}

func (p *ProgrammingLanguageStats) SetTotalLabs(totalLabs int) {
	p.totalLabs = totalLabs
}

func (p *ProgrammingLanguageStats) GetCompletedLabs() int {
	return p.completedLabs
}

func (p *ProgrammingLanguageStats) SetCompletedLabs(completedLabs int) {
	p.completedLabs = completedLabs
}

func (p *ProgrammingLanguageStats) GetPercentage() float64 {
	return p.percentage
}

func (p *ProgrammingLanguageStats) SetPercentage(percentage float64) {
	p.percentage = percentage
}

// GeneralStats represents the DTO for user general lab statistics
type GeneralStats struct {
	totalLabs        int
	totalPercentage  float64
	easyLabs         int
	easyPercentage   float64
	mediumLabs       int
	mediumPercentage float64
	hardLabs         int
	hardPercentage   float64
}

// Getter and Setter methods for GeneralStats
func (g *GeneralStats) GetTotalLabs() int {
	return g.totalLabs
}

func (g *GeneralStats) SetTotalLabs(totalLabs int) {
	g.totalLabs = totalLabs
}

func (g *GeneralStats) GetTotalPercentage() float64 {
	return g.totalPercentage
}

func (g *GeneralStats) SetTotalPercentage(totalPercentage float64) {
	g.totalPercentage = totalPercentage
}

func (g *GeneralStats) GetEasyLabs() int {
	return g.easyLabs
}

func (g *GeneralStats) SetEasyLabs(easyLabs int) {
	g.easyLabs = easyLabs
}

func (g *GeneralStats) GetEasyPercentage() float64 {
	return g.easyPercentage
}

func (g *GeneralStats) SetEasyPercentage(easyPercentage float64) {
	g.easyPercentage = easyPercentage
}

func (g *GeneralStats) GetMediumLabs() int {
	return g.mediumLabs
}

func (g *GeneralStats) SetMediumLabs(mediumLabs int) {
	g.mediumLabs = mediumLabs
}

func (g *GeneralStats) GetMediumPercentage() float64 {
	return g.mediumPercentage
}

func (g *GeneralStats) SetMediumPercentage(mediumPercentage float64) {
	g.mediumPercentage = mediumPercentage
}

func (g *GeneralStats) GetHardLabs() int {
	return g.hardLabs
}

func (g *GeneralStats) SetHardLabs(hardLabs int) {
	g.hardLabs = hardLabs
}

func (g *GeneralStats) GetHardPercentage() float64 {
	return g.hardPercentage
}

func (g *GeneralStats) SetHardPercentage(hardPercentage float64) {
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

// QuestLab represents a coding challenge or task.
type QuestLab struct {
	difficulty int
	funcName   string
	tests      []TestLab
	params     []ParamLab
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
	id          int
	name        string
	dockerImage string
	iconPath    string
	labs        []Lab
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

func (l *Labs) GetLabs() []Lab {
	return l.labs
}

func (l *Labs) SetLabs(labs []Lab) {
	l.labs = labs
}

// NewProgrammingLanguageStats creates a new instance of ProgrammingLanguageStats
func NewProgrammingLanguageStats(totalLabs, completedLabs int, percentage float64) *ProgrammingLanguageStats {
	return &ProgrammingLanguageStats{
		totalLabs:     totalLabs,
		completedLabs: completedLabs,
		percentage:    percentage,
	}
}

// NewGeneralStats creates a new instance of GeneralStats
func NewGeneralStats(totalLabs int, totalPercentage, easyPercentage, mediumPercentage, hardPercentage float64, easyLabs, mediumLabs, hardLabs int) *GeneralStats {
	return &GeneralStats{
		totalLabs:        totalLabs,
		totalPercentage:  totalPercentage,
		easyLabs:         easyLabs,
		easyPercentage:   easyPercentage,
		mediumLabs:       mediumLabs,
		mediumPercentage: mediumPercentage,
		hardLabs:         hardLabs,
		hardPercentage:   hardPercentage,
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

// NewQuestLab creates a new instance of Quest
func NewQuestLab(difficulty int, funcName string, tests []TestLab, params []ParamLab) *QuestLab {
	return &QuestLab{
		difficulty: difficulty,
		funcName:   funcName,
		tests:      tests,
		params:     params,
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
func NewLabs(id int, name, dockerImage, iconPath string, labs []Lab) *Labs {
	return &Labs{
		id:          id,
		name:        name,
		dockerImage: dockerImage,
		iconPath:    iconPath,
		labs:        labs,
	}
}
