package services

import (
	"strconv"

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

func (s *labRoadService) GetInventoryInformation(programmingID string) (inventoryInformation *domains.InventoryInformation, err error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, "Invalid Programming Language ID")
	}

	inventory, err := s.parserService.GetInventory()
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(500, "error by getting programming language information")
	}
	for _, inv := range inventory {
		if inv.ID == intProgrammingID {
			inventoryInformation = domains.NewInventoryInformation(inv.Name, inv.DockerImage, inv.FileExtension, inv.ID, inv.Cmd, inv.ShCmd)
			break
		}
	}
	if inventoryInformation == nil {
		return nil, service_errors.NewServiceErrorWithMessage(404, "Programming Language Not Found")
	}

	return
}
