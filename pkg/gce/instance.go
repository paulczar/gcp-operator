package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type Instance interface {
	Create(i *compute.Instance) error
	Get(i *compute.Instance) (*compute.Instance, error)
	Delete(i *compute.Instance) error
	Update(i *compute.Instance) error
}

// Create an instance.
func (gce *GCEClient) Create(payload compute.Instance) error {
	op, err := gce.service.Instances.Insert(gce.projectID, payload.Zone, &payload).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForZoneOp(op, payload.Zone); err != nil {
		return err
	}
	return nil
}

// Get an Instance
func (gce *GCEClient) Get(payload compute.Instance) (*compute.Instance, error) {
	instance, err := gce.service.Instances.Get(gce.projectID, payload.Zone, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return instance, nil
}

// Delete an instance
func (gce *GCEClient) Delete(payload compute.Instance) error {
	op, err := gce.service.Instances.Delete(gce.projectID, payload.Zone, payload.Name).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForZoneOp(op, payload.Zone); err != nil {
		return err
	}
	return nil
}

// Update an instance
// currently do not support updating an instance
func (gce *GCEClient) Update(payload compute.Instance) error {
	return nil
}
