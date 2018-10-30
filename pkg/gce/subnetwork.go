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

type SubnetworkService struct {
	GCE     *GCEClient
	Payload *compute.Subnetwork
}

func NewSubnetworkService(project string, payload *compute.Subnetwork) (*SubnetworkService, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &SubnetworkService{
		GCE:     is,
		Payload: payload,
	}, nil
}

// Create an instance.
func (svc *SubnetworkService) Create() error {
	op, err := svc.GCE.compute.Subnetworks.Insert(svc.GCE.projectID, svc.Payload.Region, svc.Payload).Do()
	if err != nil {
		return err
	}
	if err = svc.GCE.waitForRegionOp(op, svc.Payload.Region); err != nil {
		return err
	}
	return nil
}

// Get an Instance
func (svc *SubnetworkService) Get() (*compute.Subnetwork, error) {
	instance, err := svc.GCE.compute.Subnetworks.Get(svc.GCE.projectID, svc.Payload.Region, svc.Payload.Name).Do()
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
func (svc *SubnetworkService) Delete() error {
	op, err := svc.GCE.compute.Subnetworks.Delete(svc.GCE.projectID, svc.Payload.Region, svc.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = svc.GCE.waitForRegionOp(op, svc.Payload.Region); err != nil {
		return err
	}
	return nil
}

// Update an instance
// currently do not support updating an instance
func (svc *SubnetworkService) Update() error {
	return nil
}
