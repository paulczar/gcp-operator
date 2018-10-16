package stub

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
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
	var updateSDK = false
	logrus.Debugf("Poll Kubernetes API for changes to known resources.")
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		p := getProjectID(o.ObjectMeta)
		o.Spec.Name = setName(o.Name, o.Spec.Name)
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
		// todo this is probably a bad test because of unstrict ordering. fix it.
		for i, v := range o.Spec.Instance.NetworkInterfaces {
			if v.Fingerprint != ni.NetworkInterfaces[i].Fingerprint {
				fmt.Printf("update from %v to %v\n", v, ni.NetworkInterfaces[i])
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
		o.Spec.Name = setName(o.Name, o.Spec.Name)
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

	case *v1alpha1.ForwardingRule:
		p := getProjectID(o.ObjectMeta)
		o.Spec.Name = setName(o.Name, o.Spec.Name)
		if event.Deleted {
			return deleteForwardingRule(p, *o.Spec.ForwardingRule)
		}
		na, err := newForwardingRule(p, *o.Spec.ForwardingRule)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create Forwarding Rule : %v", err)
			return err
		}
		if o.Spec.ForwardingRule != na {
			o.Spec.ForwardingRule = na
			updateSDK = true
		}
		s := v1alpha1.ForwardingRuleStatus{Status: "CREATED"}
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
	// todo provide a way to specify a list of instances from CRD above (tags? labels? metadata?)
	case *v1alpha1.TargetPool:
		p := getProjectID(o.ObjectMeta)
		//o.Spec.Name = setName(o.Name, o.Spec.Name)
		tp := compute.TargetPool{}
		err := mapstructure.Decode(o.Spec.Resource, &tp)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteTargetPool(p, tp)
		}
		_, err = newTargetPool(p, tp)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create Target Pool : %v", err)
			return err
		}
		//		if tp != *na {
		//			o.Spec.Resource = na
		//			updateSDK = true
		//		}
		s := v1alpha1.TargetPoolStatus{Status: "CREATED"}
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
	default:
		fmt.Println("OMG WTF I DUNNO")
	}

	return nil
}
