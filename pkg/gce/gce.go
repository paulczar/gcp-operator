package gce

import (
	"context"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	dns "google.golang.org/api/dns/v1"
)

// GCEClient is a placeholder for GCE stuff.
type GCEClient struct {
	compute   *compute.Service
	dns       *dns.Service
	projectID string
}

// CreateGCECloud creates a new instance of GCECloud.
func New(project string) (*GCEClient, error) {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.TODO()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return nil, err
	}
	c, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	d, err := dns.New(client)
	if err != nil {
		return nil, err
	}
	if project == "" {
		credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
		if err != nil {
			return nil, err
		}
		project = credentials.ProjectID
	}
	// TODO validate project and network exist
	return &GCEClient{
		compute:   c,
		dns:       d,
		projectID: project,
	}, nil
}
