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
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	var updateSDK = false
	logrus.Debugf("Poll Kubernetes API for changes to known resources.")
	switch o := event.Object.(type) {
	case *v1alpha1.Instance:
		instance := compute.Instance{}
		p := getProjectID(o.ObjectMeta)
		err := mapstructure.Decode(o.Spec, &instance)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteInstance(p, instance)
		}
		ni, err := newInstance(p, instance)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create instance : %v", err)
			return err
		}
		s := v1alpha1.ServiceStatus{
			Status: ni.Status,
			//StatusMessage: ni.StatusMessage,
		}
		if o.Status != s {
			o.Status = s
			updateSDK = true
		}
		if instance.MachineType != ni.MachineType {
			o.Spec["MachineType"] = ni.MachineType
			updateSDK = true
		}
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}

	case *v1alpha1.Address:
		p := getProjectID(o.ObjectMeta)
		address := compute.Address{}
		err := mapstructure.Decode(o.Spec, &address)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteAddress(p, address)
		}
		na, err := newAddress(p, address)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create address : %v", err)
			return err
		}
		if o.Spec["Address"] != na.Address {
			o.Spec["Address"] = na.Address
			updateSDK = true
		}
		if o.Spec["SelfLink"] != na.SelfLink {
			o.Spec["SelfLink"] = na.SelfLink
			updateSDK = true
		}
		s := v1alpha1.ServiceStatus{
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
		rule := compute.ForwardingRule{}
		//status := v1alpha1.ServiceStatus{}
		err := mapstructure.Decode(o.Spec, &rule)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteForwardingRule(p, rule)
		}
		_, err = newForwardingRule(p, rule)
		s := getStatus("", "", err)
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
		if err != nil {
			return err
		}

	// todo provide a way to specify a list of instances from CRD above (tags? labels? metadata?)
	case *v1alpha1.TargetPool:
		p := getProjectID(o.ObjectMeta)
		//o.Spec.Name = setName(o.Name, o.Spec.Name)
		tp := compute.TargetPool{}
		err := mapstructure.Decode(o.Spec, &tp)
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
		s := v1alpha1.ServiceStatus{Status: "CREATED"}
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

	case *v1alpha1.Network:
		svc := compute.Network{}
		p := getProjectID(o.ObjectMeta)
		err := mapstructure.Decode(o.Spec, &svc)
		if err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteNetwork(p, svc)
		}
		_, err = newNetwork(p, svc)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create network : %v", err)
			return err
		}
		s := v1alpha1.ServiceStatus{
			Status: "CREATED",
			//StatusMessage: ni.StatusMessage,
		}
		if o.Status != s {
			o.Status = s
			updateSDK = true
		}
		// todo update subnet list if different
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}

	case *v1alpha1.Subnetwork:
		var status v1alpha1.ServiceStatus
		var err error
		svc := compute.Subnetwork{}
		p := getProjectID(o.ObjectMeta)
		if err := mapstructure.Decode(o.Spec, &svc); err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteSubnetwork(p, svc)
		}
		_, err = newSubnetwork(p, svc)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create subnetwork : %v", err)
			status.Status = "FAILED"
			status.Message = err.Error()
		} else {
			status.Status = "CREATED"
		}
		if o.Status != status {
			o.Status = status
			updateSDK = true
		}
		// todo update if different
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}

	default:
		fmt.Println("OMG WTF I DUNNO")
	}

	return nil
}

func getStatus(stat, message string, err error) v1alpha1.ServiceStatus {
	var s = v1alpha1.ServiceStatus{
		Status:  stat,
		Message: message,
	}
	if err != nil && !errors.IsAlreadyExists(err) {
		s = v1alpha1.ServiceStatus{
			Status:  "FAILED",
			Message: err.Error(),
		}
		return s
	}
	return s
}
