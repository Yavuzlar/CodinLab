package services

import (
	"context"
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

type homeService struct {
	utils         IUtilService
	logService    domains.ILogService
	parserService domains.IParserService
	levelService  domains.ILevelService
}

func newHomeService(
	utils IUtilService,
	logService domains.ILogService,
	parserService domains.IParserService,
	levelService domains.ILevelService,
) domains.IHomeService {
	return &homeService{
		utils:         utils,
		logService:    logService,
		parserService: parserService,
		levelService:  levelService,
	}
}

func (s *homeService) GetUserLevel(ctx context.Context, userID string) (userLevel *domains.UserLevel, err error) {
	userLevel, err = s.levelService.GetUserLevel(ctx, userID)
	return
}

func (s *homeService) GetInventory(ctx context.Context) (inventory []domains.Inventory, err error) {
	inventoryP, err := s.parserService.GetInventory()
	if err != nil {
		return
	}

	for _, item := range inventoryP {
		newInventory := *domains.NewInventory(item.ID, item.Name, item.IconPath)
		inventory = append(inventory, newInventory)
	}

	if len(inventory) == 0 {
		return nil, service_errors.NewServiceErrorWithMessage(200, "Programming Languages not found")
	}

	return
}

func (s *homeService) GetUserDevelopment(ctx context.Context, userID string) (development domains.Development, err error) {
	labCompleted, err := s.logService.GetAllLogs(ctx, userID, "", "", domains.TypeLab, domains.ContentCompleted)
	if err != nil {
		return
	}
	labCompletedCount := len(labCompleted)

	roadCompleted, err := s.logService.GetAllLogs(ctx, userID, "", "", domains.TypePath, domains.ContentCompleted)
	if err != nil {
		return
	}
	roadCompletedCount := len(roadCompleted)
	var labPercentage, roadPercentage int32

	if labCompletedCount > 0 {
		programmingLanguages, _ := s.parserService.GetInventory()
		allLabs, err := s.parserService.GetLabs()
		if err != nil {
			return development, err
		}
		allLabsCount := len(allLabs) * len(programmingLanguages)
		if allLabsCount > 0 {
			if labCompletedCount > allLabsCount {
				labCompletedCount = allLabsCount
			}
			labPercentage = int32((float32(labCompletedCount) / float32(allLabsCount)) * 100)
		}
	}

	if roadCompletedCount > 0 {
		allRoads, err := s.parserService.GetRoads()
		if err != nil {
			return development, err
		}
		allRoadsCount := s.countAllRoad(allRoads)
		if allRoadsCount > 0 {
			if roadCompletedCount > allRoadsCount {
				roadCompletedCount = allRoadsCount
			}
			roadPercentage = int32((float32(roadCompletedCount) / float32(allRoadsCount)) * 100)
		}
	}

	development = *domains.NewUserDevelopment(roadPercentage, labPercentage)
	return
}

func (s *homeService) GetUserAdvancement(ctx context.Context, userID string) (advancement []domains.Advancement, err error) {
	inventoryP, err := s.parserService.GetInventory()
	if err != nil {
		return
	}

	allLabs, err := s.parserService.GetLabs()
	if err != nil {
		return
	}

	allRoads, err := s.parserService.GetRoads()
	if err != nil {
		return
	}

	for _, item := range inventoryP {
		var labPercentage, roadPercentage int32

		labCompleted, err := s.logService.GetAllLogs(ctx, userID, strconv.Itoa(item.ID), "", domains.TypeLab, domains.ContentCompleted)
		if err != nil {
			return nil, err
		}

		labCompletedCount := len(labCompleted)
		if labCompletedCount > 0 {
			labByIdCount := len(allLabs) * len(inventoryP)
			if labByIdCount > 0 {
				if labCompletedCount > labByIdCount {
					labCompletedCount = labByIdCount
				}
				labPercentage = int32((float32(labCompletedCount) / float32(labByIdCount)) * 100)
			}

		}

		roadCompleted, err := s.logService.GetAllLogs(ctx, userID, strconv.Itoa(item.ID), "", domains.TypePath, domains.ContentCompleted)
		if err != nil {
			return nil, err
		}
		roadCompletedCount := len(roadCompleted)
		if roadCompletedCount > 0 {
			roadByIdCount := s.countRoadById(allRoads, item.ID)
			if roadByIdCount > 0 {
				if roadCompletedCount > roadByIdCount {
					roadCompletedCount = roadByIdCount
				}
				roadPercentage = int32((float32(labCompletedCount) / float32(roadByIdCount)) * 100)
			}
		}

		newAdvancement := *domains.NewAdvancement(item.ID, item.Name, item.IconPath, roadPercentage, labPercentage)
		advancement = append(advancement, newAdvancement)
	}
	return
}

func (s *homeService) GetWelcomeContent() (content []domains.WelcomeContent, err error) {
	content, err = s.parserService.GetWelcomeBanner()
	return
}

func (s *homeService) GetRoadContent() (content []domains.RoadContent, err error) {
	content, err = s.parserService.GetRoadBanner()
	return
}

func (s *homeService) GetLabContent() (content []domains.LabContent, err error) {
	content, err = s.parserService.GetLabBanner()
	return
}

func (s *homeService) countRoadById(roadsP []domains.RoadP, id int) int {
	count := 0

	for _, roadGroup := range roadsP {
		if roadGroup.ID == id {
			count += len(roadGroup.Paths)
		}
	}

	return count
}

func (s *homeService) countAllRoad(roadsP []domains.RoadP) int {
	count := 0
	for _, roadGroup := range roadsP {
		count += len(roadGroup.Paths)
	}
	return count
}
