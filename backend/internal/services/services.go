package services

import "github.com/Yavuzlar/CodinLab/internal/domains"

// Tüm servisler tek bir yapıda toplayarark tek endpoint üzerinden erişim sağlamak için oluşturulmuştur.

type Services struct {
	UtilService     IUtilService
	UserService     domains.IUserService
	LogService      domains.ILogService
	ParserService   domains.IParserService
	LabService      domains.ILabService
	RoadService     domains.IRoadService
	LevelService    domains.ILevelService
	HomeService     domains.IHomeService
	AdminService    domains.IAdminService
	CodeService     domains.ICodeService
	StartService    domains.IStartService
	TemplateService domains.ITemplateService
	LabRoadService  domains.ILabRoadService
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
	levelService := newLevelService(utilsService, logService, parserService, userRepositories)
	userService := newUserService(userRepositories, logService, parserService, utilsService)
	labService := newLabService(utilsService, logService, parserService)
	roadService := newRoadService(utilsService, logService, parserService)
	homeService := newHomeService(utilsService, logService, parserService, levelService)
	adminService := newAdminService(userRepositories, logService, parserService, levelService, utilsService)
	startService := newStartService(utilsService, parserService)
	codeService := NewCodeService()
	templateService := newTemplateService(utilsService, logService, parserService, labService, roadService)
	labRoadService := newLabRoadService(parserService)
	// diğer servisler buraya eklenecek

	return &Services{
		UtilService:     utilsService,
		UserService:     userService,
		LogService:      logService,
		ParserService:   parserService,
		LabService:      labService,
		RoadService:     roadService,
		LevelService:    levelService,
		HomeService:     homeService,
		AdminService:    adminService,
		CodeService:     codeService,
		StartService:    startService,
		TemplateService: templateService,
		LabRoadService:  labRoadService,
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
