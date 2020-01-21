package state

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/toolkits/slice"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/common"
	"kube-scan/resources"
	"kube-scan/risk"
	"strings"
)

type Pod struct {
	*resources.PodResource `bson:",inline"`
	risk.WorkloadRisk      `bson:",inline"`
	IsDeleted              bool `json:"isDeleted" bson:"isDeleted"`

	NoNetworkPolicyIngress bool `json:"noNetworkPolicyIngress" bson:"noNetworkPolicyIngress"`
	NoNetworkPolicyEgress  bool `json:"noNetworkPolicyEgress" bson:"noNetworkPolicyEgress"`
	NoNetworkPolicy        bool `json:"noNetworkPolicy" bson:"noNetworkPolicy"`

	loadBalancerServices []string
	nodePortServices     []string
	ingressControllers   []string

	noSecurityContextContainers   []*resources.Container
	privilegedContainers          []*resources.Container
	privilegeEscalationContainers []*resources.Container
	readOnlyFileSystemContainers  []*resources.Container
	runAsRootContainers           []*resources.Container
	procMountContainers           []*resources.Container
	shareHostPidContainers        []*resources.Container
	shareHostIpcContainers        []*resources.Container
	shareHostNetworkContainers    []*resources.Container
	hostPortsContainers           []*resources.Container
	containerPortsContainers      []*resources.Container
	noSeLinuxOptsContainers       []*resources.Container
	capsNetRawContainers          []*resources.Container
	capsNetAdminContainers        []*resources.Container
	capsSysAdminContainers        []*resources.Container
	seccompContainers             []*resources.Container
	noAppArmorContainers          []*resources.Container
	wrVolumeContainers            []*resources.Container
	roVolumeContainers            []*resources.Container
	noCpuLimitContainers          []*resources.Container
	noMemoryLimitContainers       []*resources.Container

	parentName string
}

func (pod *Pod) SetParentName(name string) {
	pod.parentName = name
}

func (pod *Pod) GetWorkloadPod() common.WorkloadPod {
	return pod
}

var OctarineImages = []string{"proxy", "idclient", "iptables-redirect", "assetcopier"}

func (pod *Pod) IsPrivileged() bool {
	return len(pod.privilegedContainers) > 0
}

func (pod *Pod) IsCapSysAdmin() bool {
	return len(pod.capsSysAdminContainers) > 0
}

func (pod *Pod) findVolumeMount(containers map[string]*resources.Container, readOnly bool) []*resources.Container {
	res := make([]*resources.Container, 0)

	osDirectoryPaths := resources.GetOSDirectoryPaths()
	for _, container := range containers {
		for _, mount := range container.VolumeMounts {
			if volume, ok := pod.Volumes[mount.Name]; ok && mount.ReadOnly == readOnly && volume.HostPath != nil && slice.ContainsString(osDirectoryPaths, volume.HostPath.Path) {
				res = append(res, container)
				break
			}
		}
	}

	return res
}

func (pod *Pod) IsMountingOsDirectoryWithWritePermissions() bool {
	return len(pod.wrVolumeContainers) > 0
}

func (pod *Pod) IsMountingOsDirectoryWithReadOnlyPermissions() bool {
	return len(pod.roVolumeContainers) > 0
}

func (pod *Pod) IsInstrumentedByOctarine() bool {
	return pod.InstrumentedByOctarine
}

func (pod *Pod) IsInstrumentedByIstio() bool {
	return pod.InstrumentedByIstio
}

func (pod *Pod) IsNotListeningToContainerPorts() bool {
	return len(pod.containerPortsContainers) == 0
}

func (pod *Pod) IsListeningToContainerPortsLowerThan1024() bool {
	var limit int32 = 1024
	for _, container := range pod.containerPortsContainers {
		if container.SecurityContext == nil {
			continue
		}
		for _, port := range container.SecurityContext.ContainerPorts {
			if port < limit {
				return true
			}
		}
	}
	return false
}

