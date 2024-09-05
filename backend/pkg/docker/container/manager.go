package container

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type IContainerManager interface {
	CreateContainer(ctx context.Context, image string) (*container.CreateResponse, error)
	StartContainer(ctx context.Context, containerID string) error
	StopContainer(ctx context.Context, containerID string) error
	ReadContainerLogs(ctx context.Context, containerID string) (string, error)
	RunContainerWithTar(ctx context.Context, image string, cmd []string, tmpCodePath string) (string, error)
}

// Manager, manages the docker containers
type Manager struct {
	client *client.Client
}

func NewManager(cli *client.Client) IContainerManager {
	return &Manager{client: cli}
}

func (m *Manager) CreateContainer(ctx context.Context, image string) (*container.CreateResponse, error) {
	resp, err := m.client.ContainerCreate(ctx, &container.Config{
		Image:        image,
		AttachStdout: true,
		AttachStderr: true,
	}, &container.HostConfig{
		AutoRemove: false,
	}, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("error creating container: %w", err)
	}
	return &resp, nil
}

func (m *Manager) StartContainer(ctx context.Context, containerID string) error {
	err := m.client.ContainerStart(ctx, containerID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("error starting container: %w", err)
	}
	return nil
}

func (m *Manager) StopContainer(ctx context.Context, containerID string) error {
	err := m.client.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		return fmt.Errorf("error stopping container: %w", err)
	}
	return nil
}

func (m *Manager) ReadContainerLogs(ctx context.Context, containerID string) (string, error) {
	out, err := m.client.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", fmt.Errorf("error getting container logs: %w", err)
	}
	defer out.Close()

	var result strings.Builder
	reader := bufio.NewReader(out)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		// Cleans answer by removing docker frame (skips first 8 bytes)
		if len(line) > 8 {
			line = line[8:]
		}
		line = strings.TrimSuffix(line, "\n")

		result.WriteString(line)
	}

	return result.String(), nil
}

func (m *Manager) ContainerIDByName(ctx context.Context, containerName string) (string, error) {
	containers, err := m.client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", fmt.Errorf("error listing containers: %w", err)
	}

	var containerID string
	for _, container := range containers {
		if strings.Contains(container.Names[0], containerName) {
			containerID = container.ID
			break
		}
	}

	if containerID == "" {
		return "", fmt.Errorf("container not found")
	}

	return containerID, nil
}

func (m *Manager) RunContainerWithTar(ctx context.Context, image string, cmd []string, tmpCodePath string) (string, error) {
	// Creates container
	resp, err := m.client.ContainerCreate(ctx, &container.Config{
		Image:        image,
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	}, nil, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("error creating container with command: %w", err)
	}

	// Copy file to container
	if err := m.CopyToContainer(ctx, resp.ID, tmpCodePath, "/go"); err != nil {
		return "", fmt.Errorf("error copying file to container: %w", err)
	}

	// Start Container
	if err := m.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("error starting container with command: %w", err)
	}

	// Wait for container to finish
	statusCh, errCh := m.client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", fmt.Errorf("error waiting for container: %w", err)
		}
	case <-statusCh:
	}

	// Read container logs
	logs, err := m.ReadContainerLogs(ctx, resp.ID)
	if err != nil {
		return "", fmt.Errorf("error reading container logs: %w", err)
	}

	// Remove container
	if err := m.client.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true}); err != nil {
		return "", fmt.Errorf("error removing container: %w", err)
	}

	// Return logs
	return logs, nil
}

func (m *Manager) CopyToContainer(ctx context.Context, containerID, srcPath, destPath string) error {
	// Create tarball
	tarBuffer := new(bytes.Buffer)
	tw := tar.NewWriter(tarBuffer)

	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %w", err)
	}

	header, err := CreateTarHeader("main.go", int64(stat.Mode()), stat.Size())
	if err != nil {
		return fmt.Errorf("error creating tar header: %w", err)
	}

	if err := tw.WriteHeader(header); err != nil {
		return fmt.Errorf("error writing tar header: %w", err)
	}

	if _, err := io.Copy(tw, file); err != nil {
		return fmt.Errorf("error writing file content to tarball: %w", err)
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("error closing tar writer: %w", err)
	}

	// Copy tarball to container
	tarReader := bytes.NewReader(tarBuffer.Bytes())
	if err := m.client.CopyToContainer(ctx, containerID, destPath, tarReader, types.CopyToContainerOptions{}); err != nil {
		return fmt.Errorf("error copying file to container: %w", err)
	}

	return nil
}
