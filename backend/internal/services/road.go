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

			var languages []domains.LanguageR
			for _, lang := range path.Languages {
				languages = append(languages, domains.LanguageR{
					Lang:        lang.Lang,
					Title:       lang.Title,
					Description: lang.Description,
					Content:     lang.Content,
					Note:        lang.Note,
				})
			}

			var tests []domains.TestR
			for _, test := range path.Quest.Tests {
				tests = append(tests, domains.TestR(
					test,
				))
			}

			var params []domains.ParamR
			for _, param := range path.Quest.Params {
				params = append(params, domains.ParamR(
					param,
				))
			}

			quest := domains.QuestR{
				Difficulty: path.Quest.Difficulty,
				FuncName:   path.Quest.FuncName,
				Tests:      tests,
				Params:     params,
			}

			newPath := domains.Path{
				ID:         path.ID,
				Languages:  languages,
				Quest:      quest,
				IsStarted:  false, // Default value is false
				IsFinished: false, // Default value is false
			}
			pathIDString := strconv.Itoa(path.ID)

			logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", pathIDString, domains.TypePath, domains.ContentStarted)
			if err != nil {
				return nil, err
			}

			if len(logStartedStatus) > 0 {
				newPath.IsStarted = true
			}

			logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, "", pathIDString, domains.TypePath, domains.ContentCompleted)
			if err != nil {
				return nil, err
			}

			if len(logFinishedStatus) > 0 {
				newPath.IsFinished = true
			}

			newPathList = append(newPathList, newPath)
		}

		roads = append(roads, domains.Roads{
			ID:          roadCollection.ID,
			Name:        roadCollection.Name,
			DockerImage: roadCollection.DockerImage,
			IconPath:    roadCollection.IconPath,
			Paths:       newPathList,
		})
	}

	return roads, nil
}

func (s *roadService) GetRoadFilter(userID string, roadId, pathId int, isStarted, isFinished bool) ([]domains.Roads, error) {
	allRoads, err := s.getAllRoads(userID)

	if err != nil {
		return nil, err
	}

	var filteredRoads []domains.Roads
	for _, roadCollection := range allRoads {

		if roadCollection.ID != roadId {
			continue
		}

		var newRoadList []domains.Path
		for _, road := range roadCollection.Paths {
			if pathId != 0 && road.ID != pathId {
				continue
			}
			if isStarted != false && road.IsStarted != isStarted {
				continue
			}

			if isFinished != false && road.IsFinished != isFinished {
				continue
			}

			newRoadList = append(newRoadList, road)
		}

		if len(newRoadList) > 0 {
			filteredRoads = append(filteredRoads, domains.Roads{
				ID:          roadCollection.ID,
				Name:        roadCollection.Name,
				DockerImage: roadCollection.DockerImage,
				IconPath:    roadCollection.IconPath,
				Paths:       newRoadList,
			})
		}
	}

	return filteredRoads, nil

}
