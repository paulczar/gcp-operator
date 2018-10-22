package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteFirewall(project string, firewall compute.Firewall) error {
	logrus.Printf("Deleting firewall %s", firewall.Name)
	client, err := gce.NewFirewallService(project, &firewall)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete()
}

func newFirewall(project string, firewall compute.Firewall) (*compute.Firewall, error) {
	// log into GCE
	client, err := gce.NewFirewallService(project, &firewall)
	if err != nil {
		return nil, err
	}
	// check if firewall aleady exists before trying to create it
	i, err := client.Get()
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create firewall
	logrus.Printf("Creating new firewall %s", firewall.Name)
	err = client.Create()
	return &firewall, err
}
