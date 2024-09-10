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
func (s *labService) getAllLabs(userID string) ([]domains.Labs, error) {
	src, err := s.parserService.GetLabs()
	if err != nil {
		return nil, err
	}

	var labs []domains.Labs
	for _, labCollection := range src {
		var newLabList []domains.Lab

		for _, lab := range labCollection.Labs {
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

			quest := domains.NewQuest(lab.Quest.Difficulty, lab.Quest.FuncName, tests, params, returns, questImports)
			newLab := domains.NewLab(lab.ID, languages, *quest, false, false)

			labIDString := strconv.Itoa(lab.ID)
			labCollectionIDString := strconv.Itoa(labCollection.ID)
			logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, labCollectionIDString, labIDString, domains.TypeLab, domains.ContentStarted)
			if err != nil {
				return nil, err
			}
			if len(logStartedStatus) > 0 {
				newLab.SetIsStarted(true)
			}

			logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, labCollectionIDString, labIDString, domains.TypeLab, domains.ContentCompleted)
			if err != nil {
				return nil, err
			}
			if len(logFinishedStatus) > 0 {
				newLab.SetIsFinished(true)
			}

			newLabList = append(newLabList, *newLab)
		}
		labs = append(labs, *domains.NewLabs(labCollection.ID, labCollection.Name, labCollection.DockerImage, labCollection.IconPath, labCollection.FileExtension, labCollection.TemplatePath, labCollection.Cmd, newLabList))
	}

	return labs, nil
}

// Fetch labs by filters
func (s *labService) GetLabsFilter(userID string, programmingID, labId int, isStarted, isFinished *bool) ([]domains.Labs, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	if userID == "" && programmingID == 0 && labId == 0 && isStarted == nil && isFinished == nil {
		return allLabs, nil
	}

	var filteredLabs []domains.Labs
	for _, labCollection := range allLabs {
		if programmingID != 0 && labCollection.GetID() != programmingID {
			continue
		}

		var newLabList []domains.Lab
		for _, lab := range labCollection.GetLabs() {

			if labId != 0 && lab.GetID() != labId {
				continue
			}
			if isStarted != nil && lab.GetIsStarted() != *isStarted {
				continue
			}
			if isFinished != nil && lab.GetIsFinished() != *isFinished {
				continue
			}
			newLabList = append(newLabList, lab)
		}

		if len(newLabList) > 0 {
			filteredLabs = append(filteredLabs, *domains.NewLabs(
				labCollection.GetID(),
				labCollection.GetName(),
				labCollection.GetDockerImage(),
				labCollection.GetIconPath(),
				labCollection.GetFileExtension(),
				labCollection.GetTemplatePath(),
				labCollection.GetCmd(),
				newLabList,
			))
		}
	}

	return filteredLabs, nil
}

func (s *labService) GetLabByID(userID string, programmingID, labID int) (lab *domains.Lab, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	for _, programmingLang := range allLabs {
		if programmingLang.GetID() == programmingID {
			for _, lab := range programmingLang.GetLabs() {
				if lab.GetID() == labID {
					return &lab, nil
				}
			}
		}
	}

	return nil, err
}

// Fetch labs by filters
func (s *labService) CountLabsFilter(userID string, programmingID, labId int, isStarted, isFinished *bool) (counter int, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return 0, err
	}

	if userID == "" && programmingID == 0 && labId == 0 && isStarted == nil && isFinished == nil {
		return 0, nil
	}

	for _, labCollection := range allLabs {
		if programmingID != 0 && labCollection.GetID() != programmingID {
			continue
		}

		for _, lab := range labCollection.GetLabs() {

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
	}

	return counter, nil
}

// User lab level stats by programming language
func (s *labService) GetUserLanguageLabStats(userID string) (programmingLangugageStats []domains.ProgrammingLanguageStats, err error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return
	}

	totalLabs := 0
	completedLabs := 0

	for _, labCollection := range allLabs {
		totalLabs = 0
		completedLabs = 0
		for _, lab := range labCollection.GetLabs() {
			totalLabs++
			if lab.GetIsFinished() {
				completedLabs++
			}
		}

		newProgrammingLanguageStats := domains.NewProgrammingLanguageStats(
			labCollection.GetID(),
			labCollection.GetName(),
			labCollection.GetIconPath(),
			totalLabs,
			completedLabs,
			float32((float32(completedLabs)/float32(totalLabs))*100),
		)
		programmingLangugageStats = append(programmingLangugageStats, *newProgrammingLanguageStats)

	}

	return
}

// User lab progress Statistics
func (s *labService) GetUserLabProgressStats(userID string) (userLabProgressStats domains.UserLabProgressStats, err error) {
	trueValue := true
	falseValue := false

	progressLabs, err := s.CountLabsFilter(userID, 0, 0, &trueValue, &falseValue)
	if err != nil {
		return
	}
	completedLabs, err := s.CountLabsFilter(userID, 0, 0, &trueValue, &trueValue)
	if err != nil {
		return
	}

	totalLabs, err := s.CountLabsFilter(userID, 0, 0, nil, nil)
	if err != nil {
		return
	}

	progressPercentage := float32(float32(progressLabs) / float32(totalLabs) * 100)
	completedPercentage := float32(float32(completedLabs) / float32(totalLabs) * 100)

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

	for _, labCollection := range allLabs {
		for _, lab := range labCollection.GetLabs() {
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
