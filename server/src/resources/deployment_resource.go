package resources

type DeploymentResource struct {
	PodResource `bson:",inline"`
	Replicas    *int32 `json:"replicas" bson:"replicas"`
}

func NewDeploymentResource(account string, domain string, namespace string, name string) *DeploymentResource {
	return &DeploymentResource{
		PodResource: PodResource{
			BaseResource: BaseResource{
				Account:   account,
				Domain:    domain,
				Namespace: namespace,
				Kind:      "Deployment",
				Name:      name,
			},
		},
	}
}
