package services

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
)

type startService struct {
	utils         IUtilService
	parserService domains.IParserService
}

func newStartService(
	utils IUtilService,
	parserService domains.IParserService,
) domains.IStartService {
	return &startService{
		utils:         utils,
		parserService: parserService,
	}
}

// Retrieves name, dockerImage and icon path
func (s *startService) GetProgrammingInformation(programmingID int) (*domains.StartProgramming, error) {
	src, err := s.parserService.GetInventory()
	if err != nil {
		return nil, err
	}

	var programmingLanguage domains.StartProgramming
	for _, pl := range src {
		if pl.ID == programmingID {
			programmingLanguage.SetProgrammingID(programmingID)
			programmingLanguage.SetDockerImage(pl.DockerImage)
			break
		}
	}

	return &programmingLanguage, nil
}
