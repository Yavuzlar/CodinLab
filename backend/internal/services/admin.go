package services

//author: yasir

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/google/uuid"
)

type adminService struct {
	userRepositories domains.IUserRepository
	logService       domains.ILogService
	utils            IUtilService
	parserService    domains.IParserService
}

func newAdminService(
	userRepositories domains.IUserRepository,
	logService domains.ILogService,
	parserService domains.IParserService,
	utils IUtilService,
) domains.IAdminService {
	return &adminService{
		userRepositories: userRepositories,
		logService:       logService,
		parserService:    parserService,
		utils:            utils,
	}
}

func (s *adminService) CreateUser(ctx context.Context, username, name, surname, password, role, githubProfile string) (err error) {
	if role != "admin" && role != "user"{
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

func (s *adminService) GetAllUsers(ctx context.Context) (users []domains.User, err error) {
	// Retrieves all users whose role is 'user' from the database
	users, _, err = s.userRepositories.Filter(ctx, domains.UserFilter{
		Role: "user",
	}, 1000000, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}

	return
}

func (s *adminService) GetProfile(ctx context.Context, userID string) (user *domains.User, err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	//Checking if the user exists and retrieving user
	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{
		Id: userIDU,
	}, 1, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "invalid request")
	}
	user = &users[0]

	return
}

func (s *adminService) UpdateUser(ctx context.Context, userID, password, newPassword, username, githubProfile, name, surname string) (err error) {
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
			oldUsername := filter[0].Username()
			if oldUsername != username {
				return service_errors.NewServiceErrorWithMessageAndError(400, "username already being used", err)
			}
		}
		if err := user.SetUsername(username); err != nil {
			return err
		}
	}

	// Checking if password is being updated
	if newPassword != "" {
		if err := user.SetPassword(newPassword); err != nil {
			return err
		}
	}

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

	if err = s.userRepositories.Update(ctx, user); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while updating user", err)
	}

	return
}

func (s *adminService) DeleteUser(ctx context.Context, userID string) (err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{Id: userIDU}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return service_errors.NewServiceErrorWithMessage(400, "invalid request")
	}

	if err = s.userRepositories.Delete(ctx, userIDU); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while deleting the user", err)
	}
	return
}

// Find users most used languages
func (s *adminService) BestLanguage(ctx context.Context, userID string) (bestLanguage string, err error) {
	languageCount := make(map[int32]int)
	if s.logService == nil || s.parserService == nil {
		return "", service_errors.NewServiceErrorWithMessage(500, "service is not initialized")
	}
	if logs, err := s.logService.GetByUserID(ctx, userID); err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while getting logs", err)
	} else {
		for _, log := range logs {
			languageCount[log.LanguageID()]++
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
