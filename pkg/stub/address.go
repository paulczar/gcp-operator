package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteAddress(project string, address compute.Address) error {
	logrus.Printf("Deleting address %s", address.Name)
	client, err := gce.New(project)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.AddressDelete(address)
}

func newAddress(project string, address compute.Address) (*compute.Address, error) {
	// log into GCE
	client, err := gce.New(project)
	if err != nil {
		return nil, err
	}
	// check if address aleady exists before trying to create it
	i, err := client.AddressGet(address)
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create address
	logrus.Printf("Creating new address %s", address.Name)
	err = client.AddressCreate(address)
	return &address, err
}
