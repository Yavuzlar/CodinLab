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
func (s *labService) GetLabsFilter(userID string, labsId, labId int, isStartedStr, isFinishedStr string) ([]domains.Labs, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return nil, err
	}

	if userID == "" && labsId == 0 && labId == 0 && isStartedStr == "" && isFinishedStr == "" {
		return allLabs, nil
	}

	// Converting string to a bool
	var isStarted bool
	var isFinished bool
	if isStartedStr != "" {
		b, err := strconv.ParseBool(isStartedStr)
		if err != nil {
			return nil, err // Handle the error if conversion fails
		}
		isStarted = b
	}

	if isFinishedStr != "" {
		b, err := strconv.ParseBool(isFinishedStr)
		if err != nil {
			return nil, err // Handle the error if conversion fails
		}
		isFinished = b
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
			if lab.GetIsStarted() != isStarted {
				continue
			}
			if lab.GetIsFinished() != isFinished {
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

// Lab Sayfasindaki Dil bazli lab istatistikleri
func (s *labService) UserLanguageLabStats(userID string, language string) (domains.ProgrammingLanguageStats, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return domains.ProgrammingLanguageStats{}, err
	}

	totalLabs := 0
	solvedLabs := 0
	for _, labCollection := range allLabs {
		for _, lab := range labCollection.GetLabs() {
			for _, lang := range lab.GetLanguages() {
				if lang.GetLang() == language {
					totalLabs++
					if lab.GetIsFinished() {
						solvedLabs++
					}
					break
				}
			}
		}
	}
	returnval := *domains.NewProgrammingLanguageStats(
		totalLabs,
		solvedLabs,
		float64(solvedLabs)/float64(totalLabs)*100,
	)

	return returnval, nil
}

// Labs sayfasindaki genel lab istatistikleri
func (s *labService) UserGeneralLabStats(userID string) (domains.GeneralStats, error) {
	allLabs, err := s.getAllLabs(userID)
	if err != nil {
		return domains.GeneralStats{}, err
	}

	totalLabs := 0
	solvedLabs := 0
	easyLabs := 0
	mediumLabs := 0
	hardLabs := 0
	for _, labCollection := range allLabs {
		for _, lab := range labCollection.GetLabs() {
			totalLabs++
			if lab.GetIsFinished() {
				solvedLabs++
			}
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

	totalPercentage := float64(solvedLabs) / float64(totalLabs) * 100
	easyPercentage := float64(easyLabs) / float64(totalLabs) * 100
	mediumPercentage := float64(mediumLabs) / float64(totalLabs) * 100
	hardPercentage := float64(hardLabs) / float64(totalLabs) * 100
	returnVal := *domains.NewGeneralStats(
		totalLabs,
		totalPercentage,
		easyPercentage,
		mediumPercentage,
		hardPercentage,
		easyLabs,
		mediumLabs,
		hardLabs,
	)
	return returnVal, nil

}
