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

			var tests []domains.TestLab
			for _, test := range lab.Quest.Tests {
				tests = append(tests, *domains.NewTestLab(test.Input, test.Output))
			}

			var params []domains.ParamLab
			for _, param := range lab.Quest.Params {
				params = append(params, *domains.NewParamLab(param.Name, param.Type))
			}

			quest := domains.NewQuestLab(lab.Quest.Difficulty, lab.Quest.FuncName, tests, params)
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

			newLabList = append(newLabList, *newLab)
		}
		labs = append(labs, *domains.NewLabs(labCollection.ID, labCollection.Name, labCollection.DockerImage, labCollection.IconPath, newLabList))
	}

	return labs, nil
}

// Fetch labs by filters
func (s *labService) GetLabsFilter(userID string, labsId, labId int, isStarted, isFinished *bool) ([]domains.Labs, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	if userID == "" && labsId == 0 && labId == 0 && isStarted == nil && isFinished == nil {
		return allLabs, nil
	}

	var filteredLabs []domains.Labs
	for _, labCollection := range allLabs {
		if labsId != 0 && labCollection.GetID() != labsId {
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
				newLabList,
			))
		}
	}

	return filteredLabs, nil
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