func (pod *Pod) findNotConfiguredCpuContainers(containers map[string]*resources.Container) []*resources.Container {
	res := make([]*resources.Container, 0)
	if containers == nil {
		return res
	}

	for _, container := range containers {
		if container.QuotasResources == nil || container.QuotasResources.CPU == nil || container.QuotasResources.CPU.Limit == nil {
			res = append(res, container)
		}
	}

	return res
}

func (pod *Pod) findNotConfiguredMemoryContainers(containers map[string]*resources.Container) []*resources.Container {
	res := make([]*resources.Container, 0)
	if containers == nil {
		return res
	}

	for _, container := range containers {
		if container.QuotasResources == nil || container.QuotasResources.Memory == nil || container.QuotasResources.Memory.Limit == nil {
			res = append(res, container)
		}
	}

	return res
}

func (pod *Pod) IsNotConfiguredCpuOrMemoryLimit() bool {
	return len(pod.noCpuLimitContainers) > 0 || len(pod.noMemoryLimitContainers) > 0
}

func (pod *Pod) IsPrivilegedEscalation() bool {
	return len(pod.privilegeEscalationContainers) > 0
}

func (pod *Pod) IsCapNetRaw() bool {
	return len(pod.capsNetRawContainers) > 0
}

func (pod *Pod) IsWritableFileSystem() bool {
	return len(pod.readOnlyFileSystemContainers) < len(pod.Containers)+len(pod.InitContainers)
}

func (pod *Pod) IsUnmaskedProcMount() bool {
	return len(pod.procMountContainers) > 0
}

func (pod *Pod) IsAllowedUnsafeSysctls() bool {
	if pod.SecurityContext != nil {
		for _, ctl := range pod.SecurityContext.Sysctls {
			if ctl.Name == "net.ipv4.route.min_pmtu" {
				continue
			}

			if common.HasPrefix(ctl.Name, "kernel.msg", "kernel.sem", "kernel.shm", "fs.mqueue.", "net.") {
				return true
			}
		}
	}

	return false
}

func (pod *Pod) IsIngressPolicy() bool {
	return !pod.NoNetworkPolicyIngress
}

func (pod *Pod) IsEgressPolicy() bool {
	return !pod.NoNetworkPolicyEgress
}

func (pod *Pod) IsRunningAsRoot() bool {
	return len(pod.runAsRootContainers) > 0
}

func (pod *Pod) IsSecComp() bool {
	return len(pod.seccompContainers) == 0
}

func (pod *Pod) IsAppArmor() bool {
	return len(pod.noAppArmorContainers) == 0
}

func (pod *Pod) IsSelinux() bool {
	return len(pod.noSeLinuxOptsContainers) < len(pod.Containers)+len(pod.InitContainers)
}

func (pod *Pod) IsExposedByLoadBalancer() bool {
	return len(pod.loadBalancerServices) > 0
}

func (pod *Pod) IsExposedByNodePort() bool {
	return len(pod.nodePortServices) > 0
}

func (pod *Pod) IsExposedByIngress() bool {
	return len(pod.ingressControllers) > 0
}

func (pod *Pod) IsHostPort() bool {
	return len(pod.hostPortsContainers) > 0
}

func (pod *Pod) IsShareHostNetwork() bool {
	return len(pod.shareHostNetworkContainers) > 0
}

func (pod *Pod) IsShareHostPID() bool {
	return len(pod.shareHostPidContainers) > 0
}

func (pod *Pod) IsShareHostIPC() bool {
	return len(pod.shareHostIpcContainers) > 0
}

func (pod *Pod) SetNetworkPoliciesState(ns *Namespace) {
	ingressPolicies, egressPolicies := pod.getNetworkPolicies(ns)
	pod.NoNetworkPolicy = len(ingressPolicies) == 0 && len(egressPolicies) == 0
	pod.NoNetworkPolicyIngress = len(egressPolicies) == 0
	pod.NoNetworkPolicyEgress = len(egressPolicies) == 0
}

