package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type CronJob struct {
	*resources.CronJobResource `bson:",inline"`
	risk.WorkloadRisk          `bson:",inline"`
	Pod                        *Pod `json:"pod" bson:"pod"`
}

func (cronJob *CronJob) GetPod() *Pod {
	return cronJob.Pod
}

func (cronJob *CronJob) GetWorkloadPod() common.WorkloadPod {
	return cronJob.GetPod()
}

func (cronJob *CronJob) AggregatePod(pod *Pod) {
	cronJobPod := pod.Clone()
	cronJob.OctarineName = cronJobPod.OctarineName
	cronJobPod.Kind = cronJob.GetKind()
	cronJobPod.Name = cronJob.GetName()
	cronJobPod.OwnerReferenceKind = ""
	cronJobPod.OwnerReferenceName = ""
	cronJobPod.Labels = cronJob.Labels

	cronJob.Pod = cronJobPod
}

func NewCronJob(namespace *Namespace, cronJobResource *resources.CronJobResource) *CronJob {
	cronJob := &CronJob{
		Pod:             NewPod(namespace, &cronJobResource.PodResource),
		CronJobResource: cronJobResource,
	}

	return cronJob
}
