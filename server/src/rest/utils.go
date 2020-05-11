package rest

import (
	"github.com/toolkits/slice"
	"kube-scan/risk"
	"kube-scan/state"
)

type ClusterRiskData struct {
	Data risk.WorkloadRiskDataList `json:"data"`
}

func (w ClusterRiskData) Sanitized() interface{} {
	return w
}

func GetClusterRiskWorkloads(cluster *state.Cluster) risk.WorkloadRiskDataList {
	result := make([]*risk.WorkloadRiskData, 0)

	for _, namespace := range cluster.Namespaces {
		isSystemNamespace := slice.ContainsString(state.SystemNamespaces, namespace.Name)
		for _, workload := range namespace.GetAllRiskWorkloads() {
			result = append(result, risk.ToWorkloadRiskData(workload, isSystemNamespace))
		}
	}

	return result
}
