package resources

import (
	"github.com/docker/docker/oci/caps"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type OSDirectoryHostPath string

const (
	EtcDirectoryPath = "/etc"
	BinDirectoryPath = "/bin"
	VarDirectoryPath = "/var"
)

func GetOSDirectoryPaths() []string {
	return []string{
		EtcDirectoryPath,
		VarDirectoryPath,
		BinDirectoryPath,
	}
}

type PodResource struct {
	BaseResource       `bson:",inline"`
	LabelsResource     `bson:",inline"`
	OwnerReferenceKind string                `json:"ownerReferenceKind,omitempty" bson:"ownerReferenceKind,omitempty"`
	OwnerReferenceName string                `json:"ownerReferenceName,omitempty" bson:"ownerReferenceName,omitempty"`
	Containers         map[string]*Container `json:"containers" bson:"containers"`
	InitContainers     map[string]*Container `json:"initContainers" bson:"initContainers"`
	SecurityContext    *PodSecurityContext   `json:"securityContext" bson:"securityContext"`
	Volumes            map[string]*Volume    `json:"volumes" bson:"volumes"`
	PodAnnotations     map[string]string     `json:"podAnnotations" bson:"podAnnotations"`

	InstrumentedByOctarine bool    `json:"instrumentedByOctarine" bson:"instrumentedByOctarine"`
	OctarineVersion        *string `json:"octarineVersion" bson:"octarineVersion"`
	InstrumentedByIstio    bool    `json:"instrumentedByIstio" bson:"instrumentedByIstio"`
	IstioVersion           *string `json:"istioVersion" bson:"istioVersion"`
}

type ContainerVolumeMount struct {
	Name      string
	MountPath string
	ReadOnly  bool
}

type QuotasResource struct {
	Request *resource.Quantity `json:"request,omitempty" bson:"request"`
	Limit   *resource.Quantity `json:"limit,omitempty" bson:"limit"`
}

type ContainerQuotasResources struct {
	CPU    *QuotasResource `json:"cpu" bson:"cpu"`
	Memory *QuotasResource `json:"memory" bson:"memory"`
}

type Volume struct {
	Name     string
	HostPath *v1.HostPathVolumeSource
}

type Container struct {
	Name            string                           `json:"name" bson:"name"`
	Image           string                           `json:"image" bson:"image"`
	Command         []string                         `json:"command" bson:"command"`
	Env             map[string]string                `json:"env" bson:"env"`
	SecurityContext *ContainerSecurityContext        `json:"securityContext" bson:"securityContext"`
	VolumeMounts    map[string]*ContainerVolumeMount `json:"volumeMounts" bson:"volumeMounts"`
	QuotasResources *ContainerQuotasResources        `json:"quotasResources" bson:"quotasResources"`
	IsInitContainer bool                             `json:"isInitContainer" bson:"isInitContainer"`

	//Insight fields
	CapNetRawContainer       *bool `json:"capNetRawContainer,omitempty" bson:"isInitContainer"`
	CapNetAdminContainer     *bool `json:"capNetAdminContainer,omitempty" bson:"isInitContainer"`
	CapSysAdminContainer     *bool `json:"capSysAdminContainer,omitempty" bson:"isInitContainer"`
	SecurityContextContainer *bool `json:"securityContextContainer,omitempty" bson:"securityContextContainer"`
	SeccompContainer         *bool `json:"seccompContainer,omitempty" bson:"seccompContainer"`
}

func ContainsContainer(cs []*Container, c *Container) bool {
	for _, curC := range cs {
		if curC.Name == c.Name && curC.IsInitContainer == c.IsInitContainer {
			return true
		}
	}
	return false
}

func ContainersDiff(c1 []*Container, c2 []*Container) []*Container {
	res := make([]*Container, 0)
	for _, c := range c1 {
		if !ContainsContainer(c2, c) {
			res = append(res, c)
		}
	}
	return res
}

type ContainerSecurityContext struct {
	IsDefined           bool            `json:"isDefined" bson:"isDefined"`
	PrivilegeEscalation *bool           `json:"privilegeEscalation" bson:"privilegeEscalation"`
	Privileged          *bool           `json:"privileged" bson:"privileged"`
	ProcMount           *string         `json:"procMount" bson:"procMount"`
	RootFileSystem      *bool           `json:"rootFileSystem" bson:"rootFileSystem"`
	User                *int64          `json:"user" bson:"user"`
	Group               *int64          `json:"group" bson:"group"`
	RunAsRoot           *bool           `json:"runAsRoot" bson:"runAsRoot"`
	SELinuxOptions      *SELinuxOptions `json:"seLinuxOptions" bson:"seLinuxOptions"`
	Capabilities        *Capabilities   `json:"capabilities" bson:"capabilities"`
	HostNetwork         bool            `json:"hostNetwork" bson:"hostNetwork"`
	HostPID             bool            `json:"hostPID" bson:"hostPID"`
	HostIPC             bool            `json:"hostIPC" bson:"hostIPC"`
	HostPorts           []int32         `json:"hostPorts" bson:"hostPorts"`
	ContainerPorts      []int32         `json:"containerPorts" bson:"containerPorts"`
}

type SELinuxOptions struct {
	User  string `json:"user" bson:"user"`
	Role  string `json:"role" bson:"role"`
	Type  string `json:"type" bson:"type"`
	Level string `json:"level" bson:"level"`
}

type Capabilities struct {
	Add  []string `json:"add" bson:"add"`
	Drop []string `json:"drop" bson:"drop"`
}

type PodSecurityContext struct {
	FsGroup            *int64          `json:"fsGroup" bson:"fsGroup"`
	RunAsUser          *int64          `json:"runAsUser" bson:"runAsUser"`
	RunAsGroup         *int64          `json:"runAsGroup" bson:"runAsGroup"`
	SupplementalGroups []int64         `json:"supplementalGroups" bson:"supplementalGroups"`
	RunAsNonRoot       *bool           `json:"runAsNonRoot" bson:"runAsNonRoot"`
	SELinuxOptions     *SELinuxOptions `json:"seLinuxOptions" bson:"seLinuxOptions"`
	Sysctls            []NameValue     `json:"sysctls" bson:"sysctls"`
}

func (sc *PodSecurityContext) InSysCtls(k string) bool {
	if sc == nil {
		return false
	}

	for _, sysctl := range sc.Sysctls {
		if sysctl.Name == k {
			return true
		}
	}

	return false
}

type NameValue struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

func (podResource *PodResource) Clone() *PodResource {
	return &PodResource{
		BaseResource: BaseResource{
			Account:      podResource.Account,
			Domain:       podResource.Domain,
			Namespace:    podResource.Namespace,
			Kind:         podResource.Kind,
			Name:         podResource.Name,
			OctarineName: podResource.OctarineName,
		},
		LabelsResource:     podResource.LabelsResource,
		OwnerReferenceKind: podResource.OwnerReferenceKind,
		OwnerReferenceName: podResource.OwnerReferenceName,

		Containers:      podResource.Containers,
		InitContainers:  podResource.InitContainers,
		SecurityContext: podResource.SecurityContext,
		Volumes:         podResource.Volumes,
		PodAnnotations:  podResource.PodAnnotations,

		InstrumentedByOctarine: podResource.InstrumentedByOctarine,
		OctarineVersion:        podResource.OctarineVersion,
		InstrumentedByIstio:    podResource.InstrumentedByIstio,
		IstioVersion:           podResource.IstioVersion,
	}
}

func NewPodResource(account string, domain string, namespace string, name string) *PodResource {
	return &PodResource{
		BaseResource: BaseResource{
			Account:   account,
			Domain:    domain,
			Namespace: namespace,
			Kind:      "Pod",
			Name:      name,
		},
	}
}

var DefaultCapabilities = []string{
	"CAP_CHOWN",
	"CAP_DAC_OVERRIDE",
	"CAP_FSETID",
	"CAP_FOWNER",
	"CAP_MKNOD",
	"CAP_NET_RAW",
	"CAP_SETGID",
	"CAP_SETUID",
	"CAP_SETFCAP",
	"CAP_SETPCAP",
	"CAP_NET_BIND_SERVICE",
	"CAP_SYS_CHROOT",
	"CAP_KILL",
	"CAP_AUDIT_WRITE",
}

func GetDefaultCapabilities() []string {
	return DefaultCapabilities
}

func GetContainerCaps(capabilities *Capabilities, privileged bool) ([]string, error) {
	if capabilities == nil {
		return GetDefaultCapabilities(), nil
	}

	add, err := caps.NormalizeLegacyCapabilities(capabilities.Add)
	if err != nil {
		return nil, err
	}

	drop, err := caps.NormalizeLegacyCapabilities(capabilities.Drop)
	if err != nil {
		return nil, err
	}

	return caps.TweakCapabilities(GetDefaultCapabilities(), add, drop, nil, privileged)
}
