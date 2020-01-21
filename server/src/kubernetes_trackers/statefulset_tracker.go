package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"kube-scan/resources"
)

const StatefulSetTrackerKind TrackerKind = "StatefulSet"

type OctarineStatefulSetTracker struct {
	Account string
	Domain  string
}

func (statefulSetTracker OctarineStatefulSetTracker) GetKind() TrackerKind {
	return StatefulSetTrackerKind
}

func (statefulSetTracker OctarineStatefulSetTracker) TrackStatefulset(statefulset appsv1.StatefulSet) *resources.StatefulSetResource {
	statefulsetResource := resources.NewStatefulsetResource(statefulSetTracker.Account, statefulSetTracker.Domain, statefulset.Namespace, statefulset.Name)

	statefulsetResource.Labels = statefulset.Spec.Template.Labels
	statefulsetResource.SecurityContext = newPodSecurityContext(statefulset.Spec.Template.Spec)
	statefulsetResource.Containers = getContainers(statefulset.Spec.Template.Spec)
	statefulsetResource.InitContainers = getInitContainers(statefulset.Spec.Template.Spec)
	statefulsetResource.PodAnnotations = statefulset.Spec.Template.Annotations

	return statefulsetResource
}

func (statefulSetTracker OctarineStatefulSetTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var statefulset appsv1.StatefulSet

	ok, _ := unmarshelResource(raw, &statefulset)
	if ok {
		return statefulSetTracker.TrackStatefulset(statefulset), nil
	}

	return nil, fmt.Errorf("error tracking statefulset")
}

func (statefulSetTracker OctarineStatefulSetTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	statefulsetResource := resources.NewStatefulsetResource(statefulSetTracker.Account, statefulSetTracker.Domain, req.Namespace, req.Name)

	return statefulsetResource, nil
}
