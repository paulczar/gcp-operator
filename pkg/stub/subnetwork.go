package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteSubnetwork(project string, subnetwork compute.Subnetwork) error {
	logrus.Printf("Deleting subnetwork %s", subnetwork.Name)
	client, err := gce.NewSubnetworkService(project, &subnetwork)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete()
}

func newSubnetwork(project string, subnetwork compute.Subnetwork) (*compute.Subnetwork, error) {
	// log into GCE
	client, err := gce.NewSubnetworkService(project, &subnetwork)
	if err != nil {
		return nil, err
	}
	// check if subnetwork aleady exists before trying to create it
	i, err := client.Get()
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create subnetwork
	logrus.Printf("Creating new subnetwork %s", subnetwork.Name)
	err = client.Create()
	return &subnetwork, err
}
