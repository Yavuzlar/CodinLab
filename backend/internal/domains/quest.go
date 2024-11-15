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

type CodeTemplate struct {
	programmingID int
	templatePath  string
}

func (ct *CodeTemplate) GetProgrammingID() int {
	return ct.programmingID
}

func (ct *CodeTemplate) SetProgrammingID(id int) {
	ct.programmingID = id
}

func (ct *CodeTemplate) GetTemplatePath() string {
	return ct.templatePath
}

func (ct *CodeTemplate) SetTemplatePath(templatePath string) {
	ct.templatePath = templatePath
}

// Quest represents a coding challenge or task.
type Quest struct {
	needAnswer   bool
	difficulty   int
	funcName     string
	tests        []Test
	codeTemplate []CodeTemplate
}

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

// NewTest creates a new instance of Test
func NewTest(input, output []interface{}) *Test {
	return &Test{
		input:  input,
		output: output,
	}
}

func NewCodeTemplate(programmingID int, templatePath string) *CodeTemplate {
	return &CodeTemplate{
		programmingID: programmingID,
		templatePath:  templatePath,
	}
}

// NewQuest creates a new instance of Quest
func NewQuest(difficulty int, funcName string, tests []Test, codeTemplates []CodeTemplate) *Quest {
	return &Quest{
		difficulty:   difficulty,
		funcName:     funcName,
		tests:        tests,
		codeTemplate: codeTemplates,
	}
}
