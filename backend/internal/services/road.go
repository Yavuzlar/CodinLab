package services

import (
	"context"
	"fmt"

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

func (s *roadService) getAllRoads(userID string) ([]domains.Road, error) {
	src, err := s.parserService.GetRoads()
	if err != nil {
		return nil, err
	}

	var roads []domains.Road
	for _, roadCollection := range src {
		var newPathList []domains.Path

		for _, path := range roadCollection.Paths {

			var languages []domains.LanguageRoad
			for _, lang := range path.Languages {
				languages = append(languages, *domains.NewLanguageRoad(lang.Lang, lang.Title, lang.Description, lang.Content, lang.Note))
			}

			var tests []domains.Test
			for _, test := range path.Quest.Tests {
				tests = append(tests, *domains.NewTest(test.Input, test.Output))
			}

			var params []domains.Param
			for _, param := range path.Quest.Params {
				params = append(params, *domains.NewParam(param.Name, param.Type))
			}

			var returns []domains.Returns
			for _, returnedParam := range path.Quest.Returns {
				returns = append(returns, *domains.NewReturn(returnedParam.Name, returnedParam.Type))
			}

			var questImports []string
			for _, questImport := range path.Quest.QuestImports {
				questImports = append(questImports, questImport)
			}

			var codeTemplates []domains.CodeTemplate
			for _, codeTemplateParser := range path.Quest.CodeTemplates {
				codeTemplates = append(codeTemplates, *domains.NewCodeTemplate(codeTemplateParser.ProgrammingID, codeTemplateParser.Frontend, codeTemplateParser.Template, codeTemplateParser.Check))
			}

			quest := domains.NewQuest(path.Quest.Difficulty, path.Quest.FuncName, tests, params, returns, questImports, codeTemplates)
			newPath := domains.NewPath(path.ID, languages, *quest, false, false)

			isStarted, isFinished, err := s.getPathStatuses(userID, fmt.Sprintf("%v", roadCollection.ID), fmt.Sprintf("%v", path.ID))
			if err != nil {
				return nil, err
			}

			newPath.SetIsStarted(*isStarted)
			newPath.SetIsFinished(*isFinished)

			newPathList = append(newPathList, *newPath)
		}

		isStarted, isFinished, err := s.getRoadStatuses(userID, fmt.Sprintf("%v", roadCollection.ID))
		if err != nil {
			return nil, err
		}

		roads = append(roads, *domains.NewRoads(roadCollection.ID, roadCollection.Name, roadCollection.DockerImage, roadCollection.IconPath, roadCollection.FileExtension, roadCollection.TemplatePath, newPathList, roadCollection.Cmd, *isStarted, *isFinished))
	}

	return roads, nil
}

func (s *roadService) getRoadStatuses(userID, programmingID string) (*bool, *bool, error) {
	var isStarted, isFinished bool
	logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, "", domains.TypeRoad, domains.ContentStarted)
	if err != nil {
		return nil, nil, err
	}
	if len(logStartedStatus) > 0 {
		isStarted = true
	}

	logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, "", domains.TypeRoad, domains.ContentCompleted)
	if err != nil {
		return nil, nil, err
	}
	if len(logFinishedStatus) > 0 {
		isFinished = true
	}

	return &isStarted, &isFinished, err
}

func (s *roadService) getPathStatuses(userID, programmingID, pathID string) (*bool, *bool, error) {
	var isStarted, isFinished bool
	logStartedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, pathID, domains.TypePath, domains.ContentStarted)
	if err != nil {
		return nil, nil, err
	}
	if len(logStartedStatus) > 0 {
		isStarted = true
	}

	logFinishedStatus, err := s.logService.GetAllLogs(context.TODO(), userID, programmingID, pathID, domains.TypePath, domains.ContentCompleted)
	if err != nil {
		return nil, nil, err
	}
	if len(logFinishedStatus) > 0 {
		isFinished = true
	}

	return &isStarted, &isFinished, err
}

func (s *roadService) GetRoadFilter(userID string, programmingID, pathId int, isStarted, isFinished *bool) ([]domains.Road, error) {
	allRoads, err := s.getAllRoads(userID)

	if err != nil {
		return nil, err
	}

	if userID == "" && programmingID == 0 && pathId == 0 && isStarted == nil && isFinished == nil {
		return allRoads, nil
	}

	var filteredRoads []domains.Road
	for _, roadCollection := range allRoads {

		if roadCollection.GetID() != programmingID {
			continue
		}

		var newRoadList []domains.Path
		for _, path := range roadCollection.GetPaths() {
			if pathId != 0 && path.GetID() != pathId {
				continue
			}

			if isStarted != nil && path.GetIsStarted() != *isStarted {
				continue
			}

			if isFinished != nil && path.GetIsFinished() != *isFinished {
				continue
			}

			newRoadList = append(newRoadList, path)
		}

		isStarted, isFinished, err := s.getRoadStatuses(userID, fmt.Sprintf("%v", roadCollection.GetID()))
		if err != nil {
			return nil, err
		}

		if len(newRoadList) > 0 {
			filteredRoads = append(filteredRoads, *domains.NewRoads(roadCollection.GetID(), roadCollection.GetName(), roadCollection.GetDockerImage(), roadCollection.GetIconPath(), roadCollection.GetFileExtension(), roadCollection.GetTemplatePath(), newRoadList, roadCollection.GetCmd(), *isStarted, *isFinished))
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

func (s *roadService) GetRoadByID(userID string, programmingID, pathID int) (path *domains.Path, err error) {
	road, err := s.getAllRoads(userID)
	if err != nil {
		return nil, err
	}
	for _, roadCollection := range road {
		if roadCollection.GetID() == int(programmingID) {
			for _, path := range roadCollection.GetPaths() {
				if path.GetID() == pathID {
					return &path, nil
				}
			}
		}
	}

	return nil, err
}
