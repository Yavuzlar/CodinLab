package services

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

type userService struct {
	userRepositories domains.IUserRepository
	logService       domains.ILogService
	utils            IUtilService
}

func newUserService(
	userRepositories domains.IUserRepository,
	logService domains.ILogService,
	utils IUtilService,
) domains.IUserService {
	return &userService{
		userRepositories: userRepositories,
		logService:       logService,
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
	ok, err := s.utils.Hasher().CompareHashAndPassword(user.Password(), password)
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

	// Hashing password for db
	hashedPassword, err := s.utils.Hasher().HashPassword(password)
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while hashing the password", err)
	}

	// Creating New User Model
	newUser, err := domains.NewUser(username, hashedPassword, name, surname, "", githubProfile)
	if err != nil {
		return err
	}

	// We save the new user to the database
	if err = s.userRepositories.Add(ctx, newUser); err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while adding the user", err)
	}

	return
}
