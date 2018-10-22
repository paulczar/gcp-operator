package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteNetwork(project string, network compute.Network) error {
	logrus.Printf("Deleting network %s", network.Name)
	client, err := gce.NewNetworkService(project, &network)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete()
}

func newNetwork(project string, network compute.Network) (*compute.Network, error) {
	// log into GCE
	client, err := gce.NewNetworkService(project, &network)
	if err != nil {
		return nil, err
	}
	// check if network aleady exists before trying to create it
	i, err := client.Get()
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create network
	logrus.Printf("Creating new network %s", network.Name)
	err = client.Create()
	return &network, err
}
