package services

import "github.com/Yavuzlar/CodinLab/internal/domains"

// Tüm servisler tek bir yapıda toplayarark tek endpoint üzerinden erişim sağlamak için oluşturulmuştur.

type Services struct {
	UtilService   IUtilService
	UserService   domains.IUserService
	LogService    domains.ILogService
	DockerService domains.IDockerService
	ParserService domains.IParserService
	// diğer servisler buraya eklenecek
}

func CreateNewServices(
	userRepositories domains.IUserRepository,
	logRepositories domains.ILogRepository,
	hasherService IHashService,
	validatorService IValidatorService,
	// diğer servisler buraya eklenecek

) *Services {
	utilsService := newUtilService(hasherService, validatorService)
	logService := newLogService(logRepositories, utilsService)
	userService := newUserService(userRepositories, logService, utilsService)
	dockerService := newDockerService(utilsService)
	parserService := newParserService(utilsService)
	// diğer servisler buraya eklenecek

	return &Services{
		UtilService:   utilsService,
		DockerService: dockerService,
		UserService:   userService,
		LogService:    logService,
		ParserService: parserService,
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

// ------------------ HASH SERVICE ------------------
// HashService interface'i password hashleme ve hash ile password karşılaştırma işlemlerini yapar.
type IHashService interface {
	HashPassword(password string) (hashedPassword string, err error)
	CompareHashAndPassword(hashedPassword string, password string) (ok bool, err error)
}

// ------------------ VALIDATOR SERVICE ------------------
// ValidatorService interface'i struct'ın içindeki alanların doğruluğunu kontrol eder.

type IValidatorService interface {
	ValidateStruct(s any) error
}

// ------------------ UTIL SERVICE ------------------
// UtilService struct'ı HashService ve ValidatorService'yi içerir.
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
