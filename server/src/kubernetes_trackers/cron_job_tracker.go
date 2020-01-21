package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	batchv1 "k8s.io/api/batch/v1beta1"
	"kube-scan/resources"
)

const CronJobTrackerKind TrackerKind = "CronJob"

type OctarineCronJobTracker struct {
	Account string
	Domain  string
}

func (cronJobTracker OctarineCronJobTracker) GetKind() TrackerKind {
	return CronJobTrackerKind
}

func (cronJobTracker OctarineCronJobTracker) TrackCronJob(cronJob batchv1.CronJob) *resources.CronJobResource {
	cronJobResource := resources.NewCronJobResource(cronJobTracker.Account, cronJobTracker.Domain, cronJob.Namespace, cronJob.Name)

	cronJobResource.Labels = cronJob.Spec.JobTemplate.Spec.Template.Labels
	cronJobResource.SecurityContext = newPodSecurityContext(cronJob.Spec.JobTemplate.Spec.Template.Spec)
	cronJobResource.Containers = getContainers(cronJob.Spec.JobTemplate.Spec.Template.Spec)
	cronJobResource.InitContainers = getInitContainers(cronJob.Spec.JobTemplate.Spec.Template.Spec)
	cronJobResource.PodAnnotations = cronJob.Spec.JobTemplate.Annotations

	return cronJobResource
}

func (cronJobTracker OctarineCronJobTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var cronJob batchv1.CronJob

	ok, _ := unmarshelResource(raw, &cronJob)
	if ok {
		return cronJobTracker.TrackCronJob(cronJob), nil
	}

	return nil, fmt.Errorf("error tracking cronJob")
}

func (cronJobTracker OctarineCronJobTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	cronJobResource := resources.NewCronJobResource(cronJobTracker.Account, cronJobTracker.Domain, req.Namespace, req.Name)

	return cronJobResource, nil
}
