package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	extv1 "k8s.io/api/extensions/v1beta1"
	"kube-scan/resources"
)

const IngressTrackerKind TrackerKind = "Ingress"

type OctarineIngressControllerTracker struct {
	Account string
	Domain  string
}

func (ingressTracker OctarineIngressControllerTracker) GetKind() TrackerKind {
	return IngressTrackerKind
}

func (ingressTracker OctarineIngressControllerTracker) TrackIngress(ingress extv1.Ingress) *resources.IngressControllerResource {
	ingressResource := resources.NewIngressControllerResource(ingressTracker.Account, ingressTracker.Domain, ingress.Namespace, ingress.Name)

	if ingress.Spec.Backend != nil {
		ingressResource.Backend = resources.NewIngressBackend(ingress.Spec.Backend.ServiceName)
	}

	rules := make([]*resources.IngressRule, 0)
	for _, r := range ingress.Spec.Rules {
		paths := make([]*resources.IngressPath, 0)
		for _, p := range r.HTTP.Paths {
			paths = append(paths, resources.NewIngressPath(p.Path, resources.NewIngressBackend(p.Backend.ServiceName)))
		}

		rules = append(rules, resources.NewIngressRule(r.Host, paths))
	}

	ingressResource.Rules = rules

	return ingressResource
}

func (ingressTracker OctarineIngressControllerTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var ingress extv1.Ingress

	ok, _ := unmarshelResource(raw, &ingress)
	if ok {
		return ingressTracker.TrackIngress(ingress), nil
	}

	return nil, fmt.Errorf("error tracking ingress controller")
}

func (ingressTracker OctarineIngressControllerTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	ingressResource := resources.NewIngressControllerResource(ingressTracker.Account, ingressTracker.Domain, req.Namespace, req.Name)

	return ingressResource, nil
}
