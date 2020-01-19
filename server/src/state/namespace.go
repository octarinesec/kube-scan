package state

import (
	"kube-scan/resources"
	"kube-scan/risk"
)

type Namespace struct {
	Name string `json:"name" bson:"name"`

	Services        map[string]*Service           `json:"services" bson:"services"`
	NetworkPolicies map[string]*NetworkPolicy     `json:"networkPolicies" bson:"networkPolicies"`
	Ingress         map[string]*IngressController `json:"ingress" bson:"ingress"`
	ReplicaSets     map[string]*ReplicaSet        `json:"replicaSets" bson:"replicaSets"`
	Pods            map[string]*Pod               `json:"pods" bson:"pods"`

	Deployments            map[string]*Deployment            `json:"deployments" bson:"deployments"`
	ReplicationControllers map[string]*ReplicationController `json:"replicationControllers" bson:"replicationControllers"`
	Daemonsets             map[string]*Daemonset             `json:"daemonsets" bson:"daemonsets"`
	StatefulSets           map[string]*StatefulSet           `json:"statefulsets" bson:"statefulsets"`
	Jobs                   map[string]*Job                   `json:"jobs" bson:"jobs"`
	CronJobs               map[string]*CronJob               `json:"cronJobs" bson:"cronJobs"`

	RoleBindings map[string]*RoleBinding `json:"roleBindings" bson:"roleBindings"`
}

type OwnerResource interface {
	AggregatePod(pod *Pod)
	GetPod() *Pod
}

func (namespace *Namespace) GetAllRiskWorkloads() []risk.IWorloadRisk {
	workloads := make([]risk.IWorloadRisk, 0)

	for _, deployment := range namespace.Deployments {
		workloads = append(workloads, deployment)
	}

	for _, daemonset := range namespace.Daemonsets {
		workloads = append(workloads, daemonset)
	}

	for _, cronJob := range namespace.CronJobs {
		workloads = append(workloads, cronJob)
	}

	for _, rc := range namespace.ReplicationControllers {
		workloads = append(workloads, rc)
	}

	for _, statefulset := range namespace.StatefulSets {
		workloads = append(workloads, statefulset)
	}

	for _, job := range namespace.Jobs {
		if job.OwnerReferenceName == "" && job.OwnerReferenceKind == "" {
			workloads = append(workloads, job)
		}
	}

	for _, pod := range namespace.Pods {
		if pod.OwnerReferenceName == "" && pod.OwnerReferenceKind == "" {
			workloads = append(workloads, pod)
		}
	}

	return workloads
}

func (namespace *Namespace) DeleteNonActiveResources() {
	for podName, pod := range namespace.Pods {
		if pod.IsDeleted {
			delete(namespace.Pods, podName)
		}
	}
}

func (namespace *Namespace) updateService(serviceResource *resources.ServiceResource) {
	namespace.Services[serviceResource.GetName()] = NewService(serviceResource)
}

func (namespace *Namespace) updateIngressController(ingressResource *resources.IngressControllerResource) {
	namespace.Ingress[ingressResource.GetName()] = NewIngressController(ingressResource)
}

func (namespace *Namespace) updateNetworkPolicy(npResource *resources.NetworkPolicyResource) {
	namespace.NetworkPolicies[npResource.GetName()] = NewNetworkPolicy(npResource)
}

func (namespace *Namespace) updatePod(podResource *resources.PodResource) {
	newPod := NewPod(namespace, podResource)

	hasOr, or := namespace.GetPodOwnerReference(newPod.OwnerReferenceKind, newPod.OwnerReferenceName)
	if hasOr {
		or.AggregatePod(newPod)
	} else {
		namespace.Pods[newPod.GetName()] = newPod
	}
}

func (namespace *Namespace) updateDeployment(deploymentResource *resources.DeploymentResource) {
	namespace.Deployments[deploymentResource.GetName()] = NewDeployment(namespace, deploymentResource)
}

func (namespace *Namespace) updateDaemonset(daemonsetResource *resources.DaemonsetResource) {
	namespace.Daemonsets[daemonsetResource.GetName()] = NewDaemonset(namespace, daemonsetResource)
}

func (namespace *Namespace) updateStatefulSet(statefulSetResource *resources.StatefulSetResource) {
	namespace.StatefulSets[statefulSetResource.GetName()] = NewStatefulSet(namespace, statefulSetResource)
}

func (namespace *Namespace) updateReplicationController(replicationControllerResource *resources.ReplicationControllerResource) {
	namespace.ReplicationControllers[replicationControllerResource.GetName()] = NewReplicationController(namespace, replicationControllerResource)
}

func (namespace *Namespace) updateCronJob(cronJobResource *resources.CronJobResource) {
	namespace.CronJobs[cronJobResource.GetName()] = NewCronJob(namespace, cronJobResource)
}

func (namespace *Namespace) updateJob(jobResource *resources.JobResource) {
	namespace.Jobs[jobResource.GetName()] = NewJob(namespace, jobResource)
}

func (namespace *Namespace) updateReplicaSet(replicaSetResource *resources.ReplicaSetResource) {
	namespace.ReplicaSets[replicaSetResource.Name] = NewReplicaSet(replicaSetResource)
}

func (namespace *Namespace) updateRoleBinding(rbResource *resources.RoleBindingResource) {
	namespace.RoleBindings[rbResource.GetName()] = NewRoleBinding(rbResource)
}

func (namespace *Namespace) GetPodOwnerReference(kind string, name string) (bool, OwnerResource) {
	switch kind {
	case "ReplicaSet":
		for rsName, rs := range namespace.ReplicaSets {
			if rsName == name {
				return namespace.GetPodOwnerReference(rs.OwnerReferenceKind, rs.OwnerReferenceName)
			}
		}
	case "Deployment":
		for dName, d := range namespace.Deployments {
			if dName == name {
				return true, d
			}
		}
	case "ReplicationController":
		for rcName, rc := range namespace.ReplicationControllers {
			if rcName == name {
				return true, rc
			}
		}
	case "DaemonSet":
		for dName, d := range namespace.Daemonsets {
			if dName == name {
				return true, d
			}
		}
	case "StatefulSet":
		for dName, d := range namespace.StatefulSets {
			if dName == name {
				return true, d
			}
		}
	case "Job":
		for jobName, job := range namespace.Jobs {
			if jobName == name {
				hasOr, or := namespace.GetPodOwnerReference(job.OwnerReferenceKind, job.OwnerReferenceName)
				if hasOr {
					return true, or
				}

				return true, job
			}
		}
	case "CronJob":
		for dName, d := range namespace.CronJobs {
			if dName == name {
				return true, d
			}
		}
	}

	return false, nil
}

func NewNamespaceState(name string /*, withPending bool*/) *Namespace {
	return &Namespace{
		Name: name,

		Services:        make(map[string]*Service),
		NetworkPolicies: make(map[string]*NetworkPolicy),
		Ingress:         make(map[string]*IngressController),
		Pods:            make(map[string]*Pod),
		ReplicaSets:     make(map[string]*ReplicaSet),

		Deployments:            make(map[string]*Deployment),
		ReplicationControllers: make(map[string]*ReplicationController),
		Daemonsets:             make(map[string]*Daemonset),
		StatefulSets:           make(map[string]*StatefulSet),
		Jobs:                   make(map[string]*Job),
		CronJobs:               make(map[string]*CronJob),

		RoleBindings: make(map[string]*RoleBinding),
	}
}
