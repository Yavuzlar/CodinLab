package services

import (
	"context"
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
)

type roadService struct {
	utils         IUtilService
	logService    domains.ILogService
	parserService domains.IParserService
}

func newRoadService(
	utils IUtilService,
	logService domains.ILogService,
	parserService domains.IParserService,
) domains.IRoadService {
	return &roadService{
		utils:         utils,
		logService:    logService,
		parserService: parserService,
	}
}

func (s *roadService) getAllRoads(userID string) ([]domains.Roads, error) {
	src, err := s.parserService.GetRoads()
	if err != nil {
		return nil, err
	}

	var roads []domains.Roads

	for _, roadCollection := range src {
		var newPathList []domains.Path

		for _, path := range roadCollection.Paths {

			var languages []domains.LanguageRoad
			for _, lang := range path.Languages {
				languages = append(languages, *domains.NewLanguageRoad(lang.Lang, lang.Title, lang.Description, lang.Content, lang.Note))
			}

			var tests []domains.TestRoad
			for _, test := range path.Quest.Tests {
				tests = append(tests, *domains.NewTestRoad(test.Input, test.Output))
			}

			var params []domains.ParamRoad
			for _, param := range path.Quest.Params {
				params = append(params, *domains.NewParamRoad(param.Name, param.Type))
			}
			quest := domains.NewQuestRoad(path.Quest.Difficulty, path.Quest.FuncName, tests, params)
			newPath := domains.NewPath(path.ID, languages, *quest, false, false)

			pathIDString := strconv.Itoa(path.ID)
			programmingLanguageIDString := strconv.Itoa(roadCollection.ID)
			logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingLanguageIDString, pathIDString, domains.TypePath, domains.ContentStarted)
			if err != nil {
				return nil, err
			}
			if len(logStartedStatus) > 0 {
				newPath.SetIsStarted(true)
			}

			logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingLanguageIDString, pathIDString, domains.TypePath, domains.ContentCompleted)
			if err != nil {
				return nil, err
			}
			if len(logFinishedStatus) > 0 {
				newPath.SetIsFinished(true)
			}

			newPathList = append(newPathList, *newPath)
		}
		roads = append(roads, *domains.NewRoads(roadCollection.ID, roadCollection.Name, roadCollection.DockerImage, roadCollection.IconPath, newPathList))
	}

	return roads, nil
}

func (s *roadService) GetRoadFilter(userID string, programmingID, pathId int, isStarted, isFinished *bool) ([]domains.Roads, error) {
	allRoads, err := s.getAllRoads(userID)

	if err != nil {
		return nil, err
	}

	if userID == "" && programmingID == 0 && pathId == 0 && isStarted == nil && isFinished == nil {
		return allRoads, nil
	}

	var filteredRoads []domains.Roads
	for _, roadCollection := range allRoads {

		if roadCollection.GetID() != programmingID {
			continue
		}

		var newRoadList []domains.Path

		for _, road := range roadCollection.GetPaths() {
			if pathId != 0 && road.GetID() != pathId {
				continue
			}

			if isStarted != nil && road.GetIsStarted() != *isStarted {
				continue
			}

			if isFinished != nil && road.GetIsFinished() != *isFinished {
				continue
			}

			newRoadList = append(newRoadList, road)
		}

		if len(newRoadList) > 0 {
			filteredRoads = append(filteredRoads, *domains.NewRoads(roadCollection.GetID(), roadCollection.GetName(), roadCollection.GetDockerImage(), roadCollection.GetIconPath(), newRoadList))
		}
	}

	return filteredRoads, nil

}

func (s *roadService) GetUserLanguageRoadStats(userID string) (programmingLangugageStats []domains.RoadStats, err error) {
	allRoads, err := s.getAllRoads(userID)
	if err != nil {
		return nil, err
	}
	for _, road := range allRoads {
		totalRoads := 0
		finishedRoads := 0
		for _, path := range road.GetPaths() {
			if path.GetIsFinished() {
				finishedRoads++
			}
			totalRoads++
		}
		newRoadStats := domains.NewRoadStats(
			road.GetID(),
			road.GetName(),
			road.GetIconPath(),
			totalRoads,
			finishedRoads,
			float32(finishedRoads)/float32(totalRoads)*100,
		)
		programmingLangugageStats = append(programmingLangugageStats, *newRoadStats)
	}

	return
}

func (s *roadService) GetUserRoadProgressStats(userID string) (progressStats *domains.RoadProgressStats, err error) {
	progress := 0
	completed := 0
	totalRoads := 0
	allRoads, err := s.getAllRoads(userID)
	if err != nil {
		return nil, err
	}
	for _, road := range allRoads {
		for _, path := range road.GetPaths() {
			if path.GetIsStarted() && !path.GetIsFinished() {
				progress++
			}
			if path.GetIsFinished() && path.GetIsStarted() {
				completed++
			}
			totalRoads++
		}
	}
	progressStats = domains.NewRoadProgressStats(
		float32(progress)/float32(totalRoads)*100,
		float32(completed)/float32(totalRoads)*100,
	)
	return
}
