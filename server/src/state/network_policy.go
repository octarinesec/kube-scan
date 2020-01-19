package state

import (
	"kube-scan/resources"
)

type NetworkPolicy struct {
	*resources.NetworkPolicyResource `bson:",inline"`
}

func NewNetworkPolicy(npResource *resources.NetworkPolicyResource) *NetworkPolicy {
	np := &NetworkPolicy{
		NetworkPolicyResource: npResource,
	}

	return np
}
