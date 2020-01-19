package resources

type RoleBindingResource struct {
	BaseResource `bson:",inline"`
	RoleRef      string                 `json:"role" bson:"role"`
	Subjects     []*RoleBindingSubjects `json:"subjects" bson:"subjects"`
}

func NewRoleBindingResource(account string, domain string, namespace string, name string) *RoleBindingResource {
	return &RoleBindingResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "RoleBinding",
			Name:      name,
		},
	}
}
