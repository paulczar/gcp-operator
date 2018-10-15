package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteInstance(project string, instance compute.Instance) error {
	logrus.Printf("Deleting instance %s", instance.Name)
	client, err := gce.New(project)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete(instance)
}

func newInstance(project string, instance compute.Instance) (*compute.Instance, error) {
	// log into GCE
	client, err := gce.New(project)
	if err != nil {
		return nil, err
	}
	// check if instance aleady exists before trying to create it
	i, err := client.Get(instance)
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create instance
	logrus.Printf("Creating new instance %s", instance.Name)
	err = client.Create(instance)
	return &instance, err
}