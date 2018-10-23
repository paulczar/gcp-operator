package stub

import (
	"github.com/paulczar/gcp-operator/pkg/apis/cloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getProjectID(meta metav1.ObjectMeta) string {
	a := meta.GetAnnotations()
	if val, ok := a["cloud.google.com/project-id"]; ok {
		return val
	}
	return ""
}

func setName(meta, spec string) string {
	if spec == "" {
		return meta
	} else {
		return spec
	}
}

func getStatus(stat, message, object string, err error) v1alpha1.ServiceStatus {
	if stat == "" {
		stat = "CREATED"
	}
	var s = v1alpha1.ServiceStatus{
		Status:  stat,
		Message: message,
		Object:  object,
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
