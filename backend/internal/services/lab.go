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

			var languages []domains.Language
			for _, lang := range lab.Languages {
				languages = append(languages, *domains.NewLanguage(
					lang.Lang,
					lang.Title,
					lang.Description,
					lang.Note,
					lang.Hint,
					//Language parametreleri buraya eklenecek
				))
			}

			var tests []domains.Test
			for _, test := range lab.Quest.Tests {
				tests = append(tests, *domains.NewTest(
					test.Input,
					test.Output,
					//Test parametreleri buraya eklenecek
				))
			}

			var params []domains.Param
			for _, param := range lab.Quest.Params {
				params = append(params, *domains.NewParam(
					param.Name,
					param.Type,
					//Param parametreleri buraya eklenecek
				))
			}

			quest := *domains.NewQuest(
				lab.Quest.Difficulty,
				lab.Quest.FuncName,
				tests,
				params,
				//Quest parametreleri buraya eklenecek
			)

			newLab := *domains.NewLab(
				lab.ID,
				languages,
				quest,
				false, // Varsayılan değer false
				false, // Varsayılan değer false
				//Lab parametreleri buraya eklenecek
			)
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

			newLabList = append(newLabList, newLab)
		}

		labs = append(labs, *domains.NewLabs(
			labCollection.ID,
			labCollection.Name,
			labCollection.DockerImage,
			labCollection.IconPath,
			newLabList,
		))
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
		//labs structı için filtreleme eklenebilir.

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

			//lab structı için filtreleme eklenebilir.

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
	returnval := domains.ProgrammingLanguageStats{}

	returnval = *domains.NewProgrammingLanguageStats(
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
			difficulty := lab.GetQuest()
			switch difficulty.GetDifficulty() {
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
	returnval := *domains.NewGeneralStats(
		totalLabs,
		totalPercentage,
		easyLabs,
		easyPercentage,
		mediumLabs,
		mediumPercentage,
		hardLabs,
		hardPercentage,
	)
	return returnval, nil

}
