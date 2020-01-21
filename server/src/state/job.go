package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type Job struct {
	*resources.JobResource `bson:",inline"`
	risk.WorkloadRisk      `bson:",inline"`
	Pod                    *Pod `json:"pod" bson:"pod"`
}

func (job *Job) GetPod() *Pod {
	return job.Pod
}

func (job *Job) GetWorkloadPod() common.WorkloadPod {
	return job.GetPod()
}

func (job *Job) AggregatePod(pod *Pod) {
	jobPod := pod.Clone()
	job.OctarineName = jobPod.OctarineName
	jobPod.Kind = job.GetKind()
	jobPod.Name = job.GetName()
	jobPod.OwnerReferenceKind = ""
	jobPod.OwnerReferenceName = ""
	jobPod.Labels = job.Labels

	job.Pod = jobPod
}

func NewJob(namespace *Namespace, jobResource *resources.JobResource) *Job {
	job := &Job{
		Pod:         NewPod(namespace, &jobResource.PodResource),
		JobResource: jobResource,
	}

	return job
}
