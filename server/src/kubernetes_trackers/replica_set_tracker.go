package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"kube-scan/resources"
)

const ReplicaSetTrackerKind TrackerKind = "ReplicaSet"

type OctarineReplicaSetTracker struct {
	Account string
	Domain  string
}

func (replicaSetTracker OctarineReplicaSetTracker) GetKind() TrackerKind {
	return ReplicaSetTrackerKind
}

func (replicaSetTracker OctarineReplicaSetTracker) TrackReplicaSet(replicaSet appsv1.ReplicaSet) *resources.ReplicaSetResource {
	replicaSetResource := resources.NewReplicaSetResource(replicaSetTracker.Account, replicaSetTracker.Domain, replicaSet.Namespace, replicaSet.Name)

	for _, or := range replicaSet.OwnerReferences {
		replicaSetResource.OwnerReferenceKind = or.Kind
		replicaSetResource.OwnerReferenceName = or.Name
	}

	return replicaSetResource
}

func (replicaSetTracker OctarineReplicaSetTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var replicaSet appsv1.ReplicaSet

	ok, _ := unmarshelResource(raw, &replicaSet)
	if ok {
		return replicaSetTracker.TrackReplicaSet(replicaSet), nil
	}

	return nil, fmt.Errorf("error tracking replicaSet")
}

func (replicaSetTracker OctarineReplicaSetTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	replicaSetResource := resources.NewReplicaSetResource(replicaSetTracker.Account, replicaSetTracker.Domain, req.Namespace, req.Name)

	return replicaSetResource, nil
}
