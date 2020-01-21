package resources

import "encoding/json"

type IngressControllerResource struct {
	BaseResource `bson:",inline"`
	Backend      *IngressBackend `json:"backend" bson:"backend"`
	Rules        []*IngressRule  `json:"rules" bson:"rules"`
}

func (ingress *IngressControllerResource) GetAllServices() []string {
	result := make([]string, 0)

	for _, rule := range ingress.Rules {
		for _, path := range rule.Paths {
			result = append(result, path.Backend.ServiceName)
		}
	}

	if ingress.Backend != nil {
		result = append(result, ingress.Backend.ServiceName)
	}

	return result
}

func (ingress *IngressControllerResource) GetResourceData() map[string]string {
	result := make(map[string]string)

	if ingress.Backend != nil {
		result["default-backend"] = ingress.Backend.ServiceName
	}

	rulesJson, err := json.Marshal(ingress.Rules)
	if err == nil {
		result["rules"] = string(rulesJson)
	}

	return result
}

type IngressBackend struct {
	ServiceName string `json:"serviceName" bson:"serviceName"`
}

func NewIngressBackend(serviceName string) *IngressBackend {
	return &IngressBackend{
		ServiceName: serviceName,
	}
}

type IngressRule struct {
	Host  string         `json:"host" bson:"host"`
	Paths []*IngressPath `json:"paths" bson:"paths"`
}

func NewIngressRule(host string, paths []*IngressPath) *IngressRule {
	return &IngressRule{
		Host:  host,
		Paths: paths,
	}
}

type IngressPath struct {
	Path    string          `json:"path" bson:"path"`
	Backend *IngressBackend `json:"backend" bson:"backend"`
}

func NewIngressPath(path string, backend *IngressBackend) *IngressPath {
	return &IngressPath{
		Path:    path,
		Backend: backend,
	}
}

func NewIngressControllerResource(account string, domain string, namespace string, name string) *IngressControllerResource {
	return &IngressControllerResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "Ingress",
			Name:      name,
		},
	}
}
