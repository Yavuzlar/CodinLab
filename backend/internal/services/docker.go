package services

import (
	"context"

	"github.com/Yavuzlar/CodinLab/internal/domains"
	service_errors "github.com/Yavuzlar/CodinLab/internal/errors"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type dockerService struct {
	cli   *client.Client
	utils IUtilService
}

func newDockerService(
	utils IUtilService,
) domains.IDockerService {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return &dockerService{
		cli:   cli,
		utils: utils,
	}
}

func (d *dockerService) IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error) {
	// Tüm mevcut imajları listele
	images, err := d.cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return false, err
	}

	// İmaj bulundu mu kontrol et
	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == imageReference {
				return true, nil
			}
		}
	}

	// İmaj bulunamadı
	return false, nil
}

func (d *dockerService) Pull(ctx context.Context, imageReference string) (err error) {
	doneChan := make(chan error, 1)

	go func() {
		out, err := d.cli.ImagePull(ctx, imageReference, image.PullOptions{})
		if err != nil {
			doneChan <- service_errors.NewServiceErrorWithMessageAndError(500, "error while pulling an image", err)
			return
		}
		defer out.Close()

		doneChan <- nil
	}()

	select {
	case err := <-doneChan:
		return err
	case <-ctx.Done():
		return service_errors.NewServiceErrorWithMessageAndError(500, "image pull operation canceled", ctx.Err())
	}
}
