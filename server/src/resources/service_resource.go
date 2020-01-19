package resources

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"kube-scan/common"
	"strings"
)

type ServiceResource struct {
	BaseResource      `bson:",inline"`
	SelectorsResource `bson:",inline"`
	ServiceType       corev1.ServiceType `json:"serviceType,omitempty" bson:"serviceType,omitempty"`
	Ports             []*ServicePort     `json:"ports,omitempty" bson:"ports,omitempty"`
}

type ServicePort struct {
	Port     int32
	NodePort int32
	Protocol string
}

func (servicePort *ServicePort) String() string {
	return fmt.Sprintf("%v: (%v %v)", servicePort.Port, servicePort.NodePort, servicePort.Protocol)
}

func (service *ServiceResource) MatchLabels(podLabels map[string]string) bool {
	selector := labels.Set(service.Selectors).AsSelector()
	return selector.Matches(labels.Set(podLabels))
}

func (service *ServiceResource) GetResourceData() map[string]string {
	result := make(map[string]string)

	result["service-type"] = string(service.ServiceType)
	if len(service.Selectors) > 0 {
		result["selectors"] = common.LabelsAsString(service.Selectors)
	}

	portsStr := make([]string, 0)
	for _, p := range service.Ports {
		portsStr = append(portsStr, p.String())
	}

	if len(service.Ports) > 0 {
		result["service-ports"] = strings.Join(portsStr, ",")
	}

	return result
}

func (service *ServiceResource) IsExternal() bool {
	return service.IsLoadBalancer() || service.IsNodePort()
}

func (service *ServiceResource) IsLoadBalancer() bool {
	return service.ServiceType == corev1.ServiceTypeLoadBalancer
}

func (service *ServiceResource) IsNodePort() bool {
	return service.ServiceType == corev1.ServiceTypeNodePort
}

func NewServiceResource(account string, domain string, namespace string, name string) *ServiceResource {
	return &ServiceResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "Service",
			Name:      name,
		},
	}
}
