package state

import (
	"kube-scan/resources"
)

type RoleBinding struct {
	*resources.RoleBindingResource `bson:",inline"`
}

func NewRoleBinding(rbResource *resources.RoleBindingResource) *RoleBinding {
	roleBinding := &RoleBinding{
		RoleBindingResource: rbResource,
	}

	return roleBinding
}
