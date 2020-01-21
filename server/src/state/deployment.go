package state

import (
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
)

type Deployment struct {
	*resources.DeploymentResource `bson:",inline"`
	risk.WorkloadRisk             `bson:",inline"`
	Pod                           *Pod     `json:"pod" bson:"pod"`
	ReplicaSets                   []string `json:"replicaSets" bson:"replicaSets"`
}

func (deployment *Deployment) GetPod() *Pod {
	return deployment.Pod
}

func (deployment *Deployment) GetWorkloadPod() common.WorkloadPod {
	return deployment.GetPod()
}

func (deployment *Deployment) AggregatePod(pod *Pod) {
	deploymentPod := pod.Clone()
	deploymentPod.SetParentName(deployment.GetName())
	deployment.OctarineName = deploymentPod.OctarineName
	deploymentPod.Kind = deployment.GetKind()
	deploymentPod.Name = deployment.GetName()
	deploymentPod.OwnerReferenceKind = ""
	deploymentPod.OwnerReferenceName = ""
	deploymentPod.Labels = deployment.Labels

	deployment.Pod = deploymentPod
}

func NewDeployment(namespace *Namespace, deploymentResource *resources.DeploymentResource) *Deployment {
	pod := NewPod(namespace, &deploymentResource.PodResource)
	pod.SetParentName(deploymentResource.Name)
	deployment := &Deployment{
		ReplicaSets:        make([]string, 0),
		Pod:                pod,
		DeploymentResource: deploymentResource,
	}

	return deployment
}
