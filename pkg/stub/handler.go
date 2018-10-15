package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
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
	logrus.Debugf("Poll Kubernetes API for changes to known resources.")
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		if event.Deleted {
			return deleteInstance(o.Spec.ProjectID, *o.Spec.Payload)
		}
		ni, err := newInstance(o.Spec.ProjectID, *o.Spec.Payload)
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

	case *v1alpha1.Address:
		if event.Deleted {
			return deleteAddress(o.Spec.ProjectID, *o.Spec.Payload)
		}
		na, err := newAddress(o.Spec.ProjectID, *o.Spec.Payload)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create address : %v", err)
			return err
		}
		if o.Spec.Payload.Address != na.Address {
			o.Spec.Payload.Address = na.Address
		}
		if o.Spec.Payload.SelfLink != na.SelfLink {
			o.Spec.Payload.SelfLink = na.SelfLink
		}
		s := v1alpha1.AddressStatus{
			Status: na.Status,
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
