package services

import (
	"context"
	"fmt"
	"os"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	"github.com/Yavuzlar/CodinLab/pkg/docker"
)

type codeService struct {
	dockerSDK docker.IDockerSDK
}

func NewCodeService() domains.ICodeService {
	dockerSDK, err := docker.NewDockerSDK()
	if err != nil {
		panic(err)
	}

	return &codeService{
		dockerSDK: dockerSDK,
	}
}

func (s *codeService) Pull(ctx context.Context, imageReference string) (err error) {
	return s.dockerSDK.Images().Pull(ctx, imageReference)
}

func (s *codeService) IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error) {
	return s.dockerSDK.Images().IsImageExists(ctx, imageReference)
}

func (s *codeService) RunContainerWithTar(ctx context.Context, image, tmpCodePath string, cmd []string) (string, error) {
	return s.dockerSDK.Container().RunContainerWithTar(ctx, image, cmd, tmpCodePath)
}

// Bunu Answer Kısmınlarında Kullanacaksın.
func (s *codeService) UploadUserCode(ctx context.Context, userID string, programmingLanguageID, PathLabID int, codeType, fileExtention, content string) (string, error) {
	if err := s.createCodeFile(userID); err != nil {
		return "", err
	}

	mainDir := "usercodes"
	userDir := mainDir + "/" + userID
	labDir := fmt.Sprintf("%v/labs/", userDir)
	pathDir := fmt.Sprintf("%v/paths/", userDir)
	fileName := fmt.Sprintf("%v-%v.%v", programmingLanguageID, PathLabID, fileExtention)

	var codePath, codeTmpPath string

	if codeType == domains.TypeLab {
		codePath = labDir + fileName
		codeTmpPath = labDir + "tmp-" + fileName
	} else if codeType == domains.TypePath {
		codePath = pathDir + fileName
		codeTmpPath = pathDir + "tmp-" + fileName
	}

	if codeType == domains.TypeLab {
		if err := s.CreateFile(codePath, content); err != nil {
			return "", err
		}
		if err := s.CreateFile(codeTmpPath, ""); err != nil {
			return "", err
		}
	}

	if codeType == domains.TypePath {
		if err := s.CreateFile(codePath, content); err != nil {
			return "", err
		}
		if err := s.CreateFile(codeTmpPath, ""); err != nil {
			return "", err
		}
	}

	return codeTmpPath, nil
}

func (s *codeService) createCodeFile(userID string) (err error) {
	mainDir := "usercodes"
	userDir := fmt.Sprintf("%v/%v", mainDir, userID)
	labDir := fmt.Sprintf("%v/labs", userDir)
	pathDir := fmt.Sprintf("%v/paths", userDir)

	// Check and create mainDir if not exists
	if err := s.checkDir(mainDir); err != nil {
		if err = s.createDir(mainDir); err != nil {
			return err
		}
	}

	// Check and create userDir, labDir, pathDir
	if err := s.checkDir(userDir); err != nil {
		if err = s.createDir(userDir); err != nil {
			return err
		}
	}

	// Check and create labDir if not exists
	if err := s.checkDir(labDir); err != nil {
		if err = s.createDir(labDir); err != nil {
			return err
		}
	}

	// Check and create pathDir if not exists
	if err := s.checkDir(pathDir); err != nil {
		if err = s.createDir(pathDir); err != nil {
			return err
		}
	}

	return nil
}

func (s *codeService) checkDir(dir string) (err error) {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return err
	}

	if !fileInfo.IsDir() {
		return os.ErrNotExist
	}

	return nil
}

func (s *codeService) createDir(dir string) (err error) {
	err = os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (s *codeService) CreateFile(filePath, content string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if content != "" {
		file.WriteString(content)
	}

	defer file.Close()

	return nil
}
