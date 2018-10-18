package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteInstance(project string, instance compute.Instance) error {
	logrus.Printf("Deleting instance %s", instance.Name)
	client, err := gce.NewInstanceService(project, &instance)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete()
}

func newInstance(project string, instance compute.Instance) (*compute.Instance, error) {
	// log into GCE
	client, err := gce.NewInstanceService(project, &instance)
	if err != nil {
		return nil, err
	}
	// check if instance aleady exists before trying to create it
	i, err := client.Get()
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create instance
	logrus.Printf("Creating new instance %s", instance.Name)
	err = client.Create()
	return &instance, err
}
