package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	networkingV1 "k8s.io/api/networking/v1"
	"kube-scan/resources"
)

const NetworkPolicyTrackerKind TrackerKind = "NetworkPolicy"

type OctarineNetworkPolicyTracker struct {
	Account string
	Domain  string
}

func (npTracker OctarineNetworkPolicyTracker) GetKind() TrackerKind {
	return NetworkPolicyTrackerKind
}

func (npTracker OctarineNetworkPolicyTracker) TrackNetworkPolicy(networkPolicy networkingV1.NetworkPolicy) *resources.NetworkPolicyResource {
	npResource := resources.NewNetworkPolicyResource(npTracker.Account, npTracker.Domain, networkPolicy.Namespace, networkPolicy.Name)

	npResource.Selectors = networkPolicy.Spec.PodSelector.MatchLabels
	npResource.MatchExpressions = networkPolicy.Spec.PodSelector.MatchExpressions
	npResource.PolicyTypes = networkPolicy.Spec.PolicyTypes
	npResource.Ingress = networkPolicy.Spec.Ingress
	npResource.Egress = networkPolicy.Spec.Egress

	return npResource
}

func (npTracker OctarineNetworkPolicyTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var networkPolicy networkingV1.NetworkPolicy

	ok, _ := unmarshelResource(raw, &networkPolicy)
	if ok {
		return npTracker.TrackNetworkPolicy(networkPolicy), nil
	}

	return nil, fmt.Errorf("error tracking network policy")
}

func (npTracker OctarineNetworkPolicyTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	npResource := resources.NewNetworkPolicyResource(npTracker.Account, npTracker.Domain, req.Namespace, req.Name)

	return npResource, nil
}
