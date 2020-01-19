package resources

type DaemonsetResource struct {
	PodResource `bson:",inline"`
}

func NewDaemonsetResource(account string, domain string, namespace string, name string) *DaemonsetResource {
	return &DaemonsetResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "DaemonSet",
				Name:      name,
			},
		},
	}
}
