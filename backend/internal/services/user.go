package services

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

type userService struct {
	userRepositories domains.IUserRepository
	utils            IUtilService
}

func newUserService(
	userRepositories domains.IUserRepository,
	utils IUtilService,
) domains.IUserService {
	return &userService{
		userRepositories: userRepositories,
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
