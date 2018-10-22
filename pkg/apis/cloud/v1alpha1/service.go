package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Service `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// after generate fix the zz_generated.deepcopy.go with:
//    				mapstructure.Decode(val, (*out)[key])
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              map[string]interface{} `json:"spec"`
	Status            ServiceStatus          `json:"status,omitempty"`
}

type ServiceSpec interface {
	DeepCopyServiceSpec() ServiceSpec
}

type ServiceStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NetworkList struct {
	Service
}

type Network struct {
	Service
}

type SubnetworkList struct {
	Service
}

type Subnetwork struct {
	Service
}

type Image struct {
	Service
}

type ImageList struct {
	Service
}
