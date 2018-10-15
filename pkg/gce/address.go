package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type Address interface {
	AddressCreate(i *compute.Address) error
	AddressGet(i *compute.Address) (*compute.Address, error)
	AddressDelete(i *compute.Address) error
	AddressUpdate(i *compute.Address) error
}

// Create an Address.
func (gce *GCEClient) AddressCreate(payload compute.Address) error {
	op, err := gce.service.Addresses.Insert(gce.projectID, payload.Region, &payload).Do()
	if err != nil {
		return err
	}
	if err = gce.waitForRegionOp(op, payload.Region); err != nil {
		return err
	}
	return nil
}

// Get an Address
func (gce *GCEClient) AddressGet(payload compute.Address) (*compute.Address, error) {
	Address, err := gce.service.Addresses.Get(gce.projectID, payload.Region, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Address, nil
}

// Delete an Address
func (gce *GCEClient) AddressDelete(payload compute.Address) error {
	op, err := gce.service.Addresses.Delete(gce.projectID, payload.Region, payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = gce.waitForRegionOp(op, payload.Region); err != nil {
		return err
	}
	return nil
}

// Update an Address
// currently do not support updating an Address
func (gce *GCEClient) AddressUpdate(payload compute.Address) error {
	return nil
}
