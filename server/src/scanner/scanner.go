package scanner

import (
	"github.com/golang/glog"
	"kube-scan/risk"
	"kube-scan/state"
	"kube-scan/state_reader"
	"time"
)

var (
	stateReader *state_reader.ClusterStateReader
	riskFormula *risk.Formula

	ClusterState *state.Cluster
)

func InitScanner(refreshIntervalMinutes int, riskConfigFilePath string) error {
	var err error
	stateReader, err = state_reader.NewClusterStateReader()
	if err != nil {
		glog.Errorf("error creating cluster state reader: %v", err)
		return err
	}

	riskConfig := risk.NewConfigFromFile(riskConfigFilePath)
	riskFormula = risk.NewFormula(riskConfig)

	go initState(refreshIntervalMinutes)

	return nil
}

func initState(refreshIntervalMinutes int) {
	if err := readClusterState(); err != nil {
		glog.Fatalf("error refreshing cluster state: %v", err)
	}
	go refreshState(refreshIntervalMinutes)
}

func readClusterState() error {
	newClusterState, err := stateReader.ReadClusterState(ClusterState)
	if err != nil {
		return err
	}

	newClusterState.DeleteSystemData()
	newClusterState.AnalyzeCluster()
	newClusterState.CalculateRisk(riskFormula)
	ClusterState = newClusterState
	return nil
}

func refreshState(refreshIntervalMinutes int) {
	ticker := time.NewTicker(time.Duration(refreshIntervalMinutes) * time.Minute)
	for range ticker.C {
		if err := readClusterState(); err != nil {
			glog.Errorf("error refreshing cluster state: %v", err)
		}
	}
}
