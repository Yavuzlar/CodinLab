package images

import (
	"context"

	"github.com/docker/docker/api/types/image"

	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/docker/docker/client"
)

type IImageManager interface {
	IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error)
	Pull(ctx context.Context, imageReference string) (err error)
}

// Manager, manages the docker Images
type Manager struct {
	cli *client.Client
}

func NewManager(cli *client.Client) IImageManager {
	return &Manager{cli: cli}
}

func (m *Manager) IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error) {
	images, err := m.cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return false, err
	}

	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == imageReference {
				return true, nil
			}
		}
	}

	return false, nil
}

func (m *Manager) Pull(ctx context.Context, imageReference string) error {
	out, err := m.cli.ImagePull(ctx, imageReference, image.PullOptions{})
	if err != nil {
		return service_errors.NewServiceErrorWithMessageAndError(500, "error while pulling an image", err)
	}
	defer out.Close()

	return nil
}
