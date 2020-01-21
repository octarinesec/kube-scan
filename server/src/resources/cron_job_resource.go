package resources

type CronJobResource struct {
	PodResource `bson:",inline"`
}

func NewCronJobResource(account string, domain string, namespace string, name string) *CronJobResource {
	return &CronJobResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "CronJob",
				Name:      name,
			},
		},
	}
}
