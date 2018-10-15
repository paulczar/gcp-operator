package stub

import (
	"context"

	ms "github.com/mitchellh/mapstructure"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
	"github.com/sirupsen/logrus"
	compute "google.golang.org/api/compute/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	var err error
	logrus.Debugf("Poll Kubernetes API for changes to known resources.")
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		var instance compute.Instance
		err = ms.Decode(o.Spec.Payload, &instance)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteInstance(o.Spec.ProjectID, instance)
		}
		ni, err := newInstance(o.Spec.ProjectID, instance)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create instance : %v", err)
			return err
		}
		s := v1alpha1.InstanceStatus{
			Status:        ni.Status,
			StatusMessage: ni.StatusMessage,
		}
		if o.Status != s {
			o.Status = s
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
