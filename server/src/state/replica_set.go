package state

import (
	"kube-scan/resources"
)

type ReplicaSet struct {
	*resources.ReplicaSetResource `bson:",inline"`
}

func NewReplicaSet(replicaSetResource *resources.ReplicaSetResource) *ReplicaSet {
	replicaSet := &ReplicaSet{
		ReplicaSetResource: replicaSetResource,
	}

	return replicaSet
}