func (pod *Pod) DeleteSystemContainers() {
	for _, container := range pod.InitContainers {
		if strings.HasPrefix(container.Image, "octarinesec/idclient") || strings.HasPrefix(container.Image, "octarinesec/microservice-proxy") || strings.HasPrefix(container.Image, "octarinesec/agent_asset_copier") {
			delete(pod.InitContainers, container.Name)
		}
	}

	for _, container := range pod.Containers {
		if strings.HasPrefix(container.Image, "octarinesec/microservice-proxy") {
			delete(pod.Containers, container.Name)
		}
	}
}

func (pod *Pod) getNetworkPolicies(namespace *Namespace) (ingress []string, egress []string) {
	ingressPolicies := make([]string, 0)
	egressPolicies := make([]string, 0)
	for _, np := range namespace.NetworkPolicies {
		if np.MatchLabels(pod.Labels) {
			if len(np.Ingress) > 0 {
				ingressPolicies = append(ingressPolicies, np.GetName())
			}

			if len(np.Egress) > 0 {
				egressPolicies = append(egressPolicies, np.GetName())
			}
		}
	}

	return ingressPolicies, egressPolicies
}

func NewPod(namespace *Namespace, podResource *resources.PodResource) *Pod {
	pod := &Pod{
		PodResource: podResource,
		IsDeleted:   false,
	}
	pod.Analyze(namespace)
	return pod
}

func getImageVersion(image string) *string {
	index := strings.LastIndex(image, ":")
	if index == -1 {
		return nil
	}

	res := image[index+1:]
	return &res
}

