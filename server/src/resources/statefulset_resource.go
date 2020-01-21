package resources

type StatefulSetResource struct {
	PodResource `bson:",inline"`
}

func NewStatefulsetResource(account string, domain string, namespace string, name string) *StatefulSetResource {
	return &StatefulSetResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "StatefulSet",
				Name:      name,
			},
		},
	}
}
