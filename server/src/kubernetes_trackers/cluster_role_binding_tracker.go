package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"kube-scan/resources"
)

const ClusterRoleBindingTrackerKind TrackerKind = "ClusterRoleBinding"

type OctarineClusterRoleBindingTracker struct {
	Account string
	Domain  string
}

func (clusterRoleBindingTracker OctarineClusterRoleBindingTracker) GetKind() TrackerKind {
	return ClusterRoleBindingTrackerKind
}

func (clusterRoleBindingTracker OctarineClusterRoleBindingTracker) TrackClusterRoleBinding(clusterRoleBinding rbacv1.ClusterRoleBinding) *resources.ClusterRoleBindingResource {
	crbResource := resources.NewClusterRoleBindingResource(clusterRoleBindingTracker.Account, clusterRoleBindingTracker.Domain, clusterRoleBinding.Namespace, clusterRoleBinding.Name)

	crbResource.RoleRef = clusterRoleBinding.RoleRef.Name

	subjects := make([]*resources.RoleBindingSubjects, 0)
	for _, s := range clusterRoleBinding.Subjects {
		rbSubjects := &resources.RoleBindingSubjects{
			Kind:      s.Kind,
			APIGroup:  s.APIGroup,
			Name:      s.Name,
			Namespace: s.Namespace,
		}
		subjects = append(subjects, rbSubjects)
	}

	crbResource.Subjects = subjects

	return crbResource
}

func (clusterRoleBindingTracker OctarineClusterRoleBindingTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var clusterRoleBinding rbacv1.ClusterRoleBinding

	ok, _ := unmarshelResource(raw, &clusterRoleBinding)
	if ok {
		return clusterRoleBindingTracker.TrackClusterRoleBinding(clusterRoleBinding), nil
	}

	return nil, fmt.Errorf("error tracking role binding")
}

func (clusterRoleBindingTracker OctarineClusterRoleBindingTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	rbResource := resources.NewClusterRoleBindingResource(clusterRoleBindingTracker.Account, clusterRoleBindingTracker.Domain, req.Namespace, req.Name)

	return rbResource, nil
}
