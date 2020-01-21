package kubernetes_trackers

import (
	"encoding/json"
	"github.com/golang/glog"
	"k8s.io/api/admission/v1beta1"
	"kube-scan/resources"
)

var Account string

type TrackerKind string

type OctarineKindTracker interface {
	GetKind() TrackerKind
	TrackResource(namespace string, name string, raw []byte) (resources.Resource, error)
	TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error)
}

func unmarshelResource(raw []byte, value interface{}) (bool, error) {
	if raw == nil {
		return false, nil
	}

	if err := json.Unmarshal(raw, &value); err != nil {
		glog.Errorf("Could not unmarshal raw object: %v", err)
		return false, err
	}

	return true, nil
}
