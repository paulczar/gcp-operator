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

type NetworkService struct {
	GCE     *GCEClient
	Payload *compute.Network
}

func NewNetworkService(project string, payload *compute.Network) (*NetworkService, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &NetworkService{
		GCE:     is,
		Payload: payload,
	}, nil
}

// Create an instance.
func (svc *NetworkService) Create() error {
	// otherwise will create a legacy network and baby jesus will cry
	svc.Payload.ForceSendFields = []string{"AutoCreateSubnetworks"}
	if !svc.Payload.AutoCreateSubnetworks {
		svc.Payload.AutoCreateSubnetworks = false
	}
	op, err := svc.GCE.service.Networks.Insert(svc.GCE.projectID, svc.Payload).Do()
	if err != nil {
		return err
	}
	if err = svc.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Get an Instance
func (svc *NetworkService) Get() (*compute.Network, error) {
	instance, err := svc.GCE.service.Networks.Get(svc.GCE.projectID, svc.Payload.Name).Do()
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
func (svc *NetworkService) Delete() error {
	op, err := svc.GCE.service.Networks.Delete(svc.GCE.projectID, svc.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = svc.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Update an instance
// currently do not support updating an instance
func (svc *NetworkService) Update() error {
	return nil
}
