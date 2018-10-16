package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TargetPool `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TargetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TargetPoolSpec   `json:"spec"`
	Status            TargetPoolStatus `json:"status,omitempty"`
}

type TargetPoolSpec struct {
	//*compute.TargetPool
	Resource map[string]interface{} `json:"resource"`
	//Resource *compute.TargetPool `json:"resource"`
}

type TargetPoolStatus struct {
	Status string `json:"status"`
}
