package stub

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
	"github.com/paulczar/gcp-operator/pkg/gce"
	"github.com/sirupsen/logrus"
)

func deleteInstance(cr *v1alpha1.Instance) error {
	logrus.Printf("Deleting instance %s", cr.Spec.Instance.Name)
	client, err := gce.New(cr.Spec.ProjectID)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.Delete(cr.Spec.Instance)
}

func newInstance(cr *v1alpha1.Instance) error {
	// log into GCE
	client, err := gce.New(cr.Spec.ProjectID)
	if err != nil {
		return err
	}
	// check if instance aleady exists before trying to create it
	i, err := client.Get(cr.Spec.Instance)
	if err != nil {
		return err
	}

	// if it exists, update resource's status
	if i != nil {
		s := v1alpha1.InstanceStatus{
			Status:        i.Status,
			StatusMessage: i.StatusMessage,
		}
		if cr.Status != s {
			cr.Status = s
			err := sdk.Update(cr)
			if err != nil {
				return err
			}
		}
		return nil
	}
	// create instance
	logrus.Printf("Creating new instance %s", cr.Spec.Instance.Name)
	err = client.Create(cr.Spec.Instance)
	if err != nil {
		cr.Status = v1alpha1.InstanceStatus{
			Status:        "FAILED",
			StatusMessage: err.Error(),
		}
		err := sdk.Update(cr)
		if err != nil {
			return err
		}
	}
	return err
}
