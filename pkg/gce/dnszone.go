package gce

import (
	dns "google.golang.org/api/dns/v1"
)

type ManagedZone struct {
	GCE     *GCEClient
	Payload *dns.ManagedZone
}

func NewManagedZone(project string, managedZone *dns.ManagedZone) (*ManagedZone, error) {
	is, err := New(project)
	if err != nil {
		return nil, err
	}
	return &ManagedZone{
		GCE:     is,
		Payload: managedZone,
	}, nil
}

// Create an managedZone.
func (is *ManagedZone) Create() error {
	_, err := is.GCE.dns.ManagedZones.Create(is.GCE.projectID, is.Payload).Do()
	if err != nil {
		return err
	}
	return nil
}

// Get an ManagedZone
func (is *ManagedZone) Get() (*dns.ManagedZone, error) {
	managedZone, err := is.GCE.dns.ManagedZones.Get(is.GCE.projectID, is.Payload.Name).Do()
	if err != nil {
		if isHTTPErrorCode(err, 404) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return managedZone, nil
}

// Delete an managedZone
func (is *ManagedZone) Delete() error {
	_ = is.GCE.dns.ManagedZones.Delete(is.GCE.projectID, is.Payload.Name).Do()
	return nil
}

// Update an managedZone
// currently do not support updating an managedZone
func (is *ManagedZone) Update() error {
	return nil
}
