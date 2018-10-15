package stub

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func getProjectID(meta metav1.ObjectMeta) string {
	a := meta.GetAnnotations()
	if val, ok := a["cloud.google.com/project-id"]; ok {
		return val
	}
	return ""
}
