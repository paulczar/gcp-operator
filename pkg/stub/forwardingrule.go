package stub

import (
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
)

func deleteForwardingRule(project string, forwardingRule compute.ForwardingRule) error {
	logrus.Printf("Deleting forwardingRule %s", forwardingRule.Name)
	client, err := gce.New(project)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.ForwardingRuleDelete(forwardingRule)
}

func newForwardingRule(project string, forwardingRule compute.ForwardingRule) (*compute.ForwardingRule, error) {
	// log into GCE
	client, err := gce.New(project)
	if err != nil {
		return nil, err
	}
	// check if forwardingRule aleady exists before trying to create it
	i, err := client.ForwardingRuleGet(forwardingRule)
	if err != nil {
		return nil, err
	}
	if i != nil {
		return i, nil
	}

	// create forwardingRule
	logrus.Printf("Creating new forwardingRule %s", forwardingRule.Name)
	err = client.ForwardingRuleCreate(forwardingRule)
	return &forwardingRule, err
}
