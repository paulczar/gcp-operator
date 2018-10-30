package gce

import (
	compute "google.golang.org/api/compute/v1"
)

type FirewallService struct {
	GCE     *GCEClient
	Payload *compute.Firewall
}

func NewFirewallService(project string, firewall *compute.Firewall) (*FirewallService, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &FirewallService{
		GCE:     is,
		Payload: firewall,
	}, nil
}

// Create an firewall.
func (is *FirewallService) Create() error {
	op, err := is.GCE.compute.Firewalls.Insert(is.GCE.projectID, is.Payload).Do()
	if err != nil {
		return err
	}
	if err = is.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Get an Firewall
func (is *FirewallService) Get() (*compute.Firewall, error) {
	firewall, err := is.GCE.compute.Firewalls.Get(is.GCE.projectID, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return firewall, nil
}

// Delete an firewall
func (is *FirewallService) Delete() error {
	op, err := is.GCE.compute.Firewalls.Delete(is.GCE.projectID, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil
		}
		return err
	}
	if err = is.GCE.waitForGlobalOp(op); err != nil {
		return err
	}
	return nil
}

// Update an firewall
// currently do not support updating an firewall
func (is *FirewallService) Update() error {
	return nil
}
