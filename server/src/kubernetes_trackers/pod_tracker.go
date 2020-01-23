package kubernetes_trackers

import (
	"fmt"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"kube-scan/resources"
	"strings"
)

const PodTrackerKind TrackerKind = "Pod"

type OctarinePodTracker struct {
	Account string
	Domain  string
}

func (podTracker OctarinePodTracker) GetKind() TrackerKind {
	return PodTrackerKind
}

func newPodSecurityContext(podSpec corev1.PodSpec) *resources.PodSecurityContext {
	if podSpec.SecurityContext != nil {
		sysctls := make([]resources.NameValue, 0)
		for _, s := range podSpec.SecurityContext.Sysctls {
			sysctls = append(sysctls, resources.NameValue{Name: s.Name, Value: s.Value})
		}

		var seLinuxOpts *resources.SELinuxOptions

		if podSpec.SecurityContext.SELinuxOptions != nil {
			seLinuxOpts = &resources.SELinuxOptions{
				User:  podSpec.SecurityContext.SELinuxOptions.User,
				Role:  podSpec.SecurityContext.SELinuxOptions.Role,
				Type:  podSpec.SecurityContext.SELinuxOptions.Type,
				Level: podSpec.SecurityContext.SELinuxOptions.Level,
			}
		}

		return &resources.PodSecurityContext{
			FsGroup:            podSpec.SecurityContext.FSGroup,
			RunAsUser:          podSpec.SecurityContext.RunAsUser,
			RunAsGroup:         podSpec.SecurityContext.RunAsGroup,
			SupplementalGroups: podSpec.SecurityContext.SupplementalGroups,
			RunAsNonRoot:       podSpec.SecurityContext.RunAsNonRoot,
			Sysctls:            sysctls,
			SELinuxOptions:     seLinuxOpts,
		}
	}
	return nil
}

