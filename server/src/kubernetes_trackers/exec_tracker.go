package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
)

const ExecTrackerKind TrackerKind = "PodExecOptions"

type OctarineExecTracker struct {
	Account string
	Domain  string
}

func (execTracker OctarineExecTracker) GetKind() TrackerKind {
	return ExecTrackerKind
}

func (execTracker OctarineExecTracker) TrackExec(namespace string, name string, exec corev1.PodExecOptions) *resources.ExecResource {
	return resources.NewExecResource(execTracker.Account, execTracker.Domain, namespace, name, exec.Container, exec.Command)
}

func (execTracker OctarineExecTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var podExec corev1.PodExecOptions
	ok, _ := unmarshelResource(raw, &podExec)
	if ok {
		return execTracker.TrackExec(namespace, name, podExec), nil
	}
	return nil, fmt.Errorf("error tracking service")
}

func (execTracker OctarineExecTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	return nil, fmt.Errorf("exec/cp are not a delete operations")
}
