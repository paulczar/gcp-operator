package gce

import (
	"context"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

// GCEClient is a placeholder for GCE stuff.
type GCEClient struct {
	service   *compute.Service
	projectID string
	request   *compute.Instance
}

// CreateGCECloud creates a new instance of GCECloud.
func New(project string) (*GCEClient, error) {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.TODO()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return nil, err
	}
	svc, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	// TODO validate project and network exist
	return &GCEClient{
		service:   svc,
		projectID: project,
	}, nil
}
