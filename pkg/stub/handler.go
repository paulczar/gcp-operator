package stub

import (
	"context"
	"fmt"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
	"github.com/paulczar/gcp-operator/pkg/cloud"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		//err := sdk.Create(newInstance(o))
		if event.Deleted {
			return deleteInstance(o)
		}
		err := newInstance(o)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create instance : %v", err)
			return err
		}
	}
	return nil
}

func deleteInstance(cr *v1alpha1.Instance) error {
	client, err := cloud.New(cr.Spec.ProjectID)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	return client.DeleteInstance(&cr.Spec.Instance)
}

func newInstance(cr *v1alpha1.Instance) error {
	client, err := cloud.New(cr.Spec.ProjectID)
	if err != nil {
		panic(err)
	}
	//spew.Dump(cr)
	i, err := client.GetInstance(&cr.Spec.Instance)
	if err != nil {
		return err
	}
	if i != nil {
		fmt.Printf("instance %s in %s already exists\n", i.Name, i.Zone)
		cr.Status = v1alpha1.InstanceStatus{
			Status:        i.Status,
			StatusMessage: i.StatusMessage,
		}
		_ = sdk.Update(cr)
		return nil
	}
	err = client.CreateInstance(&cr.Spec.Instance)
	if err != nil {
		panic(err)
	}
	return err
}
