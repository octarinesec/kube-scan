package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"kube-scan/resources"
)

const DeploymentTrackerKind TrackerKind = "Deployment"

type OctarineDeploymentTracker struct {
	Account string
	Domain  string
}

func (deploymentTracker OctarineDeploymentTracker) GetKind() TrackerKind {
	return DeploymentTrackerKind
}

func (deploymentTracker OctarineDeploymentTracker) TrackDeployment(deployment appsv1.Deployment) *resources.DeploymentResource {
	deploymentResource := resources.NewDeploymentResource(deploymentTracker.Account, deploymentTracker.Domain, deployment.Namespace, deployment.Name)

	deploymentResource.Replicas = deployment.Spec.Replicas
	deploymentResource.Labels = deployment.Spec.Template.Labels
	deploymentResource.SecurityContext = newPodSecurityContext(deployment.Spec.Template.Spec)
	deploymentResource.Containers = getContainers(deployment.Spec.Template.Spec)
	deploymentResource.InitContainers = getInitContainers(deployment.Spec.Template.Spec)
	deploymentResource.PodAnnotations = deployment.Spec.Template.Annotations

	return deploymentResource
}

func (deploymentTracker OctarineDeploymentTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var deployment appsv1.Deployment

	ok, _ := unmarshelResource(raw, &deployment)
	if ok {
		return deploymentTracker.TrackDeployment(deployment), nil
	}

	return nil, fmt.Errorf("error tracking deployment")
}

func (deploymentTracker OctarineDeploymentTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	deploymentResource := resources.NewDeploymentResource(deploymentTracker.Account, deploymentTracker.Domain, req.Namespace, req.Name)

	return deploymentResource, nil
}
