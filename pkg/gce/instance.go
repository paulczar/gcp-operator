package gce

import (
	compute "google.golang.org/api/compute/v1"
)

//type Instance interface {
//	InstanceCreate(i *compute.Instance) error
//	InstanceGet(i *compute.Instance) (*compute.Instance, error)
//	InstanceDelete(i *compute.Instance) error
//	InstanceUpdate(i *compute.Instance) error
//}

type InstanceService struct {
	GCE     *GCEClient
	Payload *compute.Instance
}

func NewInstanceService(project string, instance *compute.Instance) (*InstanceService, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &InstanceService{
		GCE:     is,
		Payload: instance,
	}, nil
}

// Create an instance.
func (is *InstanceService) Create() error {
	is.Payload.MachineType = machineTypeURL(is.GCE.projectID, is.Payload.Zone, is.Payload.MachineType)
	for i, n := range is.Payload.NetworkInterfaces {
		nu := networkURL(is.GCE.projectID, n.Network)
		if n.Network != nu {
			is.Payload.NetworkInterfaces[i].Network = nu
		}
	}
	op, err := is.GCE.service.Instances.Insert(is.GCE.projectID, is.Payload.Zone, is.Payload).Do()
	if err != nil {
		return err
	}
	if err = is.GCE.waitForZoneOp(op, is.Payload.Zone); err != nil {
		return err
	}
	return nil
}

// Get an Instance
func (is *InstanceService) Get() (*compute.Instance, error) {
	instance, err := is.GCE.service.Instances.Get(is.GCE.projectID, is.Payload.Zone, is.Payload.Name).Do()
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
func (is *InstanceService) Delete() error {
	op, err := is.GCE.service.Instances.Delete(is.GCE.projectID, is.Payload.Zone, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = is.GCE.waitForZoneOp(op, is.Payload.Zone); err != nil {
		return err
	}
	return nil
}

// Update an instance
// currently do not support updating an instance
func (is *InstanceService) Update() error {
	return nil
}
