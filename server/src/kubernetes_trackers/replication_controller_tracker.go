package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
)

const ReplicationControllerTrackerKind TrackerKind = "ReplicationController"

type OctarineReplicationControllerTracker struct {
	Account string
	Domain  string
}

func (rcTracker OctarineReplicationControllerTracker) GetKind() TrackerKind {
	return ReplicationControllerTrackerKind
}

func (rcTracker OctarineReplicationControllerTracker) TrackReplicationController(replicationController corev1.ReplicationController) *resources.ReplicationControllerResource {
	rcResource := resources.NewReplicationControllerResource(rcTracker.Account, rcTracker.Domain, replicationController.Namespace, replicationController.Name)

	rcResource.Labels = replicationController.Spec.Template.Labels
	rcResource.SecurityContext = newPodSecurityContext(replicationController.Spec.Template.Spec)
	rcResource.Containers = getContainers(replicationController.Spec.Template.Spec)
	rcResource.InitContainers = getInitContainers(replicationController.Spec.Template.Spec)
	rcResource.PodAnnotations = replicationController.Spec.Template.Annotations

	return rcResource
}

func (rcTracker OctarineReplicationControllerTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var rc corev1.ReplicationController

	ok, _ := unmarshelResource(raw, &rc)
	if ok {
		return rcTracker.TrackReplicationController(rc), nil
	}

	return nil, fmt.Errorf("error tracking replication controller")
}

func (rcTracker OctarineReplicationControllerTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	rcResource := resources.NewReplicationControllerResource(rcTracker.Account, rcTracker.Domain, req.Namespace, req.Name)

	return rcResource, nil
}
