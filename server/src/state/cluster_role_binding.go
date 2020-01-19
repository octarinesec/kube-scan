package state

import (
	"kube-scan/resources"
)

type ClusterRoleBinding struct {
	*resources.ClusterRoleBindingResource `bson:",inline"`
}

func NewClusterRoleBinding(crbResource *resources.ClusterRoleBindingResource) *ClusterRoleBinding {
	crb := &ClusterRoleBinding{
		ClusterRoleBindingResource: crbResource,
	}

	return crb
}
