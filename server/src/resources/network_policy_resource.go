package resources

import (
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"kube-scan/common"
)

type NetworkPolicyResource struct {
	BaseResource      `bson:",inline"`
	SelectorsResource `bson:",inline"`
	MatchExpressions  []metav1.LabelSelectorRequirement       `json:"matchExpressions,omitempty" bson:"matchExpressions,omitempty"`
	Ingress           []networkingv1.NetworkPolicyIngressRule `json:"ingress,omitempty" bson:"ingress,omitempty"`
	Egress            []networkingv1.NetworkPolicyEgressRule  `json:"egress,omitempty" bson:"egress,omitempty"`
	PolicyTypes       []networkingv1.PolicyType               `json:"policyTypes,omitempty" bson:"policyTypes,omitempty"`
}

func (networkPolicy *NetworkPolicyResource) GetResourceData() map[string]string {
	result := make(map[string]string)

	if len(networkPolicy.PolicyTypes) > 0 {
		result["policy-types"] = common.JoinPolicyTypes(networkPolicy.PolicyTypes, ",")
	}

	if len(networkPolicy.Selectors) > 0 {
		result["selectors"] = common.LabelsAsString(networkPolicy.Selectors)
	}

	return result
}

func (networkPolicy *NetworkPolicyResource) MatchLabels(podLabels map[string]string) bool {
	labelSelector := metav1.LabelSelector{
		MatchLabels:      networkPolicy.Selectors,
		MatchExpressions: networkPolicy.MatchExpressions,
	}

	selector, _ := metav1.LabelSelectorAsSelector(&labelSelector)
	return selector.Matches(labels.Set(podLabels))
}

func NewNetworkPolicyResource(account string, domain string, namespace string, name string) *NetworkPolicyResource {
	return &NetworkPolicyResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "NetworkPolicy",
			Name:      name,
		},
	}
}
