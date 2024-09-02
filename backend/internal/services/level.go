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

func (s *levelService) updateUserPoint(ctx context.Context, userID string) (oldPoint int32, newPoint int32, err error) {

	logs, err := s.logService.GetByUserID(ctx, userID)
	if err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(500, "error while getting logs", err)
	}
	labs, err := s.parserService.GetLabs()
	if err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(500, "error while getting labs", err)
	}
	roads, err := s.parserService.GetRoads()
	if err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(500, "error while getting roads", err)
	}

	for _, log := range logs { //calculates userpoints via completed labs and paths
		if log.Content() == domains.ContentCompleted {
			if log.Type() == domains.TypeLab {
				for _, lab := range labs {
					if int32(lab.ID) == log.ProgrammingID() {
						for _, l := range lab.Labs {
							if int32(l.ID) == log.LabPathID() {
								newPoint += int32(l.Quest.Difficulty)
							}
						}
					}
				}
			} else if log.Type() == domains.TypePath {
				for _, road := range roads {
					if int32(road.ID) == log.ProgrammingID() {
						for _, path := range road.Paths {
							if int32(path.ID) == log.LabPathID() {
								newPoint += int32(path.Quest.Difficulty)
							}
						}
					}
				}
			}
		}
	}

	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	user, _, err := s.userRepository.Filter(ctx, domains.UserFilter{ID: userIDU}, 1, 1)
	if err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	if len(user) == 0 {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	u := &user[0]
	oldPoint = u.TotalPoints()
	if u.TotalPoints() == newPoint {
		return
	}

	u.SetTotalPoints(newPoint)

	if err = s.userRepository.Update(ctx, u); err != nil {
		return 0, 0, service_errors.NewServiceErrorWithMessageAndError(500, "error while updating user points", err)
	}

	return
}

// Finds user level
func (s *levelService) GetUserLevel(ctx context.Context, userID string) (userLevel *domains.UserLevel, err error) {
	levels, err := s.parserService.GetLevels()
	var language []domains.LanguageLevel
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while getting levels", err)
	}
	oldPoint, newPoint, err := s.updateUserPoint(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, level := range levels {
		if oldPoint >= level.MinPoints && oldPoint < level.MaxPoints { //Checks level limits
			levelPercentage := ((oldPoint - level.MinPoints) * 100) / (level.MaxPoints - level.MinPoints) //level percentage
			for _, lang := range level.Languages {
				language = append(language, domains.NewLanguageLevel(lang.Lang, lang.Description))
			}
			userLevel = domains.NewUserLevel(level.Level, oldPoint, levelPercentage, language)
			break
		}
	}
	for _, level := range levels {
		if newPoint >= level.MinPoints && newPoint < level.MaxPoints {
			if userLevel.Level() != level.Level { //Compares the old level with the new level
				userLevel.SetLevel(level.Level)
				language = nil
				for _, lang := range level.Languages {
					language = append(language, domains.NewLanguageLevel(lang.Lang, lang.Description))
				}
				userLevel.SetLanguages(language)
				if err = s.logService.Add(ctx, userID, domains.TypeUser, domains.ContentLevelUp, 0, 0); err != nil {
					return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while adding log", err)
				}
			}
			levelPercentage := ((newPoint - level.MinPoints) * 100) / (level.MaxPoints - level.MinPoints)
			userLevel.SetLevelPercentage(levelPercentage)
			userLevel.SetTotalPoints(newPoint)
		}
	}
	return
}
