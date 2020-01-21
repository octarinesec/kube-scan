package resources

type ClusterRoleBindingResource struct {
	BaseResource `bson:",inline"`
	RoleRef      string                 `json:"role" bson:"role"`
	Subjects     []*RoleBindingSubjects `json:"subjects" bson:"role"`
}

func NewClusterRoleBindingResource(account string, domain string, namespace string, name string) *ClusterRoleBindingResource {
	return &ClusterRoleBindingResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "ClusterRoleBinding",
			Name:      name,
		},
	}
}
