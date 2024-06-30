package services

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

type levelService struct {
	utils          IUtilService
	logService     domains.ILogService
	parserService  domains.IParserService
	userRepository domains.IUserRepository
}

func newLevelService(
	utils IUtilService,
	logService domains.ILogService,
	parserService domains.IParserService,
	userRepository domains.IUserRepository,
) domains.ILevelService {
	return &levelService{
		utils:          utils,
		logService:     logService,
		parserService:  parserService,
		userRepository: userRepository,
	}
}

// If a lab or path is solved, its difficulty is sent for user point update.
func (s *levelService) UpdateUserPoint(ctx context.Context, userID string, difficulty int32) error {
	userLevel, err := s.GetUserLevel(ctx, userID)
	if err != nil {
		return err
	}
	newPoint := userLevel.TotalPoints() + difficulty //Adds the new point to the old points
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}
	user, _, err := s.userRepository.Filter(ctx, domains.UserFilter{Id: userIDU}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}
	if len(user) == 0 {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	u := &user[0]
	u.SetTotalPoints(newPoint)
	if err = s.userRepository.Update(ctx, u); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while updating user points", err)
	}

	levels, err := s.parserService.GetLevels()
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while getting levels", err)
	}

	for _, level := range levels {
		if newPoint >= level.MinPoints && newPoint < level.MaxPoints {
			if userLevel.Level() != level.Level { //Compares the old level with the new level
				if err = s.logService.Add(ctx, userID, domains.TypeUser, domains.ContentLevelUp, 0, 0); err != nil {
					return service_errors.NewServiceErrorWithMessageAndError(500, "error while creating log", err)
				}
			}
		}
	}
	return nil
}

// Finds user level
func (s *levelService) GetUserLevel(ctx context.Context, userID string) (userLevel *domains.UserLevel, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}
	user, _, err := s.userRepository.Filter(ctx, domains.UserFilter{Id: userIDU}, 1, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}
	levels, err := s.parserService.GetLevels()
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while getting levels", err)
	}
	var userPoint int32
	if len(user) == 0 {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}
	u := user[0]
	userPoint = u.TotalPoints()
	var lang []domains.LanguageL

	for _, level := range levels {
		if userPoint >= level.MinPoints && userPoint < level.MaxPoints { //Checks level limits
			levelPercentage := ((userPoint - level.MinPoints) * 100) / (level.MaxPoints - level.MinPoints) //level percentage
			for _, language := range level.Languages {
				lang = append(lang, domains.LanguageL{Lang: language.Lang, Description: language.Description})
			}
			userLevel = domains.NewUserLevel(level.Level, userPoint, levelPercentage, lang)
			break
		}
	}
	return
}