func newContainerSecurityContext(podSpec corev1.PodSpec, container corev1.Container) *resources.ContainerSecurityContext {
	privilegeEscalation := true
	privileged := false
	procMount := string(corev1.DefaultProcMount)
	rootFileSystem := false
	var user *int64
	var group *int64
	runAsRoot := true
	capabilities := resources.Capabilities{Add: []string{}, Drop: []string{}}

	var sELinuxOptions *resources.SELinuxOptions

	if container.SecurityContext != nil && container.SecurityContext.AllowPrivilegeEscalation != nil {
		privilegeEscalation = *container.SecurityContext.AllowPrivilegeEscalation
	}
	if container.SecurityContext != nil && container.SecurityContext.Privileged != nil {
		privileged = *container.SecurityContext.Privileged
	}
	if container.SecurityContext != nil && container.SecurityContext.ProcMount != nil {
		procMount = string(*container.SecurityContext.ProcMount)
	}
	if container.SecurityContext != nil && container.SecurityContext.ReadOnlyRootFilesystem != nil {
		rootFileSystem = *container.SecurityContext.ReadOnlyRootFilesystem
	}

	if container.SecurityContext != nil && container.SecurityContext.RunAsUser != nil {
		user = container.SecurityContext.RunAsUser
	} else if podSpec.SecurityContext != nil && podSpec.SecurityContext.RunAsUser != nil {
		user = podSpec.SecurityContext.RunAsUser
	}

	if container.SecurityContext != nil && container.SecurityContext.RunAsGroup != nil {
		group = container.SecurityContext.RunAsGroup
	} else if podSpec.SecurityContext != nil && podSpec.SecurityContext.RunAsGroup != nil {
		group = podSpec.SecurityContext.RunAsGroup
	}

	if container.SecurityContext != nil && container.SecurityContext.SELinuxOptions != nil {
		sELinuxOptions = &resources.SELinuxOptions{
			User:  container.SecurityContext.SELinuxOptions.User,
			Role:  container.SecurityContext.SELinuxOptions.Role,
			Type:  container.SecurityContext.SELinuxOptions.Type,
			Level: container.SecurityContext.SELinuxOptions.Level,
		}
	} else if podSpec.SecurityContext != nil && podSpec.SecurityContext.SELinuxOptions != nil {
		sELinuxOptions = &resources.SELinuxOptions{
			User:  podSpec.SecurityContext.SELinuxOptions.User,
			Role:  podSpec.SecurityContext.SELinuxOptions.Role,
			Type:  podSpec.SecurityContext.SELinuxOptions.Type,
			Level: podSpec.SecurityContext.SELinuxOptions.Level,
		}
	}

	if container.SecurityContext != nil && container.SecurityContext.Capabilities != nil {
		for _, a := range container.SecurityContext.Capabilities.Add {
			capabilities.Add = append(capabilities.Add, string(a))
		}

		for _, d := range container.SecurityContext.Capabilities.Drop {
			capabilities.Drop = append(capabilities.Drop, string(d))
		}
	}

	runAsNonRoot := false
	if container.SecurityContext != nil && container.SecurityContext.RunAsNonRoot != nil {
		runAsNonRoot = *container.SecurityContext.RunAsNonRoot
	} else if podSpec.SecurityContext != nil && podSpec.SecurityContext.RunAsNonRoot != nil {
		runAsNonRoot = *podSpec.SecurityContext.RunAsNonRoot
	}

	if runAsNonRoot {
		runAsRoot = false
	} else if container.SecurityContext != nil && container.SecurityContext.RunAsUser != nil {
		runAsRoot = *container.SecurityContext.RunAsUser == 0
	} else if podSpec.SecurityContext != nil && podSpec.SecurityContext.RunAsUser != nil {
		runAsRoot = *podSpec.SecurityContext.RunAsUser == 0
	}

	isDefined := podSpec.SecurityContext != nil && podSpec.SecurityContext.Size() > 0
	isDefined = isDefined || (container.SecurityContext != nil && container.SecurityContext.Size() > 0)

	hostPorts := make([]int32, 0)
	containerPorts := make([]int32, 0)
	for _, port := range container.Ports {

		if port.HostPort != 0 {
			hostPorts = append(hostPorts, port.HostPort)
		}

		if port.ContainerPort != 0 {
			containerPorts = append(containerPorts, port.ContainerPort)
		}

	}

	return &resources.ContainerSecurityContext{
		IsDefined:           isDefined,
		PrivilegeEscalation: &privilegeEscalation,
		Privileged:          &privileged,
		ProcMount:           &procMount,
		RootFileSystem:      &rootFileSystem,
		User:                user,
		Group:               group,
		SELinuxOptions:      sELinuxOptions,
		RunAsRoot:           &runAsRoot,
		Capabilities:        &capabilities,
		HostNetwork:         podSpec.HostNetwork,
		HostPID:             podSpec.HostPID,
		HostIPC:             podSpec.HostIPC,
		HostPorts:           hostPorts,
		ContainerPorts:      containerPorts,
	}
}

func newContainerVolumeMounts(container corev1.Container) map[string]*resources.ContainerVolumeMount {
	mounts := make(map[string]*resources.ContainerVolumeMount)
	for _, mount := range container.VolumeMounts {
		mounts[mount.Name] = &resources.ContainerVolumeMount{
			Name:      mount.Name,
			MountPath: mount.MountPath,
			ReadOnly:  mount.ReadOnly,
		}
	}
	return mounts
}

func getQuataResorce(container corev1.Container, quataResourceName corev1.ResourceName) *resources.QuotasResource {
	quataResource := resources.QuotasResource{}
	if request, ok := container.Resources.Requests[quataResourceName]; ok {
		quataResource.Request = &request
	}
	if limit, ok := container.Resources.Limits[quataResourceName]; ok {
		quataResource.Limit = &limit
	}
	return &quataResource
}

func newContainerQuataResources(container corev1.Container) *resources.ContainerQuotasResources {
	return &resources.ContainerQuotasResources{
		CPU:    getQuataResorce(container, corev1.ResourceCPU),
		Memory: getQuataResorce(container, corev1.ResourceMemory),
	}
}

