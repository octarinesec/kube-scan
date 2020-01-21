package common

type WorkloadPod interface {
	IsPrivileged() bool
	IsCapSysAdmin() bool
	IsMountingOsDirectoryWithWritePermissions() bool
	IsMountingOsDirectoryWithReadOnlyPermissions() bool
	IsInstrumentedByOctarine() bool
	IsInstrumentedByIstio() bool
	IsNotListeningToContainerPorts() bool
	IsListeningToContainerPortsLowerThan1024() bool
	IsNotConfiguredCpuOrMemoryLimit() bool
	IsPrivilegedEscalation() bool
	IsRunningAsRoot() bool
	IsCapNetRaw() bool
	IsWritableFileSystem() bool
	IsUnmaskedProcMount() bool
	IsAllowedUnsafeSysctls() bool

	IsSecComp() bool
	IsSelinux() bool
	IsAppArmor() bool
	IsIngressPolicy() bool
	IsEgressPolicy() bool

	IsExposedByLoadBalancer() bool
	IsExposedByNodePort() bool
	IsExposedByIngress() bool
	IsHostPort() bool
	IsShareHostNetwork() bool
	IsShareHostPID() bool
	IsShareHostIPC() bool
}

type Workload interface {
	GetWorkloadPod() WorkloadPod
	GetName() string
	GetKind() string
	GetNamespace() string
	GetDomain() string
}
