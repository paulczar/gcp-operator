package v1alpha1

import (
	compute "google.golang.org/api/compute/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ForwardingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ForwardingRule `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ForwardingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ForwardingRuleSpec   `json:"spec"`
	Status            ForwardingRuleStatus `json:"status,omitempty"`
}

type ForwardingRuleSpec struct {
	*compute.ForwardingRule
}

type ForwardingRuleStatus struct {
	Status string `json:"status"`
}
