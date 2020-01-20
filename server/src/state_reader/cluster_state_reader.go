package state_reader

import (
	"github.com/golang/glog"
	"golang.org/x/sync/errgroup"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	kt "kube-scan/kubernetes_trackers"
	"kube-scan/state"
)

type ClusterStateReader struct {
	kindToTrackerMap map[kt.TrackerKind]kt.OctarineKindTracker
	kubeClient       *kubernetes.Clientset

	serviceTracker               kt.OctarineServiceTracker
	deploymentTracker            kt.OctarineDeploymentTracker
	daemonsetTracker             kt.OctarineDaemonsetTracker
	statefulSetTracker           kt.OctarineStatefulSetTracker
	jobTracker                   kt.OctarineJobTracker
	cronJobTracker               kt.OctarineCronJobTracker
	replicaSetTracker            kt.OctarineReplicaSetTracker
	replicationControllerTracker kt.OctarineReplicationControllerTracker
	podTracker                   kt.OctarinePodTracker
	networkPolicyTracker         kt.OctarineNetworkPolicyTracker
	roleBindingTracker           kt.OctarineRoleBindingTracker
	clusterRoleBindingTracker    kt.OctarineClusterRoleBindingTracker
	ingressControllerTracker     kt.OctarineIngressControllerTracker
}

func NewClusterStateReader() (*ClusterStateReader, error) {
	kindToTrackerMap := kt.GetKindToTrackerMap("", "")

	//#################################
	// use this in order to run from your local machine against your kubectl context
	//#################################
	
	//kubeconfigpath := filepath.Join(
	//	os.Getenv("HOME"), ".kube", "config",
	//)
	//kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	//if err != nil {
	//	glog.Fatal("Error get configs")
	//}

	kubeConfig, err := restclient.InClusterConfig()
	if err != nil {
		glog.Errorf("error getting kube configs from InClusterConfig")
		return nil, err
	}

	client, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		glog.Errorf("error getting creating Kubernetes client")
		return nil, err
	}

	reader := &ClusterStateReader{
		kindToTrackerMap: kindToTrackerMap,
		kubeClient:       client,

		serviceTracker:               kindToTrackerMap[kt.ServiceTrackerKind].(kt.OctarineServiceTracker),
		podTracker:                   kindToTrackerMap[kt.PodTrackerKind].(kt.OctarinePodTracker),
		deploymentTracker:            kindToTrackerMap[kt.DeploymentTrackerKind].(kt.OctarineDeploymentTracker),
		daemonsetTracker:             kindToTrackerMap[kt.DaemonSetTrackerKind].(kt.OctarineDaemonsetTracker),
		statefulSetTracker:           kindToTrackerMap[kt.StatefulSetTrackerKind].(kt.OctarineStatefulSetTracker),
		jobTracker:                   kindToTrackerMap[kt.JobTrackerKind].(kt.OctarineJobTracker),
		cronJobTracker:               kindToTrackerMap[kt.CronJobTrackerKind].(kt.OctarineCronJobTracker),
		replicaSetTracker:            kindToTrackerMap[kt.ReplicaSetTrackerKind].(kt.OctarineReplicaSetTracker),
		replicationControllerTracker: kindToTrackerMap[kt.ReplicationControllerTrackerKind].(kt.OctarineReplicationControllerTracker),
		networkPolicyTracker:         kindToTrackerMap[kt.NetworkPolicyTrackerKind].(kt.OctarineNetworkPolicyTracker),
		ingressControllerTracker:     kindToTrackerMap[kt.IngressTrackerKind].(kt.OctarineIngressControllerTracker),
		roleBindingTracker:           kindToTrackerMap[kt.RoleBindingTrackerKind].(kt.OctarineRoleBindingTracker),
		clusterRoleBindingTracker:    kindToTrackerMap[kt.ClusterRoleBindingTrackerKind].(kt.OctarineClusterRoleBindingTracker),
	}

	return reader, nil
}

