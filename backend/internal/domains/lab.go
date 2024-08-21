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

type Language struct {
	lang        string
	title       string
	description string
	note        string
	hint        string
}

type Test struct {
	input  []string
	output []string
}

type Param struct {
	name     string
	typeName string
}

type Labs struct {
	id          int
	name        string
	dockerImage string
	iconPath    string
	labs        []Lab
}

type Lab struct {
	id         int
	languages  []Language
	quest      Quest
	isStarted  bool
	isFinished bool
}

type Quest struct {
	difficulty int
	funcName   string
	tests      []Test
	params     []Param
}

func NewProgrammingLanguageStats(totalLabs, completedLabs int, percentage float64) *ProgrammingLanguageStats {
	return &ProgrammingLanguageStats{
		totalLabs:     totalLabs,
		completedLabs: completedLabs,
		percentage:    percentage,
	}
}

func (s *ProgrammingLanguageStats) SetTotalLabs(totalLabs int) {
	s.totalLabs = totalLabs
}

func (s *ProgrammingLanguageStats) GetTotalLabs() int {
	return s.totalLabs
}

func (s *ProgrammingLanguageStats) SetCompletedLabs(completedLabs int) {
	s.completedLabs = completedLabs
}

func (s *ProgrammingLanguageStats) GetCompletedLabs() int {
	return s.completedLabs
}

func (s *ProgrammingLanguageStats) SetPercentage(percentage float64) {
	s.percentage = percentage
}

func (s *ProgrammingLanguageStats) GetPercentage() float64 {
	return s.percentage
}

