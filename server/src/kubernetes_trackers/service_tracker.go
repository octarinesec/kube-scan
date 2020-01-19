package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
)

const ServiceTrackerKind TrackerKind = "Service"

type OctarineServiceTracker struct {
	Account string
	Domain  string
}

func (serviceTracker OctarineServiceTracker) GetKind() TrackerKind {
	return ServiceTrackerKind
}

func (serviceTracker OctarineServiceTracker) TrackService(service corev1.Service) *resources.ServiceResource {
	serviceResource := resources.NewServiceResource(serviceTracker.Account, serviceTracker.Domain, service.Namespace, service.Name)

	serviceResource.ServiceType = service.Spec.Type
	serviceResource.Selectors = service.Spec.Selector

	ports := make([]*resources.ServicePort, 0)
	for _, p := range service.Spec.Ports {
		sp := &resources.ServicePort{
			Port:     p.Port,
			NodePort: p.NodePort,
			Protocol: string(p.Protocol),
		}
		ports = append(ports, sp)
	}
	serviceResource.Ports = ports

	return serviceResource
}

func (serviceTracker OctarineServiceTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var service corev1.Service

	ok, _ := unmarshelResource(raw, &service)
	if ok {
		return serviceTracker.TrackService(service), nil
	}

	return nil, fmt.Errorf("error tracking service")
}

func (serviceTracker OctarineServiceTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	serviceResource := resources.NewServiceResource(serviceTracker.Account, serviceTracker.Domain, req.Namespace, req.Name)

	return serviceResource, nil
}
