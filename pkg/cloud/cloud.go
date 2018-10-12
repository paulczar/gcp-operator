package cloud

import (
	"fmt"

	"github.com/paulczar/gcp-operator/pkg/cloud/gce"
	compute "google.golang.org/api/compute/v1"
)

type gceCloud struct {
	// GCE client
	client *gce.GCEClient
}

// Cloud interface
type Cloud interface {
	CreateInstance(i *compute.Instance) error
	DeleteInstance(i *compute.Instance) error
	GetInstance(i *compute.Instance) (*compute.Instance, error)
	//RemoveLoadBalancer(cfg *Config, force bool) error
}

func (c *gceCloud) CreateInstance(i *compute.Instance) error {
	fmt.Printf("--> Creating Instance %s\n", i.Name)
	err := c.client.CreateInstance(*i)
	if err != nil {
		return err
	}
	return nil
}

func (c *gceCloud) GetInstance(i *compute.Instance) (*compute.Instance, error) {
	return c.client.GetInstance(*i)
}

func (c *gceCloud) DeleteInstance(i *compute.Instance) error {
	return c.client.DeleteInstance(*i)
}

// New cloud interface
func New(projectID string) (Cloud, error) {
	// try and provision GCE client
	c, err := gce.CreateGCECloud(projectID)
	if err != nil {
		return nil, err
	}

	return &gceCloud{
		client: c,
	}, nil
}
