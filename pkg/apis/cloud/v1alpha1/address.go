package v1alpha1

import (
	compute "google.golang.org/api/compute/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Instance `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AddressSpec   `json:"spec"`
	Status            AddressStatus `json:"status,omitempty"`
}

type AddressSpec struct {
	ProjectID string           `json:"projectID"`
	Payload   *compute.Address `json:"payload"`
}

type AddressStatus struct {
	Status string `json:"status"`
}