func NewGeneralStats(
	totalLabs int,
	totalPercentage float64,
	easyLabs int,
	easyPercentage float64,
	mediumLabs int,
	mediumPercentage float64,
	hardLabs int,
	hardPercentage float64,
) *GeneralStats {
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

func (s *GeneralStats) SetTotalLabs(totalLabs int) {
	s.totalLabs = totalLabs
}

func (s *GeneralStats) GetTotalLabs() int {
	return s.totalLabs
}

func (s *GeneralStats) SetTotalPercentage(totalPercentage float64) {
	s.totalPercentage = totalPercentage
}

func (s *GeneralStats) GetTotalPercentage() float64 {
	return s.totalPercentage
}

func (s *GeneralStats) SetEasyLabs(easyLabs int) {
	s.easyLabs = easyLabs
}

func (s *GeneralStats) GetEasyLabs() int {
	return s.easyLabs
}

func (s *GeneralStats) SetEasyPercentage(easyPercentage float64) {
	s.easyPercentage = easyPercentage
}

func (s *GeneralStats) GetEasyPercentage() float64 {
	return s.easyPercentage
}

func (s *GeneralStats) SetMediumLabs(mediumLabs int) {
	s.mediumLabs = mediumLabs
}

func (s *GeneralStats) GetMediumLabs() int {
	return s.mediumLabs
}

func (s *GeneralStats) SetMediumPercentage(mediumPercentage float64) {
	s.mediumPercentage = mediumPercentage
}

func (s *GeneralStats) GetMediumPercentage() float64 {
	return s.mediumPercentage
}

func (s *GeneralStats) SetHardLabs(hardLabs int) {
	s.hardLabs = hardLabs
}

func (s *GeneralStats) GetHardLabs() int {
	return s.hardLabs
}

func (s *GeneralStats) SetHardPercentage(hardPercentage float64) {
	s.hardPercentage = hardPercentage
}

func (s *GeneralStats) GetHardPercentage() float64 {
	return s.hardPercentage
}

func NewLanguage(lang, title, description, note, hint string) *Language {
	return &Language{
		lang:        lang,
		title:       title,
		description: description,
		note:        note,
		hint:        hint,
	}
}

func (l *Language) SetLang(lang string) {
	l.lang = lang
}

func (l *Language) GetLang() string {
	return l.lang
}

func (l *Language) SetTitle(title string) {
	l.title = title
}

func (l *Language) GetTitle() string {
	return l.title
}

func (l *Language) SetDescription(description string) {
	l.description = description
}

func (l *Language) GetDescription() string {
	return l.description
}

func (l *Language) SetNote(note string) {
	l.note = note
}

func (l *Language) GetNote() string {
	return l.note
}

func (l *Language) SetHint(hint string) {
	l.hint = hint
}

func (l *Language) GetHint() string {
	return l.hint
}

func NewTest(input, output []string) *Test {
	return &Test{
		input:  input,
		output: output,
	}
}

func (t *Test) SetInput(input []string) {
	t.input = input
}

func (t *Test) GetInput() []string {
	return t.input
}

func (t *Test) SetOutput(output []string) {
	t.output = output
}

func (t *Test) GetOutput() []string {
	return t.output
}

func NewParam(name, typeName string) *Param {
	return &Param{
		name:     name,
		typeName: typeName,
	}
}

func (p *Param) SetName(name string) {
	p.name = name
}

func (p *Param) GetName() string {
	return p.name
}

func (p *Param) SetType(typeName string) {
	p.typeName = typeName
}

func (p *Param) GetType() string {
	return p.typeName
}

func NewQuest(difficulty int, funcName string, tests []Test, params []Param) *Quest {
	return &Quest{
		difficulty: difficulty,
		funcName:   funcName,
		tests:      tests,
		params:     params,
	}
}

func (q *Quest) SetDifficulty(difficulty int) {
	q.difficulty = difficulty
}

func (q *Quest) GetDifficulty() int {
	return q.difficulty
}

func (q *Quest) SetFuncName(funcName string) {
	q.funcName = funcName
}

func (q *Quest) GetFuncName() string {
	return q.funcName
}

func (q *Quest) SetTests(tests []Test) {
	q.tests = tests
}

func (q *Quest) GetTests() []Test {
	return q.tests
}

func (q *Quest) SetParams(params []Param) {
	q.params = params
}

func (q *Quest) GetParams() []Param {
	return q.params
}

func NewLab(id int, languages []Language, quest Quest, isStarted, isFinished bool) *Lab {
	return &Lab{
		id:         id,
		languages:  languages,
		quest:      quest,
		isStarted:  isStarted,
		isFinished: isFinished,
	}
}

func (l *Lab) SetID(id int) {
	l.id = id
}

func (l *Lab) GetID() int {
	return l.id
}

func (l *Lab) SetLanguages(languages []Language) {
	l.languages = languages
}

func (l *Lab) GetLanguages() []Language {
	return l.languages
}

func (l *Lab) SetQuest(quest Quest) {
	l.quest = quest
}

func (l *Lab) GetQuest() Quest {
	return l.quest
}

func (l *Lab) SetIsStarted(isStarted bool) {
	l.isStarted = isStarted
}

func (l *Lab) GetIsStarted() bool {
	return l.isStarted
}

func (l *Lab) SetIsFinished(isFinished bool) {
	l.isFinished = isFinished
}

func (l *Lab) GetIsFinished() bool {
	return l.isFinished
}

func NewLabs(id int, name, dockerImage, iconPath string, labs []Lab) *Labs {
	return &Labs{
		id:          id,
		name:        name,
		dockerImage: dockerImage,
		iconPath:    iconPath,
		labs:        labs,
	}
}

func (l *Labs) SetID(id int) {
	l.id = id
}

func (l *Labs) GetID() int {
	return l.id
}

func (l *Labs) SetName(name string) {
	l.name = name
}

func (l *Labs) GetName() string {
	return l.name
}

func (l *Labs) SetDockerImage(dockerImage string) {
	l.dockerImage = dockerImage
}

func (l *Labs) GetDockerImage() string {
	return l.dockerImage
}

func (l *Labs) SetIconPath(iconPath string) {
	l.iconPath = iconPath
}

func (l *Labs) GetIconPath() string {
	return l.iconPath
}

func (l *Labs) SetLabs(labs []Lab) {
	l.labs = labs
}

func (l *Labs) GetLabs() []Lab {
	return l.labs
}
