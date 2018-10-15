package gce

import (
	compute "google.golang.org/api/compute/v1"
)

// Address provides an interface for messing with GCP Addresses
type Address interface {
	AddressCreate(i *compute.Address) error
	AddressGet(i *compute.Address) (*compute.Address, error)
	AddressDelete(i *compute.Address) error
	AddressUpdate(i *compute.Address) error
}

// AddressCreate Creates an Address.
func (gce *GCEClient) AddressCreate(payload compute.Address) error {
	if payload.Region != "" {
		op, err := gce.service.Addresses.Insert(gce.projectID, payload.Region, &payload).Do()
		if err != nil {
			return err
		}
		if err = gce.waitForRegionOp(op, payload.Region); err != nil {
			return err
		}
		return nil
	} else {
		op, err := gce.service.GlobalAddresses.Insert(gce.projectID, &payload).Do()
		if err != nil {
			return err
		}
		if err = gce.waitForGlobalOp(op); err != nil {
			return err
		}
		return nil
	}
}

// AddressGet Gets an Address
func (gce *GCEClient) AddressGet(payload compute.Address) (*compute.Address, error) {
	var address *compute.Address
	var err error
	if payload.Region != "" {
		address, err = gce.service.Addresses.Get(gce.projectID, payload.Region, payload.Name).Do()
	} else {
		address, err = gce.service.GlobalAddresses.Get(gce.projectID, payload.Name).Do()
	}
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return address, nil
}

// AddressDelete Deletes an Address
func (gce *GCEClient) AddressDelete(payload compute.Address) error {
	if payload.Region != "" {
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
	} else {
		op, err := gce.service.GlobalAddresses.Delete(gce.projectID, payload.Name).Do()
		if err != nil {
			if isHTTPErrorCode(err, 404) {
				return nil
			}
			return err
		}
		if err = gce.waitForGlobalOp(op); err != nil {
			return err
		}
		return nil
	}
}

// AddressUpdate Updates an Address
// currently do not support updating an Address
func (gce *GCEClient) AddressUpdate(payload compute.Address) error {
	return nil
}
