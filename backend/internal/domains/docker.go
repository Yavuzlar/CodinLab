package domains

import "context"

// IDockerService is the interface that provides the methods for the docker service.
type IDockerService interface {
	IsImageExists(ctx context.Context, imageReference string) (isExsits bool, err error)
	Pull(ctx context.Context, lang string) (err error)
}