func (pod *Pod) Analyze(namespace *Namespace) {
	noSecurityContextContainers := make([]*resources.Container, 0)
	privilegedContainers := make([]*resources.Container, 0)
	privilegeEscalationContainers := make([]*resources.Container, 0)
	readOnlyFileSystemContainers := make([]*resources.Container, 0)
	runAsRootContainers := make([]*resources.Container, 0)
	procMountContainers := make([]*resources.Container, 0)
	shareHostPidContainers := make([]*resources.Container, 0)
	shareHostIpcContainers := make([]*resources.Container, 0)
	shareHostNetworkContainers := make([]*resources.Container, 0)
	noSeLinuxOptsContainers := make([]*resources.Container, 0)
	capsNetRawContainers := make([]*resources.Container, 0)
	capsNetAdminContainers := make([]*resources.Container, 0)
	capsSysAdminContainers := make([]*resources.Container, 0)
	seccompContainers := make([]*resources.Container, 0)
	noAppArmorContainers := make([]*resources.Container, 0)
	hostPortsContainers := make([]*resources.Container, 0)
	containerPortsContainers := make([]*resources.Container, 0)
	roVolumeContainers := make([]*resources.Container, 0)
	wrVolumeContainers := make([]*resources.Container, 0)
	noCpuLimitContainers := make([]*resources.Container, 0)
	noMemoryLimitContainers := make([]*resources.Container, 0)

	pod.SetNetworkPoliciesState(namespace)

	services := make([]string, 0)
	loadBalancerServices := make([]string, 0)
	nodePortServices := make([]string, 0)
	for serviceName, service := range namespace.Services {
		if service.MatchLabels(pod.Labels) {
			services = append(services, serviceName)
			if service.IsLoadBalancer() {
				loadBalancerServices = append(loadBalancerServices, serviceName)
			} else if service.IsNodePort() {
				nodePortServices = append(nodePortServices, serviceName)
			}
		}
	}

	ingressControllers := make([]string, 0)
	if len(services) > 0 {
		for ingressName, ingress := range namespace.Ingress {
			for _, ingressService := range ingress.GetAllServices() {
				if slice.ContainsString(services, ingressService) {
					ingressControllers = append(ingressControllers, ingressName)
				}
			}
		}
	}

	for _, container := range pod.InitContainers {
		if strings.HasPrefix(container.Image, "octarinesec/idclient") {
			pod.InstrumentedByOctarine = true
			pod.OctarineVersion = getImageVersion(container.Image)
		}
	}

	for _, container := range pod.Containers {
		if strings.HasPrefix(container.Image, "octarinesec/microservice-proxy") {
			pod.InstrumentedByOctarine = true
			pod.OctarineVersion = getImageVersion(container.Image)
		} else if container.Name == "istio-proxy" {
			pod.InstrumentedByIstio = true
			pod.IstioVersion = getImageVersion(container.Image)
		}
	}

	analyzeContainers := func(containers map[string]*resources.Container) {
		for cName, c := range containers {
			if strings.HasPrefix(c.Image, "octarinesec") && slice.ContainsString(OctarineImages, cName) {
				continue
			}

			if isSeccompContainer(cName, pod) {
				seccompContainers = append(seccompContainers, c)
			}

			appArmorValue, appArmorOk := pod.PodAnnotations[fmt.Sprintf("container.apparmor.security.beta.kubernetes.io/%v", cName)]
			if (appArmorOk && appArmorValue == "unconfined") || (!appArmorOk && c.SecurityContext.Privileged != nil && *c.SecurityContext.Privileged) {
				noAppArmorContainers = append(noAppArmorContainers, c)
			}

			if !c.SecurityContext.IsDefined {
				noSecurityContextContainers = append(noSecurityContextContainers, c)
			}

			if c.SecurityContext.Privileged != nil && *c.SecurityContext.Privileged {
				privilegedContainers = append(privilegedContainers, c)
			}

			if c.SecurityContext.PrivilegeEscalation != nil && *c.SecurityContext.PrivilegeEscalation {
				privilegeEscalationContainers = append(privilegeEscalationContainers, c)
			}

			if c.SecurityContext.RootFileSystem != nil && *c.SecurityContext.RootFileSystem {
				readOnlyFileSystemContainers = append(readOnlyFileSystemContainers, c)
			}

			if c.SecurityContext.RunAsRoot != nil && *c.SecurityContext.RunAsRoot {
				runAsRootContainers = append(runAsRootContainers, c)
			}

			if c.SecurityContext.ProcMount != nil && *c.SecurityContext.ProcMount == string(corev1.UnmaskedProcMount) {
				procMountContainers = append(procMountContainers, c)
			}

			if c.SecurityContext.HostPID {
				shareHostPidContainers = append(shareHostPidContainers, c)
			}

			if c.SecurityContext.HostIPC {
				shareHostIpcContainers = append(shareHostIpcContainers, c)
			}

			if c.SecurityContext.HostNetwork {
				shareHostNetworkContainers = append(shareHostNetworkContainers, c)
			}

			if len(c.SecurityContext.HostPorts) > 0 {
				hostPortsContainers = append(hostPortsContainers, c)
			}

			if len(c.SecurityContext.ContainerPorts) > 0 {
				containerPortsContainers = append(containerPortsContainers, c)
			}

			if c.SecurityContext.SELinuxOptions == nil {
				noSeLinuxOptsContainers = append(noSeLinuxOptsContainers, c)
			}

			containerCaps, err := getCaps(c)
			if err != nil {
				glog.Errorf("error getting container '%v' capabilities, pod: '%v'. using defaults, error:", cName, pod.Name, err)
				containerCaps = resources.GetDefaultCapabilities()
			}

			if slice.ContainsString(containerCaps, "CAP_NET_RAW") {
				capsNetRawContainers = append(capsNetRawContainers, c)
			}

			if slice.ContainsString(containerCaps, "CAP_NET_ADMIN") {
				capsNetAdminContainers = append(capsNetAdminContainers, c)
			}

			if !(c.SecurityContext.Privileged != nil && *c.SecurityContext.Privileged) && slice.ContainsString(containerCaps, "CAP_SYS_ADMIN") {
				capsSysAdminContainers = append(capsSysAdminContainers, c)
			}

			roVolumeContainers = append(roVolumeContainers, pod.findVolumeMount(containers, true)...)
			wrVolumeContainers = append(wrVolumeContainers, pod.findVolumeMount(containers, false)...)

			noCpuLimitContainers = append(noCpuLimitContainers, pod.findNotConfiguredCpuContainers(containers)...)
			noMemoryLimitContainers = append(noMemoryLimitContainers, pod.findNotConfiguredMemoryContainers(containers)...)
		}
	}
	analyzeContainers(pod.Containers)
	analyzeContainers(pod.InitContainers)

	pod.loadBalancerServices = loadBalancerServices
	pod.nodePortServices = nodePortServices
	pod.ingressControllers = ingressControllers
	pod.noSecurityContextContainers = noSecurityContextContainers
	pod.privilegedContainers = privilegedContainers
	pod.privilegeEscalationContainers = privilegeEscalationContainers
	pod.readOnlyFileSystemContainers = readOnlyFileSystemContainers
	pod.runAsRootContainers = runAsRootContainers
	pod.procMountContainers = procMountContainers
	pod.shareHostPidContainers = shareHostPidContainers
	pod.shareHostIpcContainers = shareHostIpcContainers
	pod.shareHostNetworkContainers = shareHostNetworkContainers
	pod.hostPortsContainers = hostPortsContainers
	pod.containerPortsContainers = containerPortsContainers
	pod.noSeLinuxOptsContainers = noSeLinuxOptsContainers
	pod.capsNetRawContainers = capsNetRawContainers
	pod.capsNetAdminContainers = capsNetAdminContainers
	pod.capsSysAdminContainers = capsSysAdminContainers
	pod.seccompContainers = seccompContainers
	pod.noAppArmorContainers = noAppArmorContainers
	pod.roVolumeContainers = roVolumeContainers
	pod.wrVolumeContainers = wrVolumeContainers
	pod.noCpuLimitContainers = noCpuLimitContainers
	pod.noMemoryLimitContainers = noMemoryLimitContainers
}

