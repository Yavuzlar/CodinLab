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
func (s *labService) GetAllLabs(userID string) ([]domains.Labs, error) {
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
				languages = append(languages, domains.Language{
					Lang:        lang.Lang,
					Title:       lang.Title,
					Description: lang.Description,
					Note:        lang.Note,
					Hint:        lang.Hint,
					//Language parametreleri buraya eklenecek
				})
			}

			var tests []domains.Test
			for _, test := range lab.Quest.Tests {
				tests = append(tests, domains.Test{
					Input:  test.Input,
					Output: test.Output,
					//Test parametreleri buraya eklenecek
				})
			}

			var params []domains.Param
			for _, param := range lab.Quest.Params {
				params = append(params, domains.Param{
					Name: param.Name,
					Type: param.Type,
					//Param parametreleri buraya eklenecek
				})
			}

			quest := domains.Quest{
				Difficulty: lab.Quest.Difficulty,
				FuncName:   lab.Quest.FuncName,
				Tests:      tests,
				Params:     params,
				//Quest parametreleri buraya eklenecek
			}

			newLab := domains.Lab{
				ID:         lab.ID,
				Languages:  languages,
				Quest:      quest,
				IsStarted:  "false", // Varsayılan değer false
				IsFinished: "false", // Varsayılan değer false
				//Lab parametreleri buraya eklenecek
			}
			labIDString := strconv.Itoa(lab.ID)

			logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", labIDString, domains.TypeLab, domains.ContentStarted)
			if err != nil {
				return nil, err
			}

			if len(logStartedStatus) > 0 {
				newLab.IsStarted = "true"
			}

			logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", labIDString, domains.TypeLab, domains.ContentCompleted)
			if err != nil {
				return nil, err
			}

			if len(logFinishedStatus) > 0 {
				newLab.IsFinished = "true"
			}

			newLabList = append(newLabList, newLab)
		}

		labs = append(labs, domains.Labs{
			ID:          labCollection.ID,
			Name:        labCollection.Name,
			DockerImage: labCollection.DockerImage,
			IconPath:    labCollection.IconPath,
			Labs:        newLabList,
		})
	}

	return labs, nil
}

// Fetch labs by filters
func (s *labService) GetLabsFilter(userID string, labsId, labId int, isStarted, isFinished string) ([]domains.Labs, error) {
	allLabs, err := s.GetAllLabs(userID)
	if err != nil {
		return nil, err
	}

	var filteredLabs []domains.Labs

	for _, labCollection := range allLabs {

		if labsId != 0 && labCollection.ID != labsId {
			continue
		}
		//labs structı için filtreleme eklenebilir.

		var newLabList []domains.Lab
		for _, lab := range labCollection.Labs {

			if labId != 0 && lab.ID != labId {
				continue
			}
			if isStarted != "" && lab.IsStarted != isStarted {
				continue
			}
			if isFinished != "" && lab.IsFinished != isFinished {
				continue
			}
			//lab structı için filtreleme eklenebilir.

			newLabList = append(newLabList, lab)
		}

		if len(newLabList) > 0 {
			filteredLabs = append(filteredLabs, domains.Labs{
				ID:          labCollection.ID,
				Name:        labCollection.Name,
				DockerImage: labCollection.DockerImage,
				IconPath:    labCollection.IconPath,
				Labs:        newLabList,
			})
		}
	}

	if len(filteredLabs) == 0 && (userID == "" && labsId == 0 && labId == 0 && isStarted == "" && isFinished == "") {
		return allLabs, nil
	}

	return filteredLabs, nil
}
