package domains

// ILabService is the interface that provides the methods for the lab service.
type ILabService interface {
	GetLabsFilter(userID string, labId, programmingID int, isStarted, isFinished *bool) ([]Lab, error)
	GetUserLanguageLabStats(userID string) (programmingLangugageStats []ProgrammingLanguageStats, err error)
	GetUserLabDifficultyStats(userID string) (userLabDifficultyStats UserLabDifficultyStats, err error)
	GetUserLabProgressStats(userID string) (userLabProgressStats *UserLabProgressStats, err error)
	GetLabByID(userID string, labID int) (lab *Lab, err error)
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

func (p *ProgrammingLanguageStats) GetIconPath() string {
	return p.iconPath
}

func (p *ProgrammingLanguageStats) SetIconPath(iconPath string) {
	p.iconPath = iconPath
}

func (p *ProgrammingLanguageStats) GetName() string {
	return p.name
}

func (p *ProgrammingLanguageStats) SetName(name string) {
	p.name = name
}

func (p *ProgrammingLanguageStats) GetID() int {
	return p.id
}

func (p *ProgrammingLanguageStats) SetID(id int) {
	p.id = id
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

// Lab represents a specific coding lab exercise.
type Lab struct {
	id            int
	programmingID int
	languages     []LanguageLab
	quest         Quest
	isStarted     bool
	isFinished    bool
}

// Getter and Setter methods for Lab
func (l *Lab) GetID() int {
	return l.id
}

func (l *Lab) SetID(id int) {
	l.id = id
}

func (l *Lab) GetProgrammingID() int {
	return l.programmingID
}

func (l *Lab) SetProgrammingID(programmingID int) {
	l.programmingID = programmingID
}

func (l *Lab) GetLanguages() []LanguageLab {
	return l.languages
}

func (l *Lab) SetLanguages(languages []LanguageLab) {
	l.languages = languages
}

func (l *Lab) GetQuest() *Quest {
	return &l.quest
}

func (l *Lab) SetQuest(quest Quest) {
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

// NewProgrammingLanguageStats creates a new instance of ProgrammingLanguageStats
func NewProgrammingLanguageStats(name, iconPath string, totalLabs, completedLabs int, percentage float32) *ProgrammingLanguageStats {
	return &ProgrammingLanguageStats{
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

// NewLab creates a new instance of Lab
func NewLab(id, programmingID int, languages []LanguageLab, quest Quest, isStarted, isFinished bool) *Lab {
	return &Lab{
		id:            id,
		languages:     languages,
		programmingID: programmingID,
		quest:         quest,
		isStarted:     isStarted,
		isFinished:    isFinished,
	}
}
