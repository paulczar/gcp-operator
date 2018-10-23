package stub

import (
	"context"
	"encoding/json"
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
		res, err := newInstance(p, instance)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus(res.Status, res.StatusMessage, string(oj), err)
		if o.Status != s {
			o.Status = s
			updateSDK = true
		}
		if instance.MachineType != res.MachineType {
			o.Spec["MachineType"] = res.MachineType
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
		res, err := newAddress(p, address)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus(res.Status, "", string(oj), err)
		if o.Status != s {
			o.Status = s
			updateSDK = true
		}
		if o.Spec["Address"] != res.Address {
			o.Spec["Address"] = res.Address
			updateSDK = true
		}
		if o.Spec["SelfLink"] != res.SelfLink {
			o.Spec["SelfLink"] = res.SelfLink
			updateSDK = true
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
		res, err := newForwardingRule(p, rule)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus("", "", string(oj), err)
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
		res, err := newTargetPool(p, tp)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus("", "", string(oj), err)
		if o.Status != s {
			o.Status = s
			updateSDK = true
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
		res, err := newNetwork(p, svc)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus("", "", string(oj), err)
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
		var err error
		svc := compute.Subnetwork{}
		p := getProjectID(o.ObjectMeta)
		if err := mapstructure.Decode(o.Spec, &svc); err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteSubnetwork(p, svc)
		}
		res, err := newSubnetwork(p, svc)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus("", "", string(oj), err)
		if o.Status != s {
			o.Status = s
			updateSDK = true
		} // todo update if different
		if updateSDK {
			err := sdk.Update(o)
			if err != nil {
				return err
			}
		}
		if err != nil && !errors.IsAlreadyExists(err) {
			return err
		}

	case *v1alpha1.Image:
		var err error
		svc := compute.Image{}
		p := getProjectID(o.ObjectMeta)
		if err := mapstructure.Decode(o.Spec, &svc); err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteImage(p, svc)
		}
		res, err := newImage(p, svc)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus(res.Status, "", string(oj), err)
		if o.Status != s {
			o.Status = s
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

	case *v1alpha1.Firewall:
		var err error
		svc := compute.Firewall{}
		p := getProjectID(o.ObjectMeta)
		if err := mapstructure.Decode(o.Spec, &svc); err != nil {
			panic(err)
		}
		if event.Deleted {
			return deleteFirewall(p, svc)
		}
		res, err := newFirewall(p, svc)
		oj, ojerr := json.Marshal(res)
		if ojerr != nil {
			oj = []byte{}
		}
		s := getStatus("", "", string(oj), err)
		if o.Status != s {
			o.Status = s
			updateSDK = true
		} // todo update if different
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
