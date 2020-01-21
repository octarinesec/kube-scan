package kubernetes_trackers

func GetTrackers(account string, domain string) []OctarineKindTracker {
	return []OctarineKindTracker{
		OctarineServiceTracker{Account: account, Domain: domain},
		OctarineDeploymentTracker{Account: account, Domain: domain},
		OctarineDaemonsetTracker{Account: account, Domain: domain},
		OctarineStatefulSetTracker{Account: account, Domain: domain},
		OctarineJobTracker{Account: account, Domain: domain},
		OctarineCronJobTracker{Account: account, Domain: domain},
		OctarineReplicaSetTracker{Account: account, Domain: domain},
		OctarineReplicationControllerTracker{Account: account, Domain: domain},
		OctarinePodTracker{Account: account, Domain: domain},
		OctarineNetworkPolicyTracker{Account: account, Domain: domain},
		OctarineRoleBindingTracker{Account: account, Domain: domain},
		OctarineClusterRoleBindingTracker{Account: account, Domain: domain},
		OctarineIngressControllerTracker{Account: account, Domain: domain},
		OctarineNamespaceTracker{Account: account, Domain: domain},
		OctarineExecTracker{Account: account, Domain: domain},
		OctarinePortForwardTracker{Account: account, Domain: domain},
	}
}

func GetKindToTrackerMap(account string, domain string) map[TrackerKind]OctarineKindTracker {
	kindToTrackerMap := make(map[TrackerKind]OctarineKindTracker)
	for _, tracker := range GetTrackers(account, domain) {
		kindToTrackerMap[tracker.GetKind()] = tracker
	}
	return kindToTrackerMap
}
