package services

import (
	"strconv"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
)

type commonService struct {
	parserService domains.IParserService
}

func newCommonService(
	parserService domains.IParserService,
) domains.ICommonService {
	return &commonService{
		parserService: parserService,
	}
}

func (s *commonService) GetInventoryInformation(programmingID, language string) (inventoryInformation *domains.InventoryInformation, err error) {
	intProgrammingID, err := strconv.Atoi(programmingID)
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(400, domains.ErrInvalidProgrammingID)
	}

	inventory, err := s.parserService.GetInventory()
	if err != nil {
		return nil, service_errors.NewServiceErrorWithMessage(500, domains.ErrGettingProgrammingLanguages)
	}
	for _, inv := range inventory {
		if inv.ID == intProgrammingID {
			var langInfo domains.InventoryLanguage
			for _, infoL := range inv.Languages {
				if infoL.Lang == language {
					langInfo.SetLang(infoL.Lang)
					langInfo.SetTitle(infoL.Title)
					langInfo.SetDescription(infoL.Description)
				}
			}

			inventoryInformation = domains.NewInventoryInformation(inv.IconPath,inv.Name, inv.DockerImage, inv.FileExtension, inv.PathDir, inv.ID, inv.Cmd, inv.ShCmd, langInfo)
			break
		}
	}
	if inventoryInformation == nil {
		return nil, service_errors.NewServiceErrorWithMessage(404, domains.ErrProgrammingLanguageNotFound)
	}

	return
}
