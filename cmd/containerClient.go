package containerClient

import (
	"context"
	"errors"
	// "fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
)

// NewClusterClient creates a new client to GKE
func NewClusterClient(pid, zone string) (*ClusterClient, error) {
	cl, err := google.DefaultClient(context.Background(), container.CloudPlatformScope)
	if err != nil {
		return nil, errors.New("Creating Cloud Client")
	}
	containerSvc, err := container.New(cl)
	if err != nil {
		return nil, errors.New("Initializing Container Service")
	}
	projSvc := container.NewProjectsService(containerSvc)
	cc := &ClusterClient{
		ContainerSvc: containerSvc,
		ProjectsSvc:  projSvc,

		pid:  pid,
		zone: zone,
	}
	return cc, nil
}

// ClusterClient is the top level struct for this package
type ClusterClient struct {
	ContainerSvc *container.Service
	ProjectsSvc  *container.ProjectsService

	pid  string
	zone string
}
