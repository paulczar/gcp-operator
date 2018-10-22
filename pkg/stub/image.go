package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteImage(project string, image compute.Image) error {
	logrus.Printf("Deleting image %s", image.Name)
	client, err := gce.NewImageService(project, &image)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete()
}

func newImage(project string, image compute.Image) (*compute.Image, error) {
	// log into GCE
	client, err := gce.NewImageService(project, &image)
	if err != nil {
		return nil, err
	}
	// check if image aleady exists before trying to create it
	i, err := client.Get()
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create image
	logrus.Printf("Creating new image %s", image.Name)
	err = client.Create()
	return &image, err
}
