package domains

// IRoadService is the interface that provides the methods for the road service.
type IStartService interface {
	GetProgrammingInformation(programmingID int) (programmingLanguage *StartProgramming, err error)
}

type StartProgramming struct {
	programmingID int
	dockerImage   string
}

func (s *StartProgramming) GetProgrammingID() int {
	return s.programmingID
}

func (s *StartProgramming) GetDockerImage() string {
	return s.dockerImage
}

func (s *StartProgramming) SetProgrammingID(programmingID int) {
	s.programmingID = programmingID
}

func (s *StartProgramming) SetDockerImage(dockerImage string) {
	s.dockerImage = dockerImage
}
