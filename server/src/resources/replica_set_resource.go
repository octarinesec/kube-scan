package resources

type ReplicaSetResource struct {
	BaseResource       `bson:",inline"`
	OwnerReferenceKind string `json:"ownerReferenceKind,omitempty" bson:"ownerReferenceKind,omitempty"`
	OwnerReferenceName string `json:"ownerReferenceName,omitempty" bson:"ownerReferenceName,omitempty"`
}

func NewReplicaSetResource(account string, domain string, namespace string, name string) *ReplicaSetResource {
	return &ReplicaSetResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "ReplicaSet",
			Name:      name,
		},
	}
}
