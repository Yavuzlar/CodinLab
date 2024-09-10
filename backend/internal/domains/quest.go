package domains

// Test represents a test case for a function.
type Test struct {
	input  []interface{}
	output []interface{}
}

// Getter and Setter methods for Test
func (t *Test) GetInput() []interface{} {
	return t.input
}

func (t *Test) SetInput(input []interface{}) {
	t.input = input
}

func (t *Test) GetOutput() []interface{} {
	return t.output
}

func (t *Test) SetOutput(output []interface{}) {
	t.output = output
}

// Param represents a parameter of a function.
type Param struct {
	name string
	typ  string
}

// Getter and Setter methods for Param
func (p *Param) GetName() string {
	return p.name
}

func (p *Param) SetName(name string) {
	p.name = name
}

func (p *Param) GetType() string {
	return p.typ
}

func (p *Param) SetType(typ string) {
	p.typ = typ
}

// Return represents a parameter of a function.
type Returns struct {
	name string
	typ  string
}

// Getter and Setter methods for Return
func (p *Returns) GetName() string {
	return p.name
}

func (p *Returns) SetName(name string) {
	p.name = name
}

func (p *Returns) GetType() string {
	return p.typ
}

func (p *Returns) SetType(typ string) {
	p.typ = typ
}

type CodeTemplate struct {
	ProgrammingID int
	Frontend      string
	Template      string
	Check         string
}

// Quest represents a coding challenge or task.
type Quest struct {
	difficulty   int
	funcName     string
	tests        []Test
	params       []Param
	returns      []Returns
	questImports []string
	codeTemplate []CodeTemplate
}

// Getter and Setter methods for Quest
func (q *Quest) GetCodeTemplates() []CodeTemplate {
	return q.codeTemplate
}

func (q *Quest) SetCodeTemplate(codeTemplate []CodeTemplate) {
	q.codeTemplate = codeTemplate
}

func (q *Quest) GetDifficulty() int {
	return q.difficulty
}

func (q *Quest) SetDifficulty(difficulty int) {
	q.difficulty = difficulty
}

func (q *Quest) GetFuncName() string {
	return q.funcName
}

func (q *Quest) SetFuncName(funcName string) {
	q.funcName = funcName
}

func (q *Quest) GetTests() []Test {
	return q.tests
}

func (q *Quest) SetTests(tests []Test) {
	q.tests = tests
}

func (q *Quest) GetParams() []Param {
	return q.params
}

func (q *Quest) SetParams(params []Param) {
	q.params = params
}

func (q *Quest) GetReturns() []Returns {
	return q.returns
}

func (q *Quest) SetReturns(returns []Returns) {
	q.returns = returns
}

func (q *Quest) GetQuestImports() []string {
	return q.questImports
}

func (q *Quest) SetQuestImports(QuestImports []string) {
	q.questImports = QuestImports
}

// NewTest creates a new instance of Test
func NewTest(input, output []interface{}) *Test {
	return &Test{
		input:  input,
		output: output,
	}
}

// NewParam creates a new instance of Param
func NewParam(name, typ string) *Param {
	return &Param{
		name: name,
		typ:  typ,
	}
}

// NewReturn creates a new instance of Return
func NewReturn(name, typ string) *Returns {
	return &Returns{
		name: name,
		typ:  typ,
	}
}

func NewCodeTemplate(programmingID int, frontend, template, check string) *CodeTemplate {
	return &CodeTemplate{
		ProgrammingID: programmingID,
		Template:      template,
		Check:         check,
		Frontend:      frontend,
	}
}

// NewQuest creates a new instance of Quest
func NewQuest(difficulty int, funcName string, tests []Test, params []Param, returns []Returns, questImports []string, codeTemplates []CodeTemplate) *Quest {
	return &Quest{
		difficulty:   difficulty,
		funcName:     funcName,
		tests:        tests,
		params:       params,
		returns:      returns,
		questImports: questImports,
		codeTemplate: codeTemplates,
	}
}
