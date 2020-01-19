package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
)

const PortForwardTrackerKind TrackerKind = "PodPortForwardOptions"

type OctarinePortForwardTracker struct {
	Account string
	Domain  string
}

func (portForwardTracker OctarinePortForwardTracker) GetKind() TrackerKind {
	return PortForwardTrackerKind
}

func (portForwardTracker OctarinePortForwardTracker) TrackPortForward(namespace string, name string, portForward corev1.PodPortForwardOptions) *resources.PortForwardResource {
	return resources.NewPortForwardResource(portForwardTracker.Account, portForwardTracker.Domain, namespace, name, portForward.Ports)
}

func (portForwardTracker OctarinePortForwardTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var portForward corev1.PodPortForwardOptions
	ok, _ := unmarshelResource(raw, &portForward)
	if ok {
		return portForwardTracker.TrackPortForward(namespace, name, portForward), nil
	}
	return nil, fmt.Errorf("error tracking service")
}

func (portForwardTracker OctarinePortForwardTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	return nil, fmt.Errorf("port forward is not a delete operation")
}
