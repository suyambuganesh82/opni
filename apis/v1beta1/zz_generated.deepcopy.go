//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	apiv1beta1 "github.com/banzaicloud/logging-operator/pkg/sdk/api/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSSpec) DeepCopyInto(out *AKSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSSpec.
func (in *AKSSpec) DeepCopy() *AKSSpec {
	if in == nil {
		return nil
	}
	out := new(AKSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthStatus) DeepCopyInto(out *AuthStatus) {
	*out = *in
	if in.AuthSecretKeyRef != nil {
		in, out := &in.AuthSecretKeyRef, &out.AuthSecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.S3AccessKey != nil {
		in, out := &in.S3AccessKey, &out.S3AccessKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.S3SecretKey != nil {
		in, out := &in.S3SecretKey, &out.S3SecretKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthStatus.
func (in *AuthStatus) DeepCopy() *AuthStatus {
	if in == nil {
		return nil
	}
	out := new(AuthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSource) DeepCopyInto(out *ContainerSource) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSource.
func (in *ContainerSource) DeepCopy() *ContainerSource {
	if in == nil {
		return nil
	}
	out := new(ContainerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrainServiceSpec) DeepCopyInto(out *DrainServiceSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrainServiceSpec.
func (in *DrainServiceSpec) DeepCopy() *DrainServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DrainServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSSpec) DeepCopyInto(out *EKSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSSpec.
func (in *EKSSpec) DeepCopy() *EKSSpec {
	if in == nil {
		return nil
	}
	out := new(EKSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticSpec) DeepCopyInto(out *ElasticSpec) {
	*out = *in
	in.Workloads.DeepCopyInto(&out.Workloads)
	if in.DefaultRepo != nil {
		in, out := &in.DefaultRepo, &out.DefaultRepo
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.KibanaImage != nil {
		in, out := &in.KibanaImage, &out.KibanaImage
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigSecret != nil {
		in, out := &in.ConfigSecret, &out.ConfigSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticSpec.
func (in *ElasticSpec) DeepCopy() *ElasticSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticWorkloadOptions) DeepCopyInto(out *ElasticWorkloadOptions) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticWorkloadOptions.
func (in *ElasticWorkloadOptions) DeepCopy() *ElasticWorkloadOptions {
	if in == nil {
		return nil
	}
	out := new(ElasticWorkloadOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticWorkloadSpec) DeepCopyInto(out *ElasticWorkloadSpec) {
	*out = *in
	in.Master.DeepCopyInto(&out.Master)
	in.Data.DeepCopyInto(&out.Data)
	in.Client.DeepCopyInto(&out.Client)
	in.Kibana.DeepCopyInto(&out.Kibana)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticWorkloadSpec.
func (in *ElasticWorkloadSpec) DeepCopy() *ElasticWorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticWorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSpec) DeepCopyInto(out *ExternalSpec) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSpec.
func (in *ExternalSpec) DeepCopy() *ExternalSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentConfigSpec) DeepCopyInto(out *FluentConfigSpec) {
	*out = *in
	if in.Fluentbit != nil {
		in, out := &in.Fluentbit, &out.Fluentbit
		*out = new(apiv1beta1.FluentbitSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(apiv1beta1.FluentdSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentConfigSpec.
func (in *FluentConfigSpec) DeepCopy() *FluentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FluentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GKESpec) DeepCopyInto(out *GKESpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GKESpec.
func (in *GKESpec) DeepCopy() *GKESpec {
	if in == nil {
		return nil
	}
	out := new(GKESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUControllerServiceSpec) DeepCopyInto(out *GPUControllerServiceSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUControllerServiceSpec.
func (in *GPUControllerServiceSpec) DeepCopy() *GPUControllerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(GPUControllerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuPolicyAdapter) DeepCopyInto(out *GpuPolicyAdapter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuPolicyAdapter.
func (in *GpuPolicyAdapter) DeepCopy() *GpuPolicyAdapter {
	if in == nil {
		return nil
	}
	out := new(GpuPolicyAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuPolicyAdapter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuPolicyAdapterList) DeepCopyInto(out *GpuPolicyAdapterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GpuPolicyAdapter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuPolicyAdapterList.
func (in *GpuPolicyAdapterList) DeepCopy() *GpuPolicyAdapterList {
	if in == nil {
		return nil
	}
	out := new(GpuPolicyAdapterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuPolicyAdapterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuPolicyAdapterSpec) DeepCopyInto(out *GpuPolicyAdapterSpec) {
	*out = *in
	out.Images = in.Images
	if in.VGPU != nil {
		in, out := &in.VGPU, &out.VGPU
		*out = new(VGPUSpec)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuPolicyAdapterSpec.
func (in *GpuPolicyAdapterSpec) DeepCopy() *GpuPolicyAdapterSpec {
	if in == nil {
		return nil
	}
	out := new(GpuPolicyAdapterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuPolicyAdapterStatus) DeepCopyInto(out *GpuPolicyAdapterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuPolicyAdapterStatus.
func (in *GpuPolicyAdapterStatus) DeepCopy() *GpuPolicyAdapterStatus {
	if in == nil {
		return nil
	}
	out := new(GpuPolicyAdapterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSource) DeepCopyInto(out *HTTPSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSource.
func (in *HTTPSource) DeepCopy() *HTTPSource {
	if in == nil {
		return nil
	}
	out := new(HTTPSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResolver) DeepCopyInto(out *ImageResolver) {
	*out = *in
	if in.DefaultRepoOverride != nil {
		in, out := &in.DefaultRepoOverride, &out.DefaultRepoOverride
		*out = new(string)
		**out = **in
	}
	if in.ImageOverride != nil {
		in, out := &in.ImageOverride, &out.ImageOverride
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResolver.
func (in *ImageResolver) DeepCopy() *ImageResolver {
	if in == nil {
		return nil
	}
	out := new(ImageResolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagesSpec) DeepCopyInto(out *ImagesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagesSpec.
func (in *ImagesSpec) DeepCopy() *ImagesSpec {
	if in == nil {
		return nil
	}
	out := new(ImagesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceServiceSpec) DeepCopyInto(out *InferenceServiceSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PretrainedModels != nil {
		in, out := &in.PretrainedModels, &out.PretrainedModels
		*out = make([]PretrainedModelReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceServiceSpec.
func (in *InferenceServiceSpec) DeepCopy() *InferenceServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InferenceServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalSpec) DeepCopyInto(out *InternalSpec) {
	*out = *in
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalSpec.
func (in *InternalSpec) DeepCopy() *InternalSpec {
	if in == nil {
		return nil
	}
	out := new(InternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K3SSpec) DeepCopyInto(out *K3SSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K3SSpec.
func (in *K3SSpec) DeepCopy() *K3SSpec {
	if in == nil {
		return nil
	}
	out := new(K3SSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAdapter) DeepCopyInto(out *LogAdapter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAdapter.
func (in *LogAdapter) DeepCopy() *LogAdapter {
	if in == nil {
		return nil
	}
	out := new(LogAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogAdapter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAdapterList) DeepCopyInto(out *LogAdapterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogAdapter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAdapterList.
func (in *LogAdapterList) DeepCopy() *LogAdapterList {
	if in == nil {
		return nil
	}
	out := new(LogAdapterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogAdapterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAdapterSpec) DeepCopyInto(out *LogAdapterSpec) {
	*out = *in
	out.OpniCluster = in.OpniCluster
	if in.AKS != nil {
		in, out := &in.AKS, &out.AKS
		*out = new(AKSSpec)
		**out = **in
	}
	if in.EKS != nil {
		in, out := &in.EKS, &out.EKS
		*out = new(EKSSpec)
		**out = **in
	}
	if in.GKE != nil {
		in, out := &in.GKE, &out.GKE
		*out = new(GKESpec)
		**out = **in
	}
	if in.K3S != nil {
		in, out := &in.K3S, &out.K3S
		*out = new(K3SSpec)
		**out = **in
	}
	if in.RKE != nil {
		in, out := &in.RKE, &out.RKE
		*out = new(RKESpec)
		**out = **in
	}
	if in.RKE2 != nil {
		in, out := &in.RKE2, &out.RKE2
		*out = new(RKE2Spec)
		**out = **in
	}
	if in.FluentConfig != nil {
		in, out := &in.FluentConfig, &out.FluentConfig
		*out = new(FluentConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RootFluentConfig != nil {
		in, out := &in.RootFluentConfig, &out.RootFluentConfig
		*out = new(FluentConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAdapterSpec.
func (in *LogAdapterSpec) DeepCopy() *LogAdapterSpec {
	if in == nil {
		return nil
	}
	out := new(LogAdapterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogAdapterStatus) DeepCopyInto(out *LogAdapterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogAdapterStatus.
func (in *LogAdapterStatus) DeepCopy() *LogAdapterStatus {
	if in == nil {
		return nil
	}
	out := new(LogAdapterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSource) DeepCopyInto(out *ModelSource) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPSource)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(ContainerSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSource.
func (in *ModelSource) DeepCopy() *ModelSource {
	if in == nil {
		return nil
	}
	out := new(ModelSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsSpec) DeepCopyInto(out *NatsSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.PasswordFrom != nil {
		in, out := &in.PasswordFrom, &out.PasswordFrom
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsSpec.
func (in *NatsSpec) DeepCopy() *NatsSpec {
	if in == nil {
		return nil
	}
	out := new(NatsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniCluster) DeepCopyInto(out *OpniCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniCluster.
func (in *OpniCluster) DeepCopy() *OpniCluster {
	if in == nil {
		return nil
	}
	out := new(OpniCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpniCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniClusterList) DeepCopyInto(out *OpniClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpniCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniClusterList.
func (in *OpniClusterList) DeepCopy() *OpniClusterList {
	if in == nil {
		return nil
	}
	out := new(OpniClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpniClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniClusterNameSpec) DeepCopyInto(out *OpniClusterNameSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniClusterNameSpec.
func (in *OpniClusterNameSpec) DeepCopy() *OpniClusterNameSpec {
	if in == nil {
		return nil
	}
	out := new(OpniClusterNameSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniClusterSpec) DeepCopyInto(out *OpniClusterSpec) {
	*out = *in
	if in.DefaultRepo != nil {
		in, out := &in.DefaultRepo, &out.DefaultRepo
		*out = new(string)
		**out = **in
	}
	in.Services.DeepCopyInto(&out.Services)
	in.Elastic.DeepCopyInto(&out.Elastic)
	in.Nats.DeepCopyInto(&out.Nats)
	in.S3.DeepCopyInto(&out.S3)
	if in.NulogHyperparameters != nil {
		in, out := &in.NulogHyperparameters, &out.NulogHyperparameters
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeployLogCollector != nil {
		in, out := &in.DeployLogCollector, &out.DeployLogCollector
		*out = new(bool)
		**out = **in
	}
	if in.GlobalNodeSelector != nil {
		in, out := &in.GlobalNodeSelector, &out.GlobalNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.GlobalTolerations != nil {
		in, out := &in.GlobalTolerations, &out.GlobalTolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniClusterSpec.
func (in *OpniClusterSpec) DeepCopy() *OpniClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OpniClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpniClusterStatus) DeepCopyInto(out *OpniClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Auth.DeepCopyInto(&out.Auth)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpniClusterStatus.
func (in *OpniClusterStatus) DeepCopy() *OpniClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OpniClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PayloadReceiverServiceSpec) DeepCopyInto(out *PayloadReceiverServiceSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PayloadReceiverServiceSpec.
func (in *PayloadReceiverServiceSpec) DeepCopy() *PayloadReceiverServiceSpec {
	if in == nil {
		return nil
	}
	out := new(PayloadReceiverServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSpec) DeepCopyInto(out *PersistenceSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	out.Request = in.Request.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSpec.
func (in *PersistenceSpec) DeepCopy() *PersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(PersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreprocessingServiceSpec) DeepCopyInto(out *PreprocessingServiceSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreprocessingServiceSpec.
func (in *PreprocessingServiceSpec) DeepCopy() *PreprocessingServiceSpec {
	if in == nil {
		return nil
	}
	out := new(PreprocessingServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PretrainedModel) DeepCopyInto(out *PretrainedModel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PretrainedModel.
func (in *PretrainedModel) DeepCopy() *PretrainedModel {
	if in == nil {
		return nil
	}
	out := new(PretrainedModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PretrainedModel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PretrainedModelList) DeepCopyInto(out *PretrainedModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PretrainedModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PretrainedModelList.
func (in *PretrainedModelList) DeepCopy() *PretrainedModelList {
	if in == nil {
		return nil
	}
	out := new(PretrainedModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PretrainedModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PretrainedModelReference) DeepCopyInto(out *PretrainedModelReference) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PretrainedModelReference.
func (in *PretrainedModelReference) DeepCopy() *PretrainedModelReference {
	if in == nil {
		return nil
	}
	out := new(PretrainedModelReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PretrainedModelSpec) DeepCopyInto(out *PretrainedModelSpec) {
	*out = *in
	in.ModelSource.DeepCopyInto(&out.ModelSource)
	if in.Hyperparameters != nil {
		in, out := &in.Hyperparameters, &out.Hyperparameters
		*out = make(map[string]intstr.IntOrString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PretrainedModelSpec.
func (in *PretrainedModelSpec) DeepCopy() *PretrainedModelSpec {
	if in == nil {
		return nil
	}
	out := new(PretrainedModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PretrainedModelStatus) DeepCopyInto(out *PretrainedModelStatus) {
	*out = *in
	out.ConfigMap = in.ConfigMap
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PretrainedModelStatus.
func (in *PretrainedModelStatus) DeepCopy() *PretrainedModelStatus {
	if in == nil {
		return nil
	}
	out := new(PretrainedModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKE2Spec) DeepCopyInto(out *RKE2Spec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKE2Spec.
func (in *RKE2Spec) DeepCopy() *RKE2Spec {
	if in == nil {
		return nil
	}
	out := new(RKE2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKESpec) DeepCopyInto(out *RKESpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKESpec.
func (in *RKESpec) DeepCopy() *RKESpec {
	if in == nil {
		return nil
	}
	out := new(RKESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	if in.Internal != nil {
		in, out := &in.Internal, &out.Internal
		*out = new(InternalSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesSpec) DeepCopyInto(out *ServicesSpec) {
	*out = *in
	in.Drain.DeepCopyInto(&out.Drain)
	in.Inference.DeepCopyInto(&out.Inference)
	in.Preprocessing.DeepCopyInto(&out.Preprocessing)
	in.PayloadReceiver.DeepCopyInto(&out.PayloadReceiver)
	in.GPUController.DeepCopyInto(&out.GPUController)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesSpec.
func (in *ServicesSpec) DeepCopy() *ServicesSpec {
	if in == nil {
		return nil
	}
	out := new(ServicesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VGPUSpec) DeepCopyInto(out *VGPUSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VGPUSpec.
func (in *VGPUSpec) DeepCopy() *VGPUSpec {
	if in == nil {
		return nil
	}
	out := new(VGPUSpec)
	in.DeepCopyInto(out)
	return out
}
