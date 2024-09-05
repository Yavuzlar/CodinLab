package docker

import (
	"github.com/Yavuzlar/CodinLab/pkg/docker/container"
	"github.com/Yavuzlar/CodinLab/pkg/docker/images"
	"github.com/docker/docker/client"
)

type IDockerSDK interface {
	Container() container.IContainerManager
	Images() images.IImageManager
}

// This is the main struct that represents docker sdk.
type DockerSDK struct {
	client           *client.Client
	containerManager container.IContainerManager
	imageManager     images.IImageManager
}

// Creates new docker sdk.
func NewDockerSDK() (IDockerSDK, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerSDK{
		client:           cli,
		containerManager: container.NewManager(cli),
		imageManager:     images.NewManager(cli),
	}, nil
}

func (d *DockerSDK) Container() container.IContainerManager {
	return d.containerManager
}

func (d *DockerSDK) Images() images.IImageManager {
	return d.imageManager
}
