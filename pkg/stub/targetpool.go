package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteTargetPool(project string, targetPool compute.TargetPool) error {
	logrus.Printf("Deleting targetPool %s", targetPool.Name)
	client, err := gce.New(project)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.TargetPoolDelete(targetPool)
}

func newTargetPool(project string, targetPool compute.TargetPool) (*compute.TargetPool, error) {
	// log into GCE
	client, err := gce.New(project)
	if err != nil {
		return nil, err
	}
	// check if targetPool aleady exists before trying to create it
	i, err := client.TargetPoolGet(targetPool)
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create targetPool
	logrus.Printf("Creating new targetPool %s", targetPool.Name)
	err = client.TargetPoolCreate(targetPool)
	return &targetPool, err
}
