// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1alpha3 "github.com/knative/pkg/apis/istio/v1alpha3"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSInfo) DeepCopyInto(out *AWSInfo) {
	*out = *in
	out.IAM = in.IAM
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSInfo.
func (in *AWSInfo) DeepCopy() *AWSInfo {
	if in == nil {
		return nil
	}
	out := new(AWSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Canary) DeepCopyInto(out *Canary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Canary.
func (in *Canary) DeepCopy() *Canary {
	if in == nil {
		return nil
	}
	out := new(Canary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAWSInfo) DeepCopyInto(out *ClusterAWSInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAWSInfo.
func (in *ClusterAWSInfo) DeepCopy() *ClusterAWSInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterAWSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConditionStatus) DeepCopyInto(out *ClusterConditionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConditionStatus.
func (in *ClusterConditionStatus) DeepCopy() *ClusterConditionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDNSGroup) DeepCopyInto(out *ClusterDNSGroup) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDNSGroup.
func (in *ClusterDNSGroup) DeepCopy() *ClusterDNSGroup {
	if in == nil {
		return nil
	}
	out := new(ClusterDNSGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngresses) DeepCopyInto(out *ClusterIngresses) {
	*out = *in
	out.Public = in.Public
	out.Private = in.Private
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngresses.
func (in *ClusterIngresses) DeepCopy() *ClusterIngresses {
	if in == nil {
		return nil
	}
	out := new(ClusterIngresses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesStatus) DeepCopyInto(out *ClusterKubernetesStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesStatus.
func (in *ClusterKubernetesStatus) DeepCopy() *ClusterKubernetesStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretSource) DeepCopyInto(out *ClusterSecretSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretSource.
func (in *ClusterSecretSource) DeepCopy() *ClusterSecretSource {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretTarget) DeepCopyInto(out *ClusterSecretTarget) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretTarget.
func (in *ClusterSecretTarget) DeepCopy() *ClusterSecretTarget {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecrets) DeepCopyInto(out *ClusterSecrets) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecrets.
func (in *ClusterSecrets) DeepCopy() *ClusterSecrets {
	if in == nil {
		return nil
	}
	out := new(ClusterSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecrets) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretsList) DeepCopyInto(out *ClusterSecretsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSecrets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretsList.
func (in *ClusterSecretsList) DeepCopy() *ClusterSecretsList {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecretsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretsSpec) DeepCopyInto(out *ClusterSecretsSpec) {
	*out = *in
	out.Source = in.Source
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretsSpec.
func (in *ClusterSecretsSpec) DeepCopy() *ClusterSecretsSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretsStatus) DeepCopyInto(out *ClusterSecretsStatus) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretsStatus.
func (in *ClusterSecretsStatus) DeepCopy() *ClusterSecretsStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ScalingFactor != nil {
		in, out := &in.ScalingFactor, &out.ScalingFactor
		*out = new(float64)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(ClusterAWSInfo)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]ClusterDNSGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Ingresses = in.Ingresses
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.Kubernetes = in.Kubernetes
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(ClusterAWSInfo)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterConditionStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalTest) DeepCopyInto(out *ExternalTest) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalTest.
func (in *ExternalTest) DeepCopy() *ExternalTest {
	if in == nil {
		return nil
	}
	out := new(ExternalTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMInfo) DeepCopyInto(out *IAMInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMInfo.
func (in *IAMInfo) DeepCopy() *IAMInfo {
	if in == nil {
		return nil
	}
	out := new(IAMInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressInfo) DeepCopyInto(out *IngressInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressInfo.
func (in *IngressInfo) DeepCopy() *IngressInfo {
	if in == nil {
		return nil
	}
	out := new(IngressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioHTTPPortConfig) DeepCopyInto(out *IstioHTTPPortConfig) {
	*out = *in
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(v1alpha3.HTTPRedirect)
		**out = **in
	}
	if in.Rewrite != nil {
		in, out := &in.Rewrite, &out.Rewrite
		*out = new(v1alpha3.HTTPRewrite)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(v1alpha3.HTTPRetry)
		**out = **in
	}
	if in.Fault != nil {
		in, out := &in.Fault, &out.Fault
		*out = new(v1alpha3.HTTPFaultInjection)
		(*in).DeepCopyInto(*out)
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(v1alpha3.Destination)
		**out = **in
	}
	if in.AppendHeaders != nil {
		in, out := &in.AppendHeaders, &out.AppendHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RemoveResponseHeaders != nil {
		in, out := &in.RemoveResponseHeaders, &out.RemoveResponseHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioHTTPPortConfig.
func (in *IstioHTTPPortConfig) DeepCopy() *IstioHTTPPortConfig {
	if in == nil {
		return nil
	}
	out := new(IstioHTTPPortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioPortConfig) DeepCopyInto(out *IstioPortConfig) {
	*out = *in
	in.HTTP.DeepCopyInto(&out.HTTP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioPortConfig.
func (in *IstioPortConfig) DeepCopy() *IstioPortConfig {
	if in == nil {
		return nil
	}
	out := new(IstioPortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mirror) DeepCopyInto(out *Mirror) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mirror.
func (in *Mirror) DeepCopy() *Mirror {
	if in == nil {
		return nil
	}
	out := new(Mirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mirror) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorList) DeepCopyInto(out *MirrorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mirror, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorList.
func (in *MirrorList) DeepCopy() *MirrorList {
	if in == nil {
		return nil
	}
	out := new(MirrorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MirrorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorSpec) DeepCopyInto(out *MirrorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorSpec.
func (in *MirrorSpec) DeepCopy() *MirrorSpec {
	if in == nil {
		return nil
	}
	out := new(MirrorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorStatus) DeepCopyInto(out *MirrorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorStatus.
func (in *MirrorStatus) DeepCopy() *MirrorStatus {
	if in == nil {
		return nil
	}
	out := new(MirrorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortInfo) DeepCopyInto(out *PortInfo) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Istio.DeepCopyInto(&out.Istio)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortInfo.
func (in *PortInfo) DeepCopy() *PortInfo {
	if in == nil {
		return nil
	}
	out := new(PortInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateInfo) DeepCopyInto(out *RateInfo) {
	*out = *in
	if in.DelaySeconds != nil {
		in, out := &in.DelaySeconds, &out.DelaySeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateInfo.
func (in *RateInfo) DeepCopy() *RateInfo {
	if in == nil {
		return nil
	}
	out := new(RateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseInfo) DeepCopyInto(out *ReleaseInfo) {
	*out = *in
	in.Rate.DeepCopyInto(&out.Rate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseInfo.
func (in *ReleaseInfo) DeepCopy() *ReleaseInfo {
	if in == nil {
		return nil
	}
	out := new(ReleaseInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManager) DeepCopyInto(out *ReleaseManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManager.
func (in *ReleaseManager) DeepCopy() *ReleaseManager {
	if in == nil {
		return nil
	}
	out := new(ReleaseManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerList) DeepCopyInto(out *ReleaseManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleaseManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerList.
func (in *ReleaseManagerList) DeepCopy() *ReleaseManagerList {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerRevisionMetricsStatus) DeepCopyInto(out *ReleaseManagerRevisionMetricsStatus) {
	*out = *in
	if in.GitReleaseSeconds != nil {
		in, out := &in.GitReleaseSeconds, &out.GitReleaseSeconds
		*out = new(float64)
		**out = **in
	}
	if in.GitDeploySeconds != nil {
		in, out := &in.GitDeploySeconds, &out.GitDeploySeconds
		*out = new(float64)
		**out = **in
	}
	if in.GitCreateSeconds != nil {
		in, out := &in.GitCreateSeconds, &out.GitCreateSeconds
		*out = new(float64)
		**out = **in
	}
	if in.RevisionDeploySeconds != nil {
		in, out := &in.RevisionDeploySeconds, &out.RevisionDeploySeconds
		*out = new(float64)
		**out = **in
	}
	if in.RevisionReleaseSeconds != nil {
		in, out := &in.RevisionReleaseSeconds, &out.RevisionReleaseSeconds
		*out = new(float64)
		**out = **in
	}
	if in.RevisionRollbackSeconds != nil {
		in, out := &in.RevisionRollbackSeconds, &out.RevisionRollbackSeconds
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerRevisionMetricsStatus.
func (in *ReleaseManagerRevisionMetricsStatus) DeepCopy() *ReleaseManagerRevisionMetricsStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerRevisionMetricsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerRevisionScaleStatus) DeepCopyInto(out *ReleaseManagerRevisionScaleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerRevisionScaleStatus.
func (in *ReleaseManagerRevisionScaleStatus) DeepCopy() *ReleaseManagerRevisionScaleStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerRevisionScaleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerRevisionStateStatus) DeepCopyInto(out *ReleaseManagerRevisionStateStatus) {
	*out = *in
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerRevisionStateStatus.
func (in *ReleaseManagerRevisionStateStatus) DeepCopy() *ReleaseManagerRevisionStateStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerRevisionStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerRevisionStatus) DeepCopyInto(out *ReleaseManagerRevisionStatus) {
	*out = *in
	in.State.DeepCopyInto(&out.State)
	if in.TriggeredAlarms != nil {
		in, out := &in.TriggeredAlarms, &out.TriggeredAlarms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.GitTimestamp != nil {
		in, out := &in.GitTimestamp, &out.GitTimestamp
		*out = (*in).DeepCopy()
	}
	if in.RevisionTimestamp != nil {
		in, out := &in.RevisionTimestamp, &out.RevisionTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CanaryStartTimestamp != nil {
		in, out := &in.CanaryStartTimestamp, &out.CanaryStartTimestamp
		*out = (*in).DeepCopy()
	}
	in.Metrics.DeepCopyInto(&out.Metrics)
	out.Scale = in.Scale
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerRevisionStatus.
func (in *ReleaseManagerRevisionStatus) DeepCopy() *ReleaseManagerRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerSpec) DeepCopyInto(out *ReleaseManagerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerSpec.
func (in *ReleaseManagerSpec) DeepCopy() *ReleaseManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerStatus) DeepCopyInto(out *ReleaseManagerStatus) {
	*out = *in
	if in.Revisions != nil {
		in, out := &in.Revisions, &out.Revisions
		*out = make([]ReleaseManagerRevisionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerStatus.
func (in *ReleaseManagerStatus) DeepCopy() *ReleaseManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Revision) DeepCopyInto(out *Revision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Revision.
func (in *Revision) DeepCopy() *Revision {
	if in == nil {
		return nil
	}
	out := new(Revision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Revision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionApp) DeepCopyInto(out *RevisionApp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionApp.
func (in *RevisionApp) DeepCopy() *RevisionApp {
	if in == nil {
		return nil
	}
	out := new(RevisionApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionList) DeepCopyInto(out *RevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Revision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionList.
func (in *RevisionList) DeepCopy() *RevisionList {
	if in == nil {
		return nil
	}
	out := new(RevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionReleaseStatus) DeepCopyInto(out *RevisionReleaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionReleaseStatus.
func (in *RevisionReleaseStatus) DeepCopy() *RevisionReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionScaleStatus) DeepCopyInto(out *RevisionScaleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionScaleStatus.
func (in *RevisionScaleStatus) DeepCopy() *RevisionScaleStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionScaleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionSpec) DeepCopyInto(out *RevisionSpec) {
	*out = *in
	out.App = in.App
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]RevisionTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(v1alpha3.TrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
	out.Sentry = in.Sentry
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionSpec.
func (in *RevisionSpec) DeepCopy() *RevisionSpec {
	if in == nil {
		return nil
	}
	out := new(RevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionStatus) DeepCopyInto(out *RevisionStatus) {
	*out = *in
	out.Sentry = in.Sentry
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]RevisionTargetStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionStatus.
func (in *RevisionStatus) DeepCopy() *RevisionStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTarget) DeepCopyInto(out *RevisionTarget) {
	*out = *in
	in.Scale.DeepCopyInto(&out.Scale)
	in.Release.DeepCopyInto(&out.Release)
	if in.ServiceMonitors != nil {
		in, out := &in.ServiceMonitors, &out.ServiceMonitors
		*out = make([]*ServiceMonitor, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceMonitor)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ServiceLevelObjectives != nil {
		in, out := &in.ServiceLevelObjectives, &out.ServiceLevelObjectives
		*out = make([]*ServiceLevelObjective, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceLevelObjective)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ConfigSelector != nil {
		in, out := &in.ConfigSelector, &out.ConfigSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.AWS = in.AWS
	if in.AlertRules != nil {
		in, out := &in.AlertRules, &out.AlertRules
		*out = make([]monitoringv1.Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	in.ExternalTest.DeepCopyInto(&out.ExternalTest)
	out.Canary = in.Canary
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTarget.
func (in *RevisionTarget) DeepCopy() *RevisionTarget {
	if in == nil {
		return nil
	}
	out := new(RevisionTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTargetStatus) DeepCopyInto(out *RevisionTargetStatus) {
	*out = *in
	out.Scale = in.Scale
	out.Release = in.Release
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTargetStatus.
func (in *RevisionTargetStatus) DeepCopy() *RevisionTargetStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLICanaryConfig) DeepCopyInto(out *SLICanaryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLICanaryConfig.
func (in *SLICanaryConfig) DeepCopy() *SLICanaryConfig {
	if in == nil {
		return nil
	}
	out := new(SLICanaryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleInfo) DeepCopyInto(out *ScaleInfo) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
	if in.TargetRequestsRate != nil {
		in, out := &in.TargetRequestsRate, &out.TargetRequestsRate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleInfo.
func (in *ScaleInfo) DeepCopy() *ScaleInfo {
	if in == nil {
		return nil
	}
	out := new(ScaleInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SentryInfo) DeepCopyInto(out *SentryInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SentryInfo.
func (in *SentryInfo) DeepCopy() *SentryInfo {
	if in == nil {
		return nil
	}
	out := new(SentryInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelIndicator) DeepCopyInto(out *ServiceLevelIndicator) {
	*out = *in
	out.Canary = in.Canary
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelIndicator.
func (in *ServiceLevelIndicator) DeepCopy() *ServiceLevelIndicator {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjective) DeepCopyInto(out *ServiceLevelObjective) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ServiceLevelIndicator = in.ServiceLevelIndicator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjective.
func (in *ServiceLevelObjective) DeepCopy() *ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitor) DeepCopyInto(out *ServiceMonitor) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitor.
func (in *ServiceMonitor) DeepCopy() *ServiceMonitor {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitor)
	in.DeepCopyInto(out)
	return out
}
