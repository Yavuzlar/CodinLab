package domains

// 400
const (
	ErrInvalidProgrammingID = "ERR_INVALID_PROGRAMMING_ID"
	ErrInvalidLabOrPathID   = "ERR_INVALID_LAB_OR_PATH_ID"
	ErrInvalidUserID        = "ERR_INVALID_USER_ID"
	ErrInvalidLabID         = "ERR_INVALID_LAB_ID"
	ErrInvalidLogID         = "ERR_INVALID_LOG_ID"
	ErrInvalidPathID        = "ERR_INVALID_PATH_ID"
	ErrInvalidCreds         = "ERR_INVALID_CREDENTIALS"
	ErrInvalidRole          = "ERR_INVALID_ROLE"

	ErrUsernameUsing = "ERR_USERNAME_BEING_USED"
	ErrStartRoad     = "ERR_NEED_ROAD_START"
	ErrSolvePath     = "ERR_SOLVE_PATH"
)

// 403
const (
	ErrNoPermissionDelete = "ERR_NO_PERMISSION_DELETE"
)

// 404
const (
	ErrUserNotFound                = "ERR_USER_NOT_FOUND"
	ErrRoadMainSHNotFound          = "ERR_MAINSH_NOT_FOUND"
	ErrDockerImageNotFound         = "ERR_IMAGE_NOT_FOUND"
	ErrLabNotFound                 = "ERR_LAB_NOT_FOUND"
	ErrPathNotFound                = "ERR_PATH_NOT_FOUND"
	ErrProgrammingLanguageNotFound = "ERR_PROGRAMMING_LANGUAGE_NOT_FOUND"
	ErrUserCodeNotFound            = "ERR_USERCODE_NOT_FOUND"
	ErrTemplateNotFound            = "ERR_TEMPLATE_NOT_FOUND"
	ErrRoadNotFound                = "ERR_ROAD_NOT_FOUND"
	ErrSolutionsNotFound           = "ERR_SOLUTIONS_NOT_FOUND"
	ErrActivityNotFound            = "ERR_ACTIVITY_NOT_FOUND"
)

// 500
const (
	ErrFilteringUsers              = "ERR_FILTERING_USERS"
	ErrFilteringLogs               = "ERR_FILTERING_LOGS"
	ErrUpdatingUser                = "ERR_UPDATING_USER"
	ErrDeletingUser                = "ERR_DELETING_USER"
	ErrGettingLogs                 = "ERR_FETCHING_LOGS"
	ErrGettingLabs                 = "ERR_FETCHING_LABS"
	ErrGettingRoads                = "ERR_FETCHING_ROADS"
	ErrGettingLevels               = "ERR_FETCHING_LEVELS"
	ErrGettingFrontendTemplate     = "ERR_FETCHING_TEMPLATE"
	ErrGettingProgrammingLanguages = "ERR_FETCHING_PROGRAMMING_LANGUAGES"
	ErrComparingPasswords          = "ERR_COMPARING_PASSWORDS"
	ErrAddingUser                  = "ERR_ADDING_USER"
	ErrAddingLog                   = "ERR_ADDING_LOG"

	ErrDockerImage              = "ERR_DOCKER_IMAGE"
	ErrDockerContainerStop      = "ERR_DOCKER_CONTAINER_STOP"
	ErrDockerLogs               = "ERR_DOCKER_LOGS"
	ErrDockerFileError          = "ERR_DOCKER_FILES"
	ErrDockerCouldNotCreateFile = "ERR_DOCKER_COULD_NOT_CREATE_FILE"

	ErrParserService = "ERR_PARSER_SERVICE"
)

// Log Types
var (
	TypeRoad                = "Road"
	TypePath                = "Path"
	TypeLab                 = "Lab"
	TypeUser                = "User"
	TypeProgrammingLanguage = "ProgrammingLanguage"
)

// Log Content
var (
	ContentStarted   = "Started"
	ContentCompleted = "Completed"
	ContentProfile   = "Profile Updated"
	ContentLevelUp   = "Level Up"
)

//Solve Rates
var (
	Low     = 3  //max -level1
	LowMid  = 6  //max -level2
	Middle  = 12 //max -level3
	MidHigh = 9  //max -level4
	High    = 13 //min -level5
)
