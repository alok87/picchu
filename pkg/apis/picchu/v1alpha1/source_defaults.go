package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	defaultReleaseMax              = 100
	defaultReleaseSchedule         = HumaneSchedule
	defaultReleaseRateIncrement    = 5
	defaultReleaseRateDelaySeconds = int64(10)
	defaultReleaseGcTTLSeconds     = int64(5 * 24 * 60 * 60)
	defaultScaleDefault            = int32(1)
	defaultScaleMax                = int32(1)
	defaultRequestsRateMetric      = "istio_requests_rate"

	defaultPortIngressPort   = int32(443)
	defaultPortContainerPort = int32(80)
	defaultPortPort          = int32(80)
	defaultPortProtocol      = corev1.ProtocolTCP
	defaultPortMode          = PortPrivate
)

func SetDefaults_ClusterSpec(spec *ClusterSpec) {
	if spec.ScalingFactor == nil {
		one := float64(1.0)
		spec.ScalingFactor = &one
	}
}

func SetDefaults_RevisionSpec(spec *RevisionSpec) {
	for i := range spec.Targets {
		SetReleaseDefaults(&spec.Targets[i].Release)
		SetScaleDefaults(&spec.Targets[i].Scale)
		for j := range spec.Targets[i].Ports {
			SetPortDefaults(&spec.Targets[i].Ports[j])
		}
	}
}

func SetReleaseDefaults(release *ReleaseInfo) {
	if release.Max == 0 {
		release.Max = defaultReleaseMax
	}
	if release.Schedule == "" {
		release.Schedule = defaultReleaseSchedule
	}
	if release.Rate.Increment == 0 {
		release.Rate.Increment = defaultReleaseRateIncrement
	}
	if release.Rate.DelaySeconds == nil {
		delay := defaultReleaseRateDelaySeconds
		release.Rate.DelaySeconds = &delay
	}
	if release.TTL == 0 {
		release.TTL = defaultReleaseGcTTLSeconds
	}
}

func SetPortDefaults(port *PortInfo) {
	if port.IngressPort == 0 {
		port.IngressPort = defaultPortIngressPort
	}
	if port.ContainerPort == 0 {
		port.ContainerPort = defaultPortContainerPort
	}
	if port.Port == 0 {
		port.Port = defaultPortPort
	}
	if port.Protocol == "" {
		port.Protocol = defaultPortProtocol
	}
	if port.Mode == "" {
		port.Mode = defaultPortMode
	}
}

func SetScaleDefaults(scale *ScaleInfo) {
	if scale.Default == 0 {
		scale.Default = defaultScaleDefault
	}
	if scale.Max == 0 {
		scale.Max = defaultScaleMax
	}
	if scale.RequestsRateMetric == "" {
		scale.RequestsRateMetric = defaultRequestsRateMetric
	}
}
