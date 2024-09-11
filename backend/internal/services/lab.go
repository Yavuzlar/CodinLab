package services

import (
	"context"
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
)

type labService struct {
	utils         IUtilService
	logService    domains.ILogService
	parserService domains.IParserService
}

func newLabService(
	utils IUtilService,
	logService domains.ILogService,
	parserService domains.IParserService,
) domains.ILabService {
	return &labService{
		utils:         utils,
		logService:    logService,
		parserService: parserService,
	}
}

// Customize and bring all labs
func (s *labService) getAllLabs(userID string) ([]domains.Lab, error) {
	src, err := s.parserService.GetLabs()
	if err != nil {
		return nil, err
	}

	var labs []domains.Lab
	for _, lab := range src {
		var languages []domains.LanguageLab
		for _, lang := range lab.Languages {
			languages = append(languages, *domains.NewLanguageLab(lang.Lang, lang.Title, lang.Description, lang.Note, lang.Hint))
		}

		var tests []domains.Test
		for _, test := range lab.Quest.Tests {
			tests = append(tests, *domains.NewTest(test.Input, test.Output))
		}

		var params []domains.Param
		for _, param := range lab.Quest.Params {
			params = append(params, *domains.NewParam(param.Name, param.Type))
		}

		var returns []domains.Returns
		for _, returnedParam := range lab.Quest.Returns {
			returns = append(returns, *domains.NewReturn(returnedParam.Name, returnedParam.Type))
		}

		var questImports []string
		for _, imp := range lab.Quest.QuestImports {
			questImports = append(questImports, imp)
		}

		var codeTemplates []domains.CodeTemplate
		for _, codeTemplateParser := range lab.Quest.CodeTemplates {
			codeTemplates = append(codeTemplates, *domains.NewCodeTemplate(codeTemplateParser.ProgrammingID, codeTemplateParser.Frontend, codeTemplateParser.Template, codeTemplateParser.Check))
		}

		quest := domains.NewQuest(lab.Quest.Difficulty, lab.Quest.FuncName, tests, params, returns, questImports, codeTemplates)
		newLab := domains.NewLab(lab.ID, languages, *quest, false, false)

		labIDString := strconv.Itoa(lab.ID)
		logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", labIDString, domains.TypeLab, domains.ContentStarted)
		if err != nil {
			return nil, err
		}
		if len(logStartedStatus) > 0 {
			newLab.SetIsStarted(true)
		}

		logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", labIDString, domains.TypeLab, domains.ContentCompleted)
		if err != nil {
			return nil, err
		}
		if len(logFinishedStatus) > 0 {
			newLab.SetIsFinished(true)
		}

		labs = append(labs, *newLab)
	}

	return labs, nil
}

// Fetch labs by filters
func (s *labService) GetLabsFilter(userID string, labId int, isStarted, isFinished *bool) ([]domains.Lab, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	if userID == "" && labId == 0 && isStarted == nil && isFinished == nil {
		return allLabs, nil
	}

	var labs []domains.Lab
	for _, lab := range allLabs {
		if labId != 0 && lab.GetID() != labId {
			continue
		}
		if isStarted != nil && lab.GetIsStarted() != *isStarted {
			continue
		}
		if isFinished != nil && lab.GetIsFinished() != *isFinished {
			continue
		}
		labs = append(labs, lab)

	}

	return labs, nil
}

func (s *labService) GetLabByID(userID string, labID int) (lab *domains.Lab, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	for _, lab := range allLabs {
		if lab.GetID() == labID {
			return &lab, nil
		}
	}

	return nil, err
}

// Fetch labs by filters
func (s *labService) CountLabsFilter(userID string, labId int, isStarted, isFinished *bool) (counter int, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return 0, err
	}

	if userID == "" && labId == 0 && isStarted == nil && isFinished == nil {
		return 0, nil
	}

	for _, lab := range allLabs {
		if labId != 0 && lab.GetID() != labId {
			continue
		}
		if isStarted != nil && lab.GetIsStarted() != *isStarted {
			continue
		}
		if isFinished != nil && lab.GetIsFinished() != *isFinished {
			continue
		}
		counter++

	}

	return counter, nil
}

// User lab level stats by programming language
func (s *labService) GetUserLanguageLabStats(userID string) (programmingLangugageStats *domains.ProgrammingLanguageStats, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return
	}
	programmingLangugages, _ := s.parserService.GetInventory()

	completedLabs := 0

	for _, lab := range allLabs {
		if lab.GetIsFinished() {
			completedLabs++
		}
	}
	programmingLangugageStats = domains.NewProgrammingLanguageStats(
		len(allLabs)*len(programmingLangugages),
		completedLabs,
		float32((float32(completedLabs)/float32(len(allLabs)*len(programmingLangugages)))*100),
	)

	return
}

// User lab progress Statistics
func (s *labService) GetUserLabProgressStats(userID string) (userLabProgressStats domains.UserLabProgressStats, err error) {
	trueValue := true
	falseValue := false

	progressLabs, err := s.CountLabsFilter(userID, 0, &trueValue, &falseValue)
	if err != nil {
		return
	}
	completedLabs, err := s.CountLabsFilter(userID, 0, &trueValue, &trueValue)
	if err != nil {
		return
	}

	totalLabs, err := s.CountLabsFilter(userID, 0, nil, nil)
	if err != nil {
		return
	}
	programmingLangugages, _ := s.parserService.GetInventory()

	progressPercentage := float32(float32(progressLabs) / float32(totalLabs*len(programmingLangugages)) * 100)
	completedPercentage := float32(float32(completedLabs) / float32(totalLabs*len(programmingLangugages)) * 100)

	userLabProgressStats = *domains.NewsUserLabProgressStats(progressPercentage, completedPercentage)

	return
}

// User lab difficulty Statistics
func (s *labService) GetUserLabDifficultyStats(userID string) (userLabDifficultyStats domains.UserLabDifficultyStats, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return
	}

	totalLabs := 0
	easyLabs := 0
	mediumLabs := 0
	hardLabs := 0

	for _, lab := range allLabs {
		totalLabs++
		if lab.GetIsFinished() {
			switch lab.GetQuest().GetDifficulty() {
			case 1:
				easyLabs++
			case 2:
				mediumLabs++
			case 3:
				hardLabs++
			}
		}
	}
	programmingLangugages, _ := s.parserService.GetInventory()
	totalLabs = totalLabs * len(programmingLangugages)

	easyPercentage := float32(float32(easyLabs) / float32(totalLabs) * 100)
	mediumPercentage := float32(float32(mediumLabs) / float32(totalLabs) * 100)
	hardPercentage := float32(float32(hardLabs) / float32(totalLabs) * 100)

	userLabDifficultyStats = *domains.NewserLabLevelStats(
		easyPercentage,
		mediumPercentage,
		hardPercentage,
	)
	return
}
