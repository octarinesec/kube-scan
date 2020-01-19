package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type ReplicationController struct {
	*resources.ReplicationControllerResource `bson:",inline"`
	risk.WorkloadRisk                        `bson:",inline"`
	Pod                                      *Pod `json:"pod" bson:"pod"`
}

func (replicationController *ReplicationController) GetPod() *Pod {
	return replicationController.Pod
}

func (replicationController *ReplicationController) GetWorkloadPod() common.WorkloadPod {
	return replicationController.GetPod()
}

func (replicationController *ReplicationController) AggregatePod(pod *Pod) {
	replicationControllerPod := pod.Clone()
	replicationController.OctarineName = replicationControllerPod.OctarineName
	replicationControllerPod.Kind = replicationController.GetKind()
	replicationControllerPod.Name = replicationController.GetName()
	replicationControllerPod.OwnerReferenceKind = ""
	replicationControllerPod.OwnerReferenceName = ""
	replicationControllerPod.Labels = replicationController.Labels

	replicationController.Pod = replicationControllerPod
}

func NewReplicationController(namespace *Namespace, replicationControllerResource *resources.ReplicationControllerResource) *ReplicationController {
	replicationController := &ReplicationController{
		Pod:                           NewPod(namespace, &replicationControllerResource.PodResource),
		ReplicationControllerResource: replicationControllerResource,
	}

	return replicationController
}
