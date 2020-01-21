package resources

type NamespaceResource struct {
	BaseResource `bson:",inline"`
}

func NewNamespaceResource(account string, domain string, name string) *NamespaceResource {
	return &NamespaceResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: name,
			Kind:      "Namespace",
			Name:      name,
		},
	}
}