func newContainer(podSpec corev1.PodSpec, container corev1.Container, isInitContainer bool) *resources.Container {
	env := make(map[string]string)
	for _, e := range container.Env {
		env[e.Name] = e.Value
	}

	return &resources.Container{
		Name:            container.Name,
		Image:           container.Image,
		Command:         container.Command,
		Env:             env,
		SecurityContext: newContainerSecurityContext(podSpec, container),
		VolumeMounts:    newContainerVolumeMounts(container),
		QuotasResources: newContainerQuataResources(container),
		IsInitContainer: isInitContainer,
	}
}

func getContainers(podSpec corev1.PodSpec) map[string]*resources.Container {
	res := make(map[string]*resources.Container)
	for _, c := range podSpec.Containers {
		res[c.Name] = newContainer(podSpec, c, false)
	}
	return res
}

func getInitContainers(podSpec corev1.PodSpec) map[string]*resources.Container {
	res := make(map[string]*resources.Container)
	for _, c := range podSpec.InitContainers {
		res[c.Name] = newContainer(podSpec, c, true)
	}
	return res
}

func getVolumes(volumes []corev1.Volume) map[string]*resources.Volume {
	resourceVolumes := make(map[string]*resources.Volume)
	for _, volume := range volumes {

		resourceVolumes[volume.Name] = &resources.Volume{
			Name:     volume.Name,
			HostPath: volume.HostPath,
		}
	}
	return resourceVolumes
}

func getImageVersion(image string) *string {
	index := strings.LastIndex(image, ":")
	if index == -1 {
		return nil
	}

	res := image[index+1:]
	return &res
}

func (podTracker OctarinePodTracker) TrackPod(pod corev1.Pod) *resources.PodResource {
	podResource := resources.NewPodResource(podTracker.Account, podTracker.Domain, pod.Namespace, pod.Name)

	podResource.Labels = pod.Labels
	podResource.SecurityContext = newPodSecurityContext(pod.Spec)
	podResource.Containers = getContainers(pod.Spec)
	podResource.Volumes = getVolumes(pod.Spec.Volumes)
	podResource.InitContainers = getInitContainers(pod.Spec)

	for _, or := range pod.OwnerReferences {
		podResource.OwnerReferenceKind = or.Kind
		podResource.OwnerReferenceName = or.Name
	}

	artifactName := ""
	domainName := ""
	for _, container := range podResource.InitContainers {
		for envName, envValue := range container.Env {
			if envName == "OCTARINE_ID_CLIENT_ARTIFACT_ID" {
				artifactName = envValue
			} else if envName == "OCTARINE_ID_CLIENT_DOMAIN_ID" {
				domainName = envValue
			}
		}

		if strings.HasPrefix(container.Image, "octarinesec/idclient") {
			podResource.InstrumentedByOctarine = true
			podResource.OctarineVersion = getImageVersion(container.Image)
		}
	}

	for _, container := range podResource.Containers {
		if strings.HasPrefix(container.Image, "octarinesec/microservice-proxy") {
			podResource.InstrumentedByOctarine = true
			podResource.OctarineVersion = getImageVersion(container.Image)
		} else if container.Name == "istio-proxy" {
			podResource.InstrumentedByIstio = true
			podResource.IstioVersion = getImageVersion(container.Image)
		}
	}

	if artifactName != "" && domainName != "" {
		podResource.OctarineName = fmt.Sprintf("%s@%s", artifactName, domainName)
	}

	podResource.PodAnnotations = pod.Annotations

	return podResource
}

func (podTracker OctarinePodTracker) TrackResource(namespace string, name string, raw []byte) (resources.Resource, error) {
	var pod corev1.Pod

	ok, _ := unmarshelResource(raw, &pod)
	if ok {
		return podTracker.TrackPod(pod), nil
	}

	return nil, fmt.Errorf("error tracking pod")
}

func (podTracker OctarinePodTracker) TrackDelete(req *v1beta1.AdmissionRequest) (resources.Resource, error) {
	podResource := resources.NewPodResource(podTracker.Account, podTracker.Domain, req.Namespace, req.Name)

	return podResource, nil
}
