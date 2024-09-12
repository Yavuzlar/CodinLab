package services

import (
	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

type labRoadService struct {
	parserService domains.IParserService
}

func newLabRoadService(
	parserService domains.IParserService,
) domains.ILabRoadService {
	return &labRoadService{
		parserService: parserService,
	}
}

func (s *labRoadService) GetInventoryInformation(programmingID int32) (inventorys *domains.InventoryInformation, err error) {
	inventory, err := s.parserService.GetInventory()
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(500, "error by getting programming language information")
	}
	for _, inv := range inventory {
		if inv.ID == int(programmingID) {
			inventorys = domains.NewInventoryInformation(inv.Name, inv.DockerImage, inv.FileExtension, inv.ID, inv.Cmd)
			break
		}
	}
	return
}
