//go:build !ignore_autogenerated

/*
Copyright 2025.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KymaPolicyConfig) DeepCopyInto(out *KymaPolicyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KymaPolicyConfig.
func (in *KymaPolicyConfig) DeepCopy() *KymaPolicyConfig {
	if in == nil {
		return nil
	}
	out := new(KymaPolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KymaPolicyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KymaPolicyConfigList) DeepCopyInto(out *KymaPolicyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KymaPolicyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KymaPolicyConfigList.
func (in *KymaPolicyConfigList) DeepCopy() *KymaPolicyConfigList {
	if in == nil {
		return nil
	}
	out := new(KymaPolicyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KymaPolicyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KymaPolicyConfigSpec) DeepCopyInto(out *KymaPolicyConfigSpec) {
	*out = *in
	if in.PolicyGroups != nil {
		in, out := &in.PolicyGroups, &out.PolicyGroups
		*out = make([]KymaPolicyGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KymaPolicyConfigSpec.
func (in *KymaPolicyConfigSpec) DeepCopy() *KymaPolicyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KymaPolicyConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KymaPolicyConfigStatus) DeepCopyInto(out *KymaPolicyConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KymaPolicyConfigStatus.
func (in *KymaPolicyConfigStatus) DeepCopy() *KymaPolicyConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KymaPolicyConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KymaPolicyGroup) DeepCopyInto(out *KymaPolicyGroup) {
	*out = *in
	if in.KyvernoPolicies != nil {
		in, out := &in.KyvernoPolicies, &out.KyvernoPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KymaPolicyGroup.
func (in *KymaPolicyGroup) DeepCopy() *KymaPolicyGroup {
	if in == nil {
		return nil
	}
	out := new(KymaPolicyGroup)
	in.DeepCopyInto(out)
	return out
}
