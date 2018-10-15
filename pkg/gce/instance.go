package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type Instance interface {
	InstanceCreate(i *compute.Instance) error
	InstanceGet(i *compute.Instance) (*compute.Instance, error)
	InstanceDelete(i *compute.Instance) error
	InstanceUpdate(i *compute.Instance) error
}

// Create an instance.
func (gce *GCEClient) InstanceCreate(payload compute.Instance) error {
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
func (gce *GCEClient) InstanceGet(payload compute.Instance) (*compute.Instance, error) {
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
func (gce *GCEClient) InstanceDelete(payload compute.Instance) error {
	op, err := gce.service.Instances.Delete(gce.projectID, payload.Zone, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = gce.waitForZoneOp(op, payload.Zone); err != nil {
		return err
	}
	return nil
}

// Update an instance
// currently do not support updating an instance
func (gce *GCEClient) InstanceUpdate(payload compute.Instance) error {
	return nil
}
