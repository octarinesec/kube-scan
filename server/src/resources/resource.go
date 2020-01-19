package resources

import (
	"reflect"
)

type Resource interface {
	GetAccount() string
	GetDomain() string
	SetDomain(domain string)

	GetNamespace() string
	SetNamespace(namespace string)
	GetKind() string
	GetName() string
	SetName(name string)
	GetOctarineName() string
}

type BaseResource struct {
	Account      string `json:"account" bson:"account"`
	Domain       string `json:"domain" bson:"domain"`
	Namespace    string `json:"namespace" bson:"namespace"`
	Kind         string `json:"kind" bson:"kind"`
	Name         string `json:"name" bson:"name"`
	OctarineName string `json:"octarineName" bson:"octarineName"`
}

func (resource *BaseResource) GetAccount() string {
	return resource.Account
}

func (resource *BaseResource) GetDomain() string {
	return resource.Domain
}

func (resource *BaseResource) SetDomain(domain string) {
	resource.Domain = domain
}

func (resource *BaseResource) GetNamespace() string {
	return resource.Namespace
}

func (resource *BaseResource) SetNamespace(namespace string) {
	resource.Namespace = namespace
}

func (resource *BaseResource) GetKind() string {
	return resource.Kind
}

func (resource *BaseResource) GetName() string {
	return resource.Name
}

func (resource *BaseResource) SetName(name string) {
	resource.Name = name
}

func (resource *BaseResource) GetOctarineName() string {
	return resource.OctarineName
}

type LabeledResource interface {
	Resource

	GetLabels() map[string]string
}

type LabelsResource struct {
	Labels map[string]string `json:"labels,omitempty"`
}

func (labelResource *LabelsResource) GetLabels() map[string]string {
	return labelResource.Labels
}

func (labelResource *LabelsResource) Equals(other *LabelsResource) bool {
	return (len(labelResource.Labels) == 0 && len(other.Labels) == 0) ||
		reflect.DeepEqual(labelResource.Labels, other.Labels)
}

type HasSelectorsResource interface {
	Resource
	MatchLabels(selectors map[string]string) bool
}

type SelectorsResource struct {
	Selectors map[string]string `json:"selectors,omitempty" bson:"selectors,omitempty"`
}

type RoleBindingSubjects struct {
	Kind      string `json:"kind,omitempty" bson:"kind,omitempty"`
	APIGroup  string `json:"apiGroup,omitempty" bson:"apiGroup,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" bson:"namespace,omitempty"`
}

func ToSubjectNames(subjects []*RoleBindingSubjects) []string {
	result := make([]string, 0)

	for _, subject := range subjects {
		result = append(result, subject.Name)
	}

	return result
}
