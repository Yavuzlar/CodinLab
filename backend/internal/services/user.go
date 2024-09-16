package services

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	hasher_service "github.com/Yavuzlar/CodinLab/pkg/hasher"
	"github.com/google/uuid"
)

type userService struct {
	userRepositories domains.IUserRepository
	logService       domains.ILogService
	utils            IUtilService
	parserService    domains.IParserService
}

func newUserService(
	userRepositories domains.IUserRepository,
	logService domains.ILogService,
	parserService domains.IParserService,
	utils IUtilService,
) domains.IUserService {
	return &userService{
		userRepositories: userRepositories,
		logService:       logService,
		parserService:    parserService,
		utils:            utils,
	}
}

func (s *userService) Login(ctx context.Context, username, password string) (user *domains.User, err error) {
	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{
		Username: username,
	}, 1, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return nil, service_errors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	user = &users[0]
	ok, err := hasher_service.CompareHashAndPassword(user.Password(), password)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while comparing password", err)
	}
	if !ok {
		return nil, service_errors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, username, name, surname, password, githubProfile string) (err error) {
	// Checking the username is already being used
	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{Username: username}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) != 0 {
		return service_errors.NewServiceErrorWithMessageAndError(400, "username already being used", err)
	}

	// Creating New User Model
	newUser, err := domains.NewUser(username, password, name, surname, "", githubProfile, 0)
	if err != nil {
		return err
	}

	// We save the new user to the database
	if err = s.userRepositories.Add(ctx, newUser); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the user", err)
	}

	return
}

func (s *userService) CreateUser(ctx context.Context, username, name, surname, password, githubProfile string) (err error) {
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

func (s *userService) GetAllUsers(ctx context.Context) (users []domains.User, err error) {
	// Retrieves all users whose role is 'user' from the database
	users, _, err = s.userRepositories.Filter(ctx, domains.UserFilter{
		Role: "user",
	}, 1000000, 1)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}

	return
}

func (s *userService) GetProfile(ctx context.Context, userID string) (user *domains.User, err error) {
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

func (s *userService) UpdateUser(ctx context.Context, userID, password, username, githubProfile, name, surname string) (err error) {
	user, err := s.GetProfile(ctx, userID)
	if err != nil {
		return err
	}

	if err := s.checkPassword(user.Password(), password); err != nil {
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

	user.SetGithubProfile(githubProfile)

	//checking if name is updated
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

func (s *userService) UpdatePassword(ctx context.Context, userID, password, newPassword, confirmPassword string) (err error) {
	user, err := s.GetProfile(ctx, userID)
	if err != nil {
		return err
	}

	if err := s.checkPassword(user.Password(), password); err != nil {
		return err
	}

	// Checking if password is being updated & password match with confirm password
	if newPassword != "" {
		if newPassword != confirmPassword {
			return service_errors.NewServiceErrorWithMessage(400, "password do not match")
		}
		if err := user.SetPassword(newPassword); err != nil {
			return err
		}
	}

	if err = s.userRepositories.Update(ctx, user); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while updating user", err)
	}
	return

}

func (s *userService) checkPassword(userPassword, password string) (err error) {
	// Checking if password matches
	ok, err := hasher_service.CompareHashAndPassword(userPassword, password)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while comparing password", err)
	}
	if !ok {
		return service_errors.NewServiceErrorWithMessage(400, "wrong password")
	}
	return
}

func (s *userService) DeleteUser(ctx context.Context, userID string) (err error) {
	userIDU, err := uuid.Parse(userID)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(400, "invalid user id", err)
	}

	users, _, err := s.userRepositories.Filter(ctx, domains.UserFilter{ID: userIDU}, 1, 1)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return service_errors.NewServiceErrorWithMessage(400, "user not found")
	}

	if err = s.userRepositories.Delete(ctx, userIDU); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while deleting the user", err)
	}
	return
}

// Find users most used programming languages
func (s *userService) BestProgrammingLanguages(ctx context.Context, userID string) (bestProgrammingLanguage string, err error) {
	programmingLanguageCount := make(map[int32]int)
	if logs, err := s.logService.GetByUserID(ctx, userID); err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while getting logs", err)
	} else {
		for _, log := range logs {
			programmingLanguageCount[log.ProgrammingID()]++
		}
	}
	max := 0
	var mostUsedProgrammingLanguageID int32
	for lang, count := range programmingLanguageCount {
		if count > max {
			max = count
			mostUsedProgrammingLanguageID = lang
		}
	}
	languages, err := s.parserService.GetInventory()
	if err != nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "error while getting languages", err)
	}
	if languages == nil {
		return "", service_errors.NewServiceErrorWithMessageAndError(500, "languages list is nil", err)
	}
	for _, programming := range languages {
		if programming.ID == int(mostUsedProgrammingLanguageID) {
			bestProgrammingLanguage = programming.Name
		}
	}
	return
}
