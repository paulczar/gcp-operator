package gce

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

// GCEClient is a placeholder for GCE stuff.
type GCEClient struct {
	service   *compute.Service
	projectID string
	//networkURL string
}

// CreateGCECloud creates a new instance of GCECloud.
func CreateGCECloud(project string) (*GCEClient, error) {
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

// CreateInstance returns an instance by name
func (gce *GCEClient) CreateInstance(ci compute.Instance) error {
	fmt.Printf("Creating Instance %s in %s\n", ci.Name, ci.Zone)
	op, err := gce.service.Instances.Insert(gce.projectID, ci.Zone, &ci).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForZoneOp(op, ci.Zone); err != nil {
		return err
	}
	return nil
}

// GetInstance returns an instance group by name
func (gce *GCEClient) GetInstance(gi compute.Instance) (*compute.Instance, error) {
	ig, err := gce.service.Instances.Get(gce.projectID, gi.Zone, gi.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return ig, nil
}

// DeleteInstance returns an instance group by name
func (gce *GCEClient) DeleteInstance(di compute.Instance) error {
	op, err := gce.service.Instances.Delete(gce.projectID, di.Zone, di.Name).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForZoneOp(op, di.Zone); err != nil {
		return err
	}
	return nil
}
