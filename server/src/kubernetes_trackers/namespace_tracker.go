package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
)

const NamespaceTrackerKind TrackerKind = "Namespace"

type OctarineNamespaceTracker struct {
	Account string
	Domain  string
}

func (namespaceTracker OctarineNamespaceTracker) GetKind() TrackerKind {
	return NamespaceTrackerKind
}

func (namespaceTracker OctarineNamespaceTracker) TrackNamespace(namespace corev1.Namespace) *resources.NamespaceResource {
	nsResource := resources.NewNamespaceResource(namespaceTracker.Account, namespaceTracker.Domain, namespace.Name)
	return nsResource
}

func (namespaceTracker OctarineNamespaceTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var ns corev1.Namespace

	ok, _ := unmarshelResource(raw, &namespace)
	if ok {
		return namespaceTracker.TrackNamespace(ns), nil
	}

	return nil, fmt.Errorf("error tracking namespace")
}

func (namespaceTracker OctarineNamespaceTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	nsResource := resources.NewNamespaceResource(namespaceTracker.Account, namespaceTracker.Domain, req.Name)

	return nsResource, nil
}
