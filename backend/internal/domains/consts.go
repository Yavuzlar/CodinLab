package domains

// 400
const (
	ErrInvalidProgrammingID = "invalid programming language id"
	ErrInvalidLabOrPathID   = "invalid lab or path id"
	ErrInvalidUserID        = "invalid user id"
	ErrInvalidLabID         = "invalid lab id"
	ErrInvalidLogID         = "invalid log id"
	ErrInvalidPathID        = "invalid path id"
	ErrInvalidCreds         = "invalid credentials"
	ErrInvalidRole          = "invalid role"

	ErrUsernameUsing = "username already being used"
	ErrStartRoad     = "you need to start road"
	ErrSolvePath     = "you need to solve %d. path first"
)

// 403
const (
	ErrNoPermissionDelete = "no permission to delete"
)

// 404
const (
	ErrUserNotFound                = "user not found"
	ErrRoadMainSHNotFound          = "main.sh not found, could not be read"
	ErrDockerImageNotFound         = "Image not found "
	ErrLabNotFound                 = "lab not found"
	ErrPathNotFound                = "path not found"
	ErrProgrammingLanguageNotFound = "programming language not found"
	ErrUserCodeNotFound            = "user code not found"
	ErrTemplateNotFound            = "template not found"
	ErrRoadNotFound                = "road not found"
)

// 500
const (
	ErrFilteringUsers              = "error while filtering user"
	ErrFilteringLogs               = "error while filtering logs"
	ErrUpdatingUser                = "error while updating user"
	ErrDeletingUser                = "error while deleting user"
	ErrGettingLogs                 = "error while getting logs"
	ErrGettingLabs                 = "error while getting labs"
	ErrGettingRoads                = "error while getting roads"
	ErrGettingLevels               = "error while getting levels"
	ErrGettingFrontendTemplate     = "Error while getting frontend template"
	ErrGettingProgrammingLanguages = "error while getting programming languages"
	ErrComparingPasswords          = "error while comparing passwords"
	ErrAddingUser                  = "error while adding user"
	ErrAddingLog                   = "error while adding log"

	ErrDockerImage              = "Docker Image Function Has an Error"
	ErrDockerContainerStop      = "Container Stop Error"
	ErrDockerLogs               = "unable to read docker logs"
	ErrDockerFileError          = "file error"
	ErrDockerCouldNotCreateFile = "could not create file and write data into it"

	ErrParserService = "parser service error"
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
