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
	var updateSDK = false
	logrus.Debugf("Poll Kubernetes API for changes to known resources.")
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		p := getProjectID(o.ObjectMeta)
		if event.Deleted {
			return deleteInstance(p, *o.Spec.Instance)
		}
		ni, err := newInstance(p, *o.Spec.Instance)
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
			updateSDK = true
		}
		if o.Spec.Instance.MachineType != ni.MachineType {
			o.Spec.Instance.MachineType = ni.MachineType
			updateSDK = true
		}
		for i, v := range o.Spec.Instance.NetworkInterfaces {
			if v != ni.NetworkInterfaces[i] {
				o.Spec.Instance.NetworkInterfaces = ni.NetworkInterfaces
				updateSDK = true
			}
		}
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}

	case *v1alpha1.Address:
		p := getProjectID(o.ObjectMeta)
		if event.Deleted {
			return deleteAddress(p, *o.Spec.Address)
		}
		na, err := newAddress(p, *o.Spec.Address)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create address : %v", err)
			return err
		}
		if o.Spec.Address.Address != na.Address {
			o.Spec.Address.Address = na.Address
			updateSDK = true
		}
		if o.Spec.Address.SelfLink != na.SelfLink {
			o.Spec.Address.SelfLink = na.SelfLink
			updateSDK = true
		}
		s := v1alpha1.AddressStatus{
			Status: na.Status,
		}
		if o.Status != s {
			o.Status = s
			updateSDK = true
		}
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
