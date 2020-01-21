package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type StatefulSet struct {
	*resources.StatefulSetResource `bson:",inline"`
	risk.WorkloadRisk              `bson:",inline"`
	Pod                            *Pod `json:"pod" bson:"pod"`
}

func (statefulSet *StatefulSet) GetPod() *Pod {
	return statefulSet.Pod
}

func (statefulSet *StatefulSet) GetWorkloadPod() common.WorkloadPod {
	return statefulSet.GetPod()
}

func (statefulSet *StatefulSet) AggregatePod(pod *Pod) {
	statefulSetPod := pod.Clone()
	statefulSet.OctarineName = statefulSetPod.OctarineName
	statefulSetPod.Kind = statefulSet.GetKind()
	statefulSetPod.Name = statefulSet.GetName()
	statefulSetPod.OwnerReferenceKind = ""
	statefulSetPod.OwnerReferenceName = ""
	statefulSetPod.Labels = statefulSet.Labels

	statefulSet.Pod = statefulSetPod
}

func NewStatefulSet(namespace *Namespace, statefulSetResource *resources.StatefulSetResource) *StatefulSet {
	statefulSet := &StatefulSet{
		Pod:                 NewPod(namespace, &statefulSetResource.PodResource),
		StatefulSetResource: statefulSetResource,
	}

	return statefulSet
}
