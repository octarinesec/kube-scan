package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"kube-scan/resources"
)

const DaemonSetTrackerKind TrackerKind = "DaemonSet"

type OctarineDaemonsetTracker struct {
	Account string
	Domain  string
}

func (daemonsetTracker OctarineDaemonsetTracker) GetKind() TrackerKind {
	return DaemonSetTrackerKind
}

func (daemonsetTracker OctarineDaemonsetTracker) TrackDaemonset(daemonset appsv1.DaemonSet) *resources.DaemonsetResource {
	daemonsetResource := resources.NewDaemonsetResource(daemonsetTracker.Account, daemonsetTracker.Domain, daemonset.Namespace, daemonset.Name)

	daemonsetResource.Labels = daemonset.Spec.Template.Labels
	daemonsetResource.SecurityContext = newPodSecurityContext(daemonset.Spec.Template.Spec)
	daemonsetResource.Containers = getContainers(daemonset.Spec.Template.Spec)
	daemonsetResource.InitContainers = getInitContainers(daemonset.Spec.Template.Spec)
	daemonsetResource.PodAnnotations = daemonset.Spec.Template.Annotations

	return daemonsetResource
}

func (daemonsetTracker OctarineDaemonsetTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var daemonset appsv1.DaemonSet

	ok, _ := unmarshelResource(raw, &daemonset)
	if ok {
		return daemonsetTracker.TrackDaemonset(daemonset), nil
	}

	return nil, fmt.Errorf("error tracking daemonset")
}

func (daemonsetTracker OctarineDaemonsetTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	daemonsetResource := resources.NewDaemonsetResource(daemonsetTracker.Account, daemonsetTracker.Domain, req.Namespace, req.Name)

	return daemonsetResource, nil
}
