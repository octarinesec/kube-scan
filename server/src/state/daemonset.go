package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type Daemonset struct {
	*resources.DaemonsetResource `bson:",inline"`
	risk.WorkloadRisk            `bson:",inline"`
	Pod                          *Pod `json:"pod" bson:"pod"`
}

func (daemonset *Daemonset) GetPod() *Pod {
	return daemonset.Pod
}

func (daemonset *Daemonset) GetWorkloadPod() common.WorkloadPod {
	return daemonset.GetPod()
}

func (daemonset *Daemonset) AggregatePod(pod *Pod) {
	daemonsetPod := pod.Clone()
	daemonset.OctarineName = daemonsetPod.OctarineName
	daemonsetPod.Kind = daemonset.GetKind()
	daemonsetPod.Name = daemonset.GetName()
	daemonsetPod.OwnerReferenceKind = ""
	daemonsetPod.OwnerReferenceName = ""
	daemonsetPod.Labels = daemonset.Labels

	daemonset.Pod = daemonsetPod
}

func NewDaemonset(namespace *Namespace, daemonsetResource *resources.DaemonsetResource) *Daemonset {
	daemonset := &Daemonset{
		Pod:               NewPod(namespace, &daemonsetResource.PodResource),
		DaemonsetResource: daemonsetResource,
	}

	return daemonset
}
