package services

import "github.com/Yavuzlar/CodinLab/internal/domains"

// Tüm servisler tek bir yapıda toplayarark tek endpoint üzerinden erişim sağlamak için oluşturulmuştur.

type Services struct {
	UtilService   IUtilService
	UserService   domains.IUserService
	LogService    domains.ILogService
	DockerService domains.IDockerService
	ParserService domains.IParserService
	LabService    domains.ILabService
	// diğer servisler buraya eklenecek
}

func CreateNewServices(
	userRepositories domains.IUserRepository,
	logRepositories domains.ILogRepository,
	validatorService IValidatorService,
	// diğer servisler buraya eklenecek

) *Services {
	utilsService := newUtilService(validatorService)
	logService := newLogService(logRepositories, utilsService)
	parserService := newParserService(utilsService)
	userService := newUserService(userRepositories, logService, parserService, utilsService)
	dockerService := newDockerService(utilsService)
	labService := newLabService(utilsService, logService, parserService)
	// diğer servisler buraya eklenecek

	return &Services{
		UtilService:   utilsService,
		DockerService: dockerService,
		UserService:   userService,
		LogService:    logService,
		ParserService: parserService,
		LabService:    labService,
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
	Validator() IValidatorService
}

// ------------------ VALIDATOR SERVICE ------------------
// ValidatorService interface'i struct'ın içindeki alanların doğruluğunu kontrol eder.

type IValidatorService interface {
	ValidateStruct(s any) error
}

// ------------------ UTIL SERVICE ------------------
// UtilService struct'ı ValidatorService içerir.
type utilService struct {
	validatorService IValidatorService
}

func newUtilService(
	validatorService IValidatorService,
) IUtilService {
	return &utilService{
		validatorService: validatorService,
	}
}

func (s *utilService) Validator() IValidatorService {
	return s.validatorService
}
