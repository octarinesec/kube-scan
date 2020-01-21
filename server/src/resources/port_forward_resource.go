package resources

type PortForwardResource struct {
	BaseResource `bson:",inline"`
	Ports        []int32
}

func NewPortForwardResource(account string, domain string, namespace string, name string, ports []int32) *PortForwardResource {
	return &PortForwardResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Name:      name,
			Kind:      "Pod",
		},
		Ports: ports,
	}
}