func isSeccompContainer(cName string, pod *Pod) bool {
	seccompContainer := false
	secCompvalue, secCompOk := pod.PodAnnotations[fmt.Sprintf("container.seccomp.security.alpha.kubernetes.io/%v", cName)]
	if !secCompOk {
		secCompvalue, secCompOk = pod.PodAnnotations["seccomp.security.alpha.kubernetes.io/pod"]
	}
	if secCompOk {
		value := strings.Split(secCompvalue, ",")
		if len(value) == 0 || slice.ContainsString(value, "unconfined") {
			seccompContainer = true
		}
	} else {
		seccompContainer = true
	}

	return seccompContainer
}

func getCaps(c *resources.Container) ([]string, error) {
	privileged := c.SecurityContext.Privileged != nil && *c.SecurityContext.Privileged
	return resources.GetContainerCaps(c.SecurityContext.Capabilities, privileged)
}

func (pod *Pod) Clone() *Pod {
	return &Pod{
		PodResource:                   pod.PodResource.Clone(),
		noSecurityContextContainers:   pod.noSecurityContextContainers,
		privilegedContainers:          pod.privilegedContainers,
		privilegeEscalationContainers: pod.privilegeEscalationContainers,
		readOnlyFileSystemContainers:  pod.readOnlyFileSystemContainers,
		runAsRootContainers:           pod.runAsRootContainers,
		procMountContainers:           pod.procMountContainers,
		shareHostPidContainers:        pod.shareHostPidContainers,
		shareHostIpcContainers:        pod.shareHostIpcContainers,
		shareHostNetworkContainers:    pod.shareHostNetworkContainers,
		hostPortsContainers:           pod.hostPortsContainers,
		noSeLinuxOptsContainers:       pod.noSeLinuxOptsContainers,
		capsNetRawContainers:          pod.capsNetRawContainers,
		capsNetAdminContainers:        pod.capsNetAdminContainers,
		capsSysAdminContainers:        pod.capsSysAdminContainers,
		seccompContainers:             pod.seccompContainers,
		noAppArmorContainers:          pod.noAppArmorContainers,
		IsDeleted:                     pod.IsDeleted,
		NoNetworkPolicy:               pod.NoNetworkPolicy,
		NoNetworkPolicyIngress:        pod.NoNetworkPolicyIngress,
		NoNetworkPolicyEgress:         pod.NoNetworkPolicyEgress,
		roVolumeContainers:            pod.roVolumeContainers,
		wrVolumeContainers:            pod.wrVolumeContainers,
		noCpuLimitContainers:          pod.noCpuLimitContainers,
		noMemoryLimitContainers:       pod.noMemoryLimitContainers,
		parentName:                    pod.parentName,
	}
}
