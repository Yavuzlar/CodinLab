package dto

type DTOManager struct {
	UserDTOManager *UserDTOManager
	LogDTOManager  *LogDTOManager
	RoadDTOManager *RoadDTOManager
	HomeDTOManager *HomeDTOManager
	LabDTOManager  *LabDTOManager
}

func CreateNewDTOManager() *DTOManager {
	userDTOManager := NewUserDTOManager()
	logDTOManager := NewLogDTOManager()
	roadDTOManager := NewRoadDTOManager()
	homeDTOManager := NewHomeDTOManager()
	labDTOManager := NewLabDTOManager()

	return &DTOManager{
		UserDTOManager: &userDTOManager,
		LogDTOManager:  &logDTOManager,
		RoadDTOManager: &roadDTOManager,
		HomeDTOManager: &homeDTOManager,
		LabDTOManager:  &labDTOManager,
	}
}
