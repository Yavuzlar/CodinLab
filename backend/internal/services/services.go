package services

import "github.com/Yavuzlar/CodinLab/internal/domains"

// Tüm servisler tek bir yapıda toplayarark tek endpoint üzerinden erişim sağlamak için oluşturulmuştur.

type Services struct {
	UtilService IUtilService
	UserService domains.IUserService
	// diğer servisler buraya eklenecek
}

func CreateNewServices(
	userRepositories domains.IUserRepository,
	hasherService IHashService,
	validatorService IValidatorService,
	// diğer servisler buraya eklenecek

) *Services {
	utilsService := newUtilService(hasherService, validatorService)
	userService := newUserService(userRepositories, utilsService)
	// diğer servisler buraya eklenecek

	return &Services{
		UtilService: utilsService,
		UserService: userService,
		// diğer servisler buraya eklenecek

	}
}

func (s *Services) Util() IUtilService {
	return s.UtilService
}

func (s *Services) User() domains.IUserService {
	return s.UserService
}

// diğer Servisler buraya eklenecek

// ------------------ UTIL SERVICE ------------------
type IUtilService interface {
	Hasher() IHashService
	Validator() IValidatorService
}

type IHashService interface {
	HashPassword(password string) (hashedPassword string, err error)
	CompareHashAndPassword(hashedPassword string, password string) (ok bool, err error)
}

type IValidatorService interface {
	ValidateStruct(s any) error
}

type utilService struct {
	hasherService    IHashService
	validatorService IValidatorService
}

func newUtilService(
	hasherService IHashService,
	validatorService IValidatorService,
) IUtilService {
	return &utilService{
		hasherService:    hasherService,
		validatorService: validatorService,
	}
}

func (s *utilService) Hasher() IHashService {
	return s.hasherService
}

func (s *utilService) Validator() IValidatorService {
	return s.validatorService
}
