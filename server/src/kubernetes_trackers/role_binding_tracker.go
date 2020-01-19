package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"kube-scan/resources"
)

const RoleBindingTrackerKind TrackerKind = "RoleBinding"

type OctarineRoleBindingTracker struct {
	Account string
	Domain  string
}

func (roleBindingTracker OctarineRoleBindingTracker) GetKind() TrackerKind {
	return RoleBindingTrackerKind
}

func (roleBindingTracker OctarineRoleBindingTracker) TrackRoleBinding(roleBinding rbacv1.RoleBinding) *resources.RoleBindingResource {
	rbResource := resources.NewRoleBindingResource(roleBindingTracker.Account, roleBindingTracker.Domain, roleBinding.Namespace, roleBinding.Name)

	rbResource.RoleRef = roleBinding.RoleRef.Name

	subjects := make([]*resources.RoleBindingSubjects, 0)
	for _, s := range roleBinding.Subjects {
		rbSubjects := &resources.RoleBindingSubjects{
			Kind:      s.Kind,
			APIGroup:  s.APIGroup,
			Name:      s.Name,
			Namespace: s.Namespace,
		}
		subjects = append(subjects, rbSubjects)
	}

	rbResource.Subjects = subjects

	return rbResource
}

func (roleBindingTracker OctarineRoleBindingTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var roleBinding rbacv1.RoleBinding

	ok, _ := unmarshelResource(raw, &roleBinding)
	if ok {
		return roleBindingTracker.TrackRoleBinding(roleBinding), nil
	}

	return nil, fmt.Errorf("error tracking role binding")
}

func (roleBindingTracker OctarineRoleBindingTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	rbResource := resources.NewRoleBindingResource(roleBindingTracker.Account, roleBindingTracker.Domain, req.Namespace, req.Name)

	return rbResource, nil
}
