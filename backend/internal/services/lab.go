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

		var codeTemplates []domains.CodeTemplate
		for _, codeTemplateParser := range lab.Quest.CodeTemplates {
			codeTemplates = append(codeTemplates, *domains.NewCodeTemplate(codeTemplateParser.ProgrammingID, codeTemplateParser.TemplatePath))
		}

		quest := domains.NewQuest(lab.Quest.Difficulty, lab.Quest.FuncName, tests, params, returns, codeTemplates, lab.Quest.QuestImports)

		labIDString := strconv.Itoa(lab.ID)
		for _, lang := range lab.Quest.CodeTemplates {
			newLab := domains.NewLab(lab.ID, 0, languages, *quest, false, false)
			programmingID := strconv.Itoa(lang.ProgrammingID)
			logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, labIDString, domains.TypeLab, domains.ContentStarted)
			if err != nil {
				return nil, err
			}
			if len(logStartedStatus) > 0 {
				newLab.SetIsStarted(true)
			}

			logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, labIDString, domains.TypeLab, domains.ContentCompleted)
			if err != nil {
				return nil, err
			}
			if len(logFinishedStatus) > 0 {
				newLab.SetIsFinished(true)
			}
			newLab.SetProgrammingID(lang.ProgrammingID)
			labs = append(labs, *newLab)
		}

	}
	return labs, nil
}

// Fetch labs by filters
func (s *labService) GetLabsFilter(userID string, labId, programmingID int, isStarted, isFinished *bool) ([]domains.Lab, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	if userID == "" && labId == 0 && isStarted == nil && isFinished == nil && programmingID == 0 {
		return allLabs, nil
	}

	var labs []domains.Lab
	for _, lab := range allLabs {
		//found := false
		if labId != 0 && lab.GetID() != labId {
			continue
		}
		if isStarted != nil && lab.GetIsStarted() != *isStarted {
			continue
		}
		if isFinished != nil && lab.GetIsFinished() != *isFinished {
			continue
		}
		if lab.GetProgrammingID() != programmingID {
			continue
		}
		/* if programmingID != 0 {
			for _, ct := range lab.GetQuest().GetCodeTemplates() {
				if ct.GetProgrammingID() == programmingID {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		} */

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

// User lab level stats by programming language
func (s *labService) GetUserLanguageLabStats(userID string) (programmingLangugageStats []domains.ProgrammingLanguageStats, err error) {
	programmingLangugages, _ := s.parserService.GetInventory()

	for _, pl := range programmingLangugages {
		allLabs, _ := s.GetLabsFilter(userID, 0, pl.ID, nil, nil)
		completedLabs := 0

		for _, lab := range allLabs {
			if lab.GetIsStarted() && lab.GetIsFinished() {
				completedLabs++
			}
		}

		var completionRate float32
		if len(allLabs) > 0 {
			completionRate = float32((float32(completedLabs) / float32(len(allLabs))) * 100)
		} else {
			completionRate = 0
		}

		programmingLanguageStat := domains.NewProgrammingLanguageStats(
			pl.Name,
			pl.IconPath,
			len(allLabs),
			completedLabs,
			completionRate,
		)

		programmingLangugageStats = append(programmingLangugageStats, *programmingLanguageStat)
	}

	return
}

// User lab progress Statistics
func (s *labService) GetUserLabProgressStats(userID string) (userLabProgressStats *domains.UserLabProgressStats, err error) {
	progress := 0
	completed := 0
	totalLabs := 0

	programmingLangugages, _ := s.parserService.GetInventory()

	for _, pl := range programmingLangugages {
		allLabs, _ := s.GetLabsFilter(userID, 0, pl.ID, nil, nil)

		for _, lab := range allLabs {
			if lab.GetIsStarted() && !lab.GetIsFinished() {
				progress++
			}
			if lab.GetIsFinished() && lab.GetIsStarted() {
				completed++
			}
			totalLabs++
		}
	}
	progressPercentage := float32(float32(progress) / float32(totalLabs) * 100)
	completedPercentage := float32(float32(completed) / float32(totalLabs) * 100)

	userLabProgressStats = domains.NewsUserLabProgressStats(progressPercentage, completedPercentage)

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
