package state

import (
	"kube-scan/resources"
)

type Service struct {
	*resources.ServiceResource `bson:",inline"`
}

func NewService(serviceResource *resources.ServiceResource) *Service {
	service := &Service{
		ServiceResource: serviceResource,
	}

	return service
}
