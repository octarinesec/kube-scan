package resources

type ReplicationControllerResource struct {
	PodResource `bson:",inline"`
}

func NewReplicationControllerResource(account string, domain string, namespace string, name string) *ReplicationControllerResource {
	return &ReplicationControllerResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "ReplicationController",
				Name:      name,
			},
		},
	}
}
