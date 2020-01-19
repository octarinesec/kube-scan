package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	batchv1 "k8s.io/api/batch/v1"
	"kube-scan/resources"
)

const JobTrackerKind TrackerKind = "Job"

type OctarineJobTracker struct {
	Account string
	Domain  string
}

func (jobTracker OctarineJobTracker) GetKind() TrackerKind {
	return JobTrackerKind
}

func (jobTracker OctarineJobTracker) TrackJob(job batchv1.Job) *resources.JobResource {
	jobResource := resources.NewJobResource(jobTracker.Account, jobTracker.Domain, job.Namespace, job.Name)

	for _, or := range job.OwnerReferences {
		jobResource.OwnerReferenceKind = or.Kind
		jobResource.OwnerReferenceName = or.Name
	}

	jobResource.Labels = job.Spec.Template.Labels
	jobResource.SecurityContext = newPodSecurityContext(job.Spec.Template.Spec)
	jobResource.Containers = getContainers(job.Spec.Template.Spec)
	jobResource.InitContainers = getInitContainers(job.Spec.Template.Spec)
	jobResource.PodAnnotations = job.Spec.Template.Annotations

	jobResource.ComplitionTime = job.Status.CompletionTime

	return jobResource
}

func (jobTracker OctarineJobTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var job batchv1.Job

	ok, _ := unmarshelResource(raw, &job)
	if ok {
		return jobTracker.TrackJob(job), nil
	}

	return nil, fmt.Errorf("error tracking job")
}

func (jobTracker OctarineJobTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	jobResource := resources.NewJobResource(jobTracker.Account, jobTracker.Domain, req.Namespace, req.Name)

	return jobResource, nil
}
