package resources

type CustomResource struct {
	BaseResource `bson:",inline"`
}

func NewCustomResource(namespace string, kind string, name string, account string, domain string) *CustomResource {
	return &CustomResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      kind,
			Name:      name,
		},
	}
}