func (reader *ClusterStateReader) ReadClusterState(prevState *state.Cluster) (*state.Cluster, error) {
	clusterState := state.NewState("current_state")

	namespaces, err := reader.kubeClient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var eg errgroup.Group
	for _, namespace := range namespaces.Items {
		namespace := namespace // https://golang.org/doc/faq#closures_and_goroutines
		eg.Go(func() error {
			return reader.readNsState(clusterState, prevState, namespace)
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	clusterRoleBindings, err := reader.kubeClient.RbacV1().ClusterRoleBindings().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, clusterRoleBinding := range clusterRoleBindings.Items {
		clusterRoleBinding.Kind = "ClusterRoleBinding"
		clusterRoleBinding.APIVersion = "rbac/v1"
		clusterState.Update(reader.clusterRoleBindingTracker.TrackClusterRoleBinding(clusterRoleBinding))
	}

	return clusterState, nil
}

func (reader *ClusterStateReader) readNsState(clusterState *state.Cluster, prevState *state.Cluster, namespace v1.Namespace) error {
	deployments, err := reader.kubeClient.AppsV1().Deployments(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, deployment := range deployments.Items {
		deployment.Kind = "Deployment"
		deployment.APIVersion = "apps/v1"
		clusterState.Update(reader.deploymentTracker.TrackDeployment(deployment))
	}
	replicaSets, err := reader.kubeClient.AppsV1().ReplicaSets(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, replicaSet := range replicaSets.Items {
		replicaSet.Kind = "ReplicaSet"
		replicaSet.APIVersion = "apps/v1"
		clusterState.Update(reader.replicaSetTracker.TrackReplicaSet(replicaSet))
	}
	replicationControllers, err := reader.kubeClient.CoreV1().ReplicationControllers(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, rc := range replicationControllers.Items {
		rc.Kind = "ReplicationController"
		rc.APIVersion = "core/v1"
		clusterState.Update(reader.replicationControllerTracker.TrackReplicationController(rc))
	}
	statefulSets, err := reader.kubeClient.AppsV1().StatefulSets(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, statefulSet := range statefulSets.Items {
		statefulSet.Kind = "StatefulSet"
		statefulSet.APIVersion = "apps/v1"
		clusterState.Update(reader.statefulSetTracker.TrackStatefulset(statefulSet))
	}
	daemonSets, err := reader.kubeClient.AppsV1().DaemonSets(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, daemonSet := range daemonSets.Items {
		daemonSet.Kind = "DaemonSet"
		daemonSet.APIVersion = "apps/v1"
		clusterState.Update(reader.daemonsetTracker.TrackDaemonset(daemonSet))
	}
	jobs, err := reader.kubeClient.BatchV1().Jobs(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, job := range jobs.Items {
		if job.Status.CompletionTime != nil {
			continue
		}
		job.Kind = "Job"
		job.APIVersion = "batch/v1"
		clusterState.Update(reader.jobTracker.TrackJob(job))
	}
	cronJobs, err := reader.kubeClient.BatchV1beta1().CronJobs(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, cronJob := range cronJobs.Items {
		cronJob.Kind = "CronJob"
		cronJob.APIVersion = "batch/v1beta1"
		clusterState.Update(reader.cronJobTracker.TrackCronJob(cronJob))
	}
	services, err := reader.kubeClient.CoreV1().Services(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, service := range services.Items {
		service.Kind = "Service"
		service.APIVersion = "core/v1"
		clusterState.Update(reader.serviceTracker.TrackService(service))
	}
	networkPolicies, err := reader.kubeClient.NetworkingV1().NetworkPolicies(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, networkPolicy := range networkPolicies.Items {
		networkPolicy.Kind = "NetworkPolicy"
		networkPolicy.APIVersion = "networking/v1"
		clusterState.Update(reader.networkPolicyTracker.TrackNetworkPolicy(networkPolicy))
	}
	ingresses, err := reader.kubeClient.ExtensionsV1beta1().Ingresses(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, ingress := range ingresses.Items {
		ingress.Kind = "Ingress"
		ingress.APIVersion = "extensions/v1beta1"
		clusterState.Update(reader.ingressControllerTracker.TrackIngress(ingress))
	}
	pods, err := reader.kubeClient.CoreV1().Pods(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, pod := range pods.Items {
		if prevState != nil {
			if ns, nsOk := prevState.Namespaces[namespace.Name]; nsOk {
				if curPod, podOk := ns.Pods[pod.Name]; podOk && curPod.IsDeleted {
					continue
				}
			}
		}

		pod.Kind = "Pod"
		pod.APIVersion = "core/v1"
		clusterState.Update(reader.podTracker.TrackPod(pod))
	}
	roleBindings, err := reader.kubeClient.RbacV1().RoleBindings(namespace.Name).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, roleBinding := range roleBindings.Items {
		roleBinding.Kind = "RoleBinding"
		roleBinding.APIVersion = "rbac/v1"
		clusterState.Update(reader.roleBindingTracker.TrackRoleBinding(roleBinding))
	}
	return nil
}
