package resources

type ExecResource struct {
	BaseResource `bson:",inline"`
	Container    string
	Command      []string
}

func NewExecResource(account string, domain string, namespace string, name string, container string, command []string) *ExecResource {
	return &ExecResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Name:      name,
			Kind:      "Pod",
		},
		Container: container,
		Command:   command,
	}
}
