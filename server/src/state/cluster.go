package state

import (
	"fmt"
	"github.com/golang/glog"
	"kube-scan/resources"
	"kube-scan/risk"
	"strings"
	"sync"
)

type Cluster struct {
	Name       string                `json:"name" bson:"name"`
	Namespaces map[string]*Namespace `json:"namespaces" bson:"namespaces"`

	ClusterRoleBindings map[string]*ClusterRoleBinding `json:"clusterRoleBindings" bson:"clusterRoleBindings"`

	mux sync.Mutex
}

var SystemNamespaces = []string{"octarine", "kube-system", "kube-public", "octarine-tiller", "istio-system", "octarine-dataplane", "kube-scan"}

func NewState(name string) *Cluster {
	return &Cluster{
		Name:       name,
		Namespaces: make(map[string]*Namespace),

		ClusterRoleBindings: make(map[string]*ClusterRoleBinding),
	}
}

func (cluster *Cluster) CalculateRiskWithRiskStatusFunc(formula *risk.Formula, statusFunc risk.GetStatusFunc) {
	for _, namespace := range cluster.Namespaces {
		for _, workload := range namespace.GetAllRiskWorkloads() {
			r, err := formula.CalculateRiskWithStatusGetter(workload, statusFunc)
			if err != nil {
				glog.Errorf("error calculating risk: %v", err)
			} else {
				workload.SetRisk(r)
			}
		}
	}
}

func (cluster *Cluster) CalculateRisk(formula *risk.Formula) {
	for _, namespace := range cluster.Namespaces {
		for _, workload := range namespace.GetAllRiskWorkloads() {
			r, err := formula.CalculateRisk(workload)
			if err != nil {
				glog.Errorf("error calculating risk: %v", err)
			} else {
				workload.SetRisk(r)
			}
		}
	}
}

func (cluster *Cluster) deleteEnvNamespaces() {
	for _, namespace := range cluster.Namespaces {
		for _, pod := range namespace.Pods {
			for _, containers := range pod.Containers {
				for name := range containers.Env {
					if strings.HasPrefix(name, "OCTARINE_") {
						delete(containers.Env, name)
					}
				}
			}
		}
	}
}

func (cluster *Cluster) deleteSystemPodsAndContainers() {
	for _, namespace := range cluster.Namespaces {
		//todo: move this to namespace
		for podName, pod := range namespace.Pods {
			for _, container := range pod.InitContainers {
				if strings.HasPrefix(container.Image, "octarinesec/idcontroller") || strings.HasPrefix(container.Image, "octarinesec/sidecar_injector") || strings.HasPrefix(container.Image, "octarinesec/istio_adapter") || strings.HasPrefix(container.Image, "octarinesec/octarine_ids") {
					delete(namespace.Pods, podName)
					continue
				}
			}

			pod.DeleteSystemContainers()
		}

		//todo: add function to get all owner resources
		for _, d := range namespace.Deployments {
			d.Pod.DeleteSystemContainers()
		}

		for _, d := range namespace.Daemonsets {
			d.Pod.DeleteSystemContainers()
		}

		for _, s := range namespace.StatefulSets {
			s.Pod.DeleteSystemContainers()
		}

		for _, j := range namespace.Jobs {
			j.Pod.DeleteSystemContainers()
		}

		for _, cj := range namespace.CronJobs {
			cj.Pod.DeleteSystemContainers()
		}

		for _, rc := range namespace.ReplicationControllers {
			rc.Pod.DeleteSystemContainers()
		}
	}
}

func (cluster *Cluster) DeleteSystemData() {
	cluster.deleteSystemPodsAndContainers()
	cluster.deleteEnvNamespaces()
}

func (cluster *Cluster) DeleteNonActiveResources() {
	for _, namespace := range cluster.Namespaces {
		namespace.DeleteNonActiveResources()
	}
}

func (cluster *Cluster) AnalyzeCluster() {
	for _, ns := range cluster.Namespaces {
		for _, pod := range ns.Pods {
			pod.Analyze(ns)
		}
		for _, deployment := range ns.Deployments {
			deployment.Pod.Analyze(ns)
		}
		for _, replicationController := range ns.ReplicationControllers {
			replicationController.Pod.Analyze(ns)
		}
		for _, daemonSet := range ns.Daemonsets {
			daemonSet.Pod.Analyze(ns)
		}
		for _, statefulSet := range ns.StatefulSets {
			statefulSet.Pod.Analyze(ns)
		}
		for _, job := range ns.Jobs {
			job.Pod.Analyze(ns)
		}
		for _, cronJob := range ns.CronJobs {
			cronJob.Pod.Analyze(ns)
		}
	}
}

func (cluster *Cluster) ensureNamespace(name string) *Namespace {
	if name == "" {
		return nil
	}

	namespace, ok := cluster.Namespaces[name]
	if !ok {
		namespace = NewNamespaceState(name)
		cluster.Namespaces[name] = namespace
	}

	return namespace
}

func (cluster *Cluster) GetNamespace(name string) (*Namespace, error) {
	if ns, ok := cluster.Namespaces[name]; ok {
		return ns, nil
	}

	return nil, fmt.Errorf("error getting namespace %v", name)
}

func (cluster *Cluster) updateResource(namespace *Namespace, resource resources.Resource) {
	switch resource.(type) {
	case *resources.ServiceResource:
		namespace.updateService(resource.(*resources.ServiceResource))
	case *resources.PodResource:
		namespace.updatePod(resource.(*resources.PodResource))
	case *resources.ReplicaSetResource:
		namespace.updateReplicaSet(resource.(*resources.ReplicaSetResource))
	case *resources.DeploymentResource:
		namespace.updateDeployment(resource.(*resources.DeploymentResource))
	case *resources.ReplicationControllerResource:
		namespace.updateReplicationController(resource.(*resources.ReplicationControllerResource))
	case *resources.DaemonsetResource:
		namespace.updateDaemonset(resource.(*resources.DaemonsetResource))
	case *resources.StatefulSetResource:
		namespace.updateStatefulSet(resource.(*resources.StatefulSetResource))
	case *resources.JobResource:
		namespace.updateJob(resource.(*resources.JobResource))
	case *resources.CronJobResource:
		namespace.updateCronJob(resource.(*resources.CronJobResource))
	case *resources.NetworkPolicyResource:
		namespace.updateNetworkPolicy(resource.(*resources.NetworkPolicyResource))
	case *resources.IngressControllerResource:
		namespace.updateIngressController(resource.(*resources.IngressControllerResource))
	case *resources.RoleBindingResource:
		namespace.updateRoleBinding(resource.(*resources.RoleBindingResource))
	case *resources.ClusterRoleBindingResource:
		cluster.updateClusterRoleBinding(resource.(*resources.ClusterRoleBindingResource))
	case *resources.NamespaceResource:
		cluster.updateNamespace(resource.(*resources.NamespaceResource))
	}
}

func (cluster *Cluster) updateNamespace(nsResource *resources.NamespaceResource) {
	cluster.ensureNamespace(nsResource.GetName())
}

func (cluster *Cluster) updateClusterRoleBinding(crbResource *resources.ClusterRoleBindingResource) {
	cluster.ClusterRoleBindings[crbResource.GetName()] = NewClusterRoleBinding(crbResource)
}

func (cluster *Cluster) Update(resource resources.Resource) {
	cluster.mux.Lock()
	defer cluster.mux.Unlock()
	namespace := cluster.ensureNamespace(resource.GetNamespace())
	cluster.updateResource(namespace, resource)
}
