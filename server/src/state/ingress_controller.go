package state

import (
	"kube-scan/resources"
)

type IngressController struct {
	*resources.IngressControllerResource `bson:",inline"`
}

func NewIngressController(ingressResource *resources.IngressControllerResource) *IngressController {
	ingress := &IngressController{
		IngressControllerResource: ingressResource,
	}

	return ingress
}
