package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PortMode string

const (
	// PortPublic sets a port to be published to the Internet.
	PortPublic PortMode = "public"

	// PortPrivate sets a port to be published on a private gateway, making it available to other
	// clusters and users on our private networks (including VPN), but not on the Internet.
	PortPrivate PortMode = "private"

	// PortInternal sets the port to not be published to any gateway, making it only available within
	// the local Kubernetes cluster.
	PortLocal PortMode = "local"
)

const (
	HumaneSchedule    = "humane"
	AlwaysSchedule    = "always"
	LabelApp          = "picchu.medium.engineering/app"
	LabelTag          = "picchu.medium.engineering/tag"
	LabelCluster      = "picchu.medium.engineering/cluster"
	LabelFleet        = "picchu.medium.engineering/fleet"
	LabelRevision     = "picchu.medium.engineering/revision"
	LabelTarget       = "picchu.medium.engineering/target"
	LabelOwnerName    = "picchu.medium.engineering/ownerName"
	LabelOwnerType    = "picchu.medium.engineering/ownerType"
	LabelIstioApp     = "app"
	LabelIstioVersion = "version"
	LabelK8sName      = "app.kubernetes.io/name"
	LabelK8sVersion   = "app.kubernetes.io/version"
	LabelCommit       = "picchu.medium.engineering/commit"
	LabelRuleType     = "picchu.medium.engineering/ruleType"
	LabelFleetPrefix  = "fleet.picchu.medium.engineering/"
	// LabelTargetDeletablePrefix is used to signal that a releasemanager no longer needs the revision and it can be deleted
	LabelTargetDeletablePrefix          = "target-deletable.picchu.medium.engineering/"
	FinalizerReleaseManager             = "picchu.medium.engineering/releasemanager"
	FinalizerCluster                    = "picchu.medium.engineering/cluster"
	FinalizerClusterSecrets             = "picchu.medium.engineering/clustersecrets"
	OwnerReleaseManager                 = "releasemanager"
	AnnotationGitCommitterTimestamp     = "git-scm.com/committer-timestamp"
	AnnotationRevisionCreationTimestamp = "revisionCreationTimestamp"
	AnnotationIAMRole                   = "iam.amazonaws.com/role"
	// TODO(bob): camelCase
	AnnotationFailedAt               = "picchu.medium.engineering/failed-at-timestamp"
	AnnotationRepo                   = "picchu.medium.engineering/repo"
	AnnotationCanaryStartedTimestamp = "picchu.medium.engineering/canaryStartedTimestamp"
)

type PortInfo struct {
	Name          string          `json:"name"`
	Hosts         []string        `json:"hosts,omitempty"`
	IngressPort   int32           `json:"ingressPort,omitempty"`
	Port          int32           `json:"port,omitempty"`
	ContainerPort int32           `json:"containerPort,omitempty"`
	Protocol      corev1.Protocol `json:"protocol,omitempty"`
	Mode          PortMode        `json:"mode"`
	// Default denotes that this port will receive the default hostnames for the service. This is only useful if a
	// service exposes multiple ports over the same ingress gateway (mode: public or private). If only one port is
	// exposed, it will be the default port. If multiple ports are exposed and no default is specified and a port with
	// the `http` name is present, it will become the default port. If multiple ports are specified and the `http` port
	// is not present, validation will fail. Only Mode == (private|public) use the default flag, as they are trafficked
	// through the same external port. Internal ports can have the same hostnames since they can be selected by port
	// number.
	Default   bool     `json:"default,omitempty"`
	Ingresses []string `json:"ingresses,omitempty"`

	Istio IstioPortConfig `json:"istio,omitempty"`
}

type IstioPortConfig struct {
	HTTP IstioHTTPPortConfig `json:"http,omitempty"`
}

type Retries struct {
	Attempts      int32            `json:"attempts,omitempty"`
	PerTryTimeout *metav1.Duration `json:"perTryTimeout,omitempty"`
	RetryOn       *string          `json:"retryOn,omitempty"`
}

type IstioHTTPPortConfig struct {
	Retries *Retries         `json:"retries,omitempty"`
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

type WorkerScaleInfo struct {
	QueueURI                string   `json:"queueUri"`
	TargetMessagesPerWorker *int32   `json:"targetMessagesPerWorker"`
	SecondsToProcessOneJob  *float64 `json:"secondsToProcessOneJob,omitempty"` // optional
	MaxDisruption           *string  `json:"maxDisruption"`                    // optional
}

type ScaleInfo struct {
	Min             *int32 `json:"min,omitempty"`
	Default         int32  `json:"default,omitempty"`
	Max             int32  `json:"max,omitempty"`
	MinReadySeconds int32  `json:"minReadySeconds,omitempty"`

	// TargetCPUUtilizationPercentage scales based on CPU percentage
	TargetCPUUtilizationPercentage *int32 `json:"targetCPUUtilizationPercentage,omitempty"`

	// TargetRequestsRate scales based on the specified RequestsRateMetric
	TargetRequestsRate *string `json:"targetRequestsRate,omitempty"`
	// RequestsRateMetric refers to a Prometheus Adapter metric. See: https://github.com/DirectXMan12/k8s-prometheus-adapter
	RequestsRateMetric string `json:"requestsRateMetric,omitempty"`

	// Worker specifies parameters for Worker Pod Autoscaler. See https://github.com/practo/k8s-worker-pod-autoscaler
	Worker *WorkerScaleInfo `json:"worker,omitempty"`
}

func (s *ScaleInfo) TargetReqeustsRateQuantity() (*resource.Quantity, error) {
	if s.TargetRequestsRate == nil {
		return nil, nil
	}
	r, e := resource.ParseQuantity(*s.TargetRequestsRate)
	if e != nil {
		return nil, e
	}
	return &r, nil
}

type ReleaseInfo struct {
	Eligible bool     `json:"eligible,omitempty"`
	Max      uint32   `json:"max,omitempty"`
	Rate     RateInfo `json:"rate,omitempty"`
	Schedule string   `json:"schedule,omitempty"`
	TTL      int64    `json:"ttl,omitempty"`
}

type RateInfo struct {
	Increment    uint32 `json:"increment,omitempty"`
	DelaySeconds *int64 `json:"delay,omitempty"`
}

type SentryInfo struct {
	Release bool `json:"release,omitempty"`
}

// TODO(lyra): PodTemplate
type AWSInfo struct {
	IAM IAMInfo `json:"iam,omitempty"`
}

type IAMInfo struct {
	RoleARN string `json:"role_arn,omitempty"`
}
