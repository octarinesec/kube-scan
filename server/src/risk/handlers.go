package risk

import (
	"fmt"
)

type Handler func(workload IWorloadRisk) bool

func GetHandler(handlerName string) (Handler, error) {
	switch handlerName {
	case "IsPrivileged":
		return isPrivileged, nil
	case "IsCapSysAdmin":
		return isCapSysAdmin, nil
	case "IsMountingOSDirectoryRW":
		return IsMountingOsDirectoryWithWritePermissions, nil
	case "IsMountingOSDirectoryRO":
		return IsMountingOsDirectoryWithReadOnlyPermissions, nil
	case "IsInstrumentedByOctarine":
		return IsInstrumentedByOctarine, nil
	case "IsInstrumentedByIstio":
		return IsInstrumentedByIstio, nil
	case "IsNotListeningToContainerPorts":
		return IsNotListeningToContainerPorts, nil
	case "IsListeningToContainerPortsLowerThan1024":
		return IsListeningToContainerPortsLowerThan1024, nil
	case "IsNotConfiguredCpuOrMemoryLimit":
		return IsNotConfiguredCpuOrMemoryLimit, nil
	case "IsRunningAsRoot":
		return isRunningAsRoot, nil
	case "IsPrivilegedEscalation":
		return isPrivilegedEscalation, nil
	case "IsCapNetRaw":
		return isCapNetRaw, nil
	case "IsWritableFileSystem":
		return isWritableFileSystem, nil
	case "IsUnmaskedProcMount":
		return isUnmaskedProcMount, nil
	case "IsAllowedUnsafeSysctls":
		return isAllowedUnsafeSysctls, nil
	case "IsSecComp":
		return isSecComp, nil
	case "IsSelinux":
		return isSelinux, nil
	case "IsAppArmor":
		return isAppArmor, nil
	case "IsIngressPolicy":
		return isIngressPolicy, nil
	case "IsEgressPolicy":
		return isEgressPolicy, nil
	case "IsExposedByLoadBalancer":
		return isExposedByLoadBalancer, nil
	case "IsExposedByNodePort":
		return isExposedByNodePort, nil
	case "IsExposedByIngress":
		return isExposedByIngress, nil
	case "IsHostPort":
		return isHostPort, nil
	case "IsShareHostNetwork":
		return isShareHostNetwork, nil
	case "IsShareHostPID":
		return isShareHostPID, nil
	case "IsShareHostIPC":
		return isShareHostIPC, nil
	default:
		return nil, fmt.Errorf("unknown handler: %s", handlerName)
	}
}

func isPrivileged(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsPrivileged()
}

func isCapSysAdmin(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsCapSysAdmin()
}

func IsMountingOsDirectoryWithWritePermissions(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsMountingOsDirectoryWithWritePermissions()
}

func IsMountingOsDirectoryWithReadOnlyPermissions(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsMountingOsDirectoryWithReadOnlyPermissions()
}

func IsInstrumentedByOctarine(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsInstrumentedByOctarine()
}

func IsInstrumentedByIstio(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsInstrumentedByIstio()
}

func IsNotListeningToContainerPorts(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsNotListeningToContainerPorts()
}

func IsListeningToContainerPortsLowerThan1024(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsListeningToContainerPortsLowerThan1024()
}

func IsNotConfiguredCpuOrMemoryLimit(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsNotConfiguredCpuOrMemoryLimit()
}

func isPrivilegedEscalation(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsPrivilegedEscalation()
}

func isRunningAsRoot(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsRunningAsRoot()
}

func isCapNetRaw(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsCapNetRaw()
}

func isWritableFileSystem(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsWritableFileSystem()
}

func isUnmaskedProcMount(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsUnmaskedProcMount()
}

func isAllowedUnsafeSysctls(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsAllowedUnsafeSysctls()
}

func isSecComp(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsSecComp()
}

func isSelinux(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsSelinux()
}

func isAppArmor(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsAppArmor()
}

func isIngressPolicy(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsIngressPolicy()
}

func isEgressPolicy(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsEgressPolicy()
}

func isExposedByLoadBalancer(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsExposedByLoadBalancer()
}

func isExposedByNodePort(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsExposedByNodePort()
}

func isExposedByIngress(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsExposedByIngress()
}

func isHostPort(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsHostPort()
}

func isShareHostNetwork(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsShareHostNetwork()
}

func isShareHostPID(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsShareHostPID()
}

func isShareHostIPC(w IWorloadRisk) bool {
	return w.GetWorkloadPod().IsShareHostIPC()
}
