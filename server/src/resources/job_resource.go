package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobResource struct {
	PodResource        `bson:",inline"`
	OwnerReferenceKind string `json:"ownerReferenceKind,omitempty" bson:"ownerReferenceKind,omitempty"`
	OwnerReferenceName string `json:"ownerReferenceName,omitempty" bson:"ownerReferenceName,omitempty"`

	ComplitionTime *metav1.Time
}

func NewJobResource(account string, domain string, namespace string, name string) *JobResource {
	return &JobResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "Job",
				Name:      name,
			},
		},
	}
}
