package services

//author: yasir

import (
	"context"
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

type adminService struct {
	userRepositories domains.IUserRepository
	logService       domains.ILogService
	utils            IUtilService
	parserService    domains.IParserService
	levelService     domains.ILevelService
}

func newAdminService(
	userRepositories domains.IUserRepository,
	logService domains.ILogService,
	parserService domains.IParserService,
	levelService domains.ILevelService,
	utils IUtilService,
) domains.IAdminService {
	return &adminService{
		userRepositories: userRepositories,
		logService:       logService,
		parserService:    parserService,
		levelService:     levelService,
		utils:            utils,
	}
}

func (s *adminService) CreateUser(ctx context.Context, username, name, surname, password, role, githubProfile string) (err error) {
	if role != "admin" && role != "user" {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid role", err)
	}

	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{
		Username: username,
	}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) != 0 {
		return service_errors.NewServiceErrorWithMessageAndError(400, "username already being used", err)
	}

	newUser, err := domains.NewUser(username, password, name, surname, "", githubProfile, 0)
	if err != nil {
		return err
	}

	if err = s.userRepositories.Add(ctx, newUser); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the user", err)
	}

	return
}

func (s *adminService) GetAllUsers(ctx context.Context) ([]domains.AdminUserDetail, error) {
	var adminUserDetail []domains.AdminUserDetail
	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{
		Role: "user",
	}, 1000000, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return nil, service_errors.NewServiceErrorWithMessageAndError(404, "user list is empty", err)
	}

	for i, user := range users {
		// Getting most used programming languaage
		mostUsedProgrammingLanguage, err := s.BestProgrammingLanguage(ctx, user.ID().String())
		if err != nil {
			return nil, err
		}

		// Gettins user level
		userLevel, err := s.levelService.GetUserLevel(ctx, user.ID().String())
		if err != nil {
			return nil, err
		}
		adminUserDetail = append(adminUserDetail, *domains.NewAdminUser(user.ID(), i+1, user.Username(), strconv.Itoa(userLevel.Level())+" Level", mostUsedProgrammingLanguage))
	}

	return adminUserDetail, nil
}

func (s *adminService) GetProfile(ctx context.Context, userID string) (user *domains.User, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	//Checking if the user exists and retrieving user
	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{
		ID: userIDU,
	}, 1, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "user not found")
	}
	user = &users[0]

	return
}

func (s *adminService) UpdateUser(ctx context.Context, userID, role, username, githubProfile, name, surname string) (err error) {
	user, err := s.GetProfile(ctx, userID)
	if err != nil {
		return err
	}

	// Checking if username is being updated
	if username != "" {
		//Checking the username is already being used
		filter, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{Username: username}, 1, 1)
		if err != nil {
			return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
		}
		if len(filter) > 0 {
			oldUsername := user.Username()
			if oldUsername != filter[0].Username() {
				return service_errors.NewServiceErrorWithMessageAndError(400, "username already being used", err)
			}
		}
		if err := user.SetUsername(username); err != nil {
			return err
		}
	}

	user.SetRole(role)
	user.SetGithubProfile(githubProfile)
	if name != "" {
		if err := user.SetName(name); err != nil {
			return err
		}
	}
	if surname != "" {
		if err := user.SetSurname(surname); err != nil {
			return err
		}
	}

	if err = s.userRepositories.AdminUpdate(ctx, user); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while updating user", err)
	}

	return
}

func (s *adminService) DeleteUser(ctx context.Context, userID string) (err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{ID: userIDU}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return service_errors.NewServiceErrorWithMessage(404, "user not found")
	}

	isAdmin, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{ID: userIDU, Role: "user"}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(isAdmin) == 0 {
		return service_errors.NewServiceErrorWithMessage(403, "no permission to delete")
	}

	if err = s.userRepositories.Delete(ctx, userIDU); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while deleting the user", err)
	}

	return
}

// Find users most used languages
func (s *adminService) BestProgrammingLanguage(ctx context.Context, userID string) (bestLanguage string, err error) {
	languageCount := make(map[int32]int)
	if s.logService == nil || s.parserService == nil {
		return "", service_errors.NewServiceErrorWithMessage(500, "service is not initialized")
	}
	if logs, err := s.logService.GetByUserID(ctx, userID); err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while getting logs", err)
	} else {
		for _, log := range logs {
			languageCount[log.ProgrammingID()]++
		}
	}
	max := 0
	var mostUsedLanguageID int32
	for lang, count := range languageCount {
		if count > max {
			max = count
			mostUsedLanguageID = lang
		}
	}
	languages, err := s.parserService.GetInventory()
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while getting languages", err)
	}
	if languages == nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "languages list is nil", err)
	}
	for _, lang := range languages {
		if lang.ID == int(mostUsedLanguageID) {
			bestLanguage = lang.Name
		}
	}
	return
}
