//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFailureDomain) DeepCopyInto(out *AWSFailureDomain) {
	*out = *in
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(AWSResourceReference)
		(*in).DeepCopyInto(*out)
	}
	out.Placement = in.Placement
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFailureDomain.
func (in *AWSFailureDomain) DeepCopy() *AWSFailureDomain {
	if in == nil {
		return nil
	}
	out := new(AWSFailureDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFailureDomainPlacement) DeepCopyInto(out *AWSFailureDomainPlacement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFailureDomainPlacement.
func (in *AWSFailureDomainPlacement) DeepCopy() *AWSFailureDomainPlacement {
	if in == nil {
		return nil
	}
	out := new(AWSFailureDomainPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSResourceFilter) DeepCopyInto(out *AWSResourceFilter) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSResourceFilter.
func (in *AWSResourceFilter) DeepCopy() *AWSResourceFilter {
	if in == nil {
		return nil
	}
	out := new(AWSResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSResourceReference) DeepCopyInto(out *AWSResourceReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = new([]AWSResourceFilter)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AWSResourceFilter, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSResourceReference.
func (in *AWSResourceReference) DeepCopy() *AWSResourceReference {
	if in == nil {
		return nil
	}
	out := new(AWSResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibabaCloudMachineProviderConfig) DeepCopyInto(out *AlibabaCloudMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]DataDiskProperties, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]AlibabaResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Bandwidth = in.Bandwidth
	out.SystemDisk = in.SystemDisk
	in.VSwitch.DeepCopyInto(&out.VSwitch)
	in.ResourceGroup.DeepCopyInto(&out.ResourceGroup)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibabaCloudMachineProviderConfig.
func (in *AlibabaCloudMachineProviderConfig) DeepCopy() *AlibabaCloudMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AlibabaCloudMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlibabaCloudMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibabaCloudMachineProviderConfigList) DeepCopyInto(out *AlibabaCloudMachineProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlibabaCloudMachineProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibabaCloudMachineProviderConfigList.
func (in *AlibabaCloudMachineProviderConfigList) DeepCopy() *AlibabaCloudMachineProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(AlibabaCloudMachineProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlibabaCloudMachineProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibabaCloudMachineProviderStatus) DeepCopyInto(out *AlibabaCloudMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibabaCloudMachineProviderStatus.
func (in *AlibabaCloudMachineProviderStatus) DeepCopy() *AlibabaCloudMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AlibabaCloudMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlibabaCloudMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibabaResourceReference) DeepCopyInto(out *AlibabaResourceReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new([]Tag)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Tag, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibabaResourceReference.
func (in *AlibabaResourceReference) DeepCopy() *AlibabaResourceReference {
	if in == nil {
		return nil
	}
	out := new(AlibabaResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureFailureDomain) DeepCopyInto(out *AzureFailureDomain) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureFailureDomain.
func (in *AzureFailureDomain) DeepCopy() *AzureFailureDomain {
	if in == nil {
		return nil
	}
	out := new(AzureFailureDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthProperties) DeepCopyInto(out *BandwidthProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthProperties.
func (in *BandwidthProperties) DeepCopy() *BandwidthProperties {
	if in == nil {
		return nil
	}
	out := new(BandwidthProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSet) DeepCopyInto(out *ControlPlaneMachineSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSet.
func (in *ControlPlaneMachineSet) DeepCopy() *ControlPlaneMachineSet {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneMachineSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetList) DeepCopyInto(out *ControlPlaneMachineSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlPlaneMachineSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetList.
func (in *ControlPlaneMachineSetList) DeepCopy() *ControlPlaneMachineSetList {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneMachineSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetSpec) DeepCopyInto(out *ControlPlaneMachineSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.Strategy = in.Strategy
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetSpec.
func (in *ControlPlaneMachineSetSpec) DeepCopy() *ControlPlaneMachineSetSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetStatus) DeepCopyInto(out *ControlPlaneMachineSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetStatus.
func (in *ControlPlaneMachineSetStatus) DeepCopy() *ControlPlaneMachineSetStatus {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetStrategy) DeepCopyInto(out *ControlPlaneMachineSetStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetStrategy.
func (in *ControlPlaneMachineSetStrategy) DeepCopy() *ControlPlaneMachineSetStrategy {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetTemplate) DeepCopyInto(out *ControlPlaneMachineSetTemplate) {
	*out = *in
	if in.OpenShiftMachineV1Beta1Machine != nil {
		in, out := &in.OpenShiftMachineV1Beta1Machine, &out.OpenShiftMachineV1Beta1Machine
		*out = new(OpenShiftMachineV1Beta1MachineTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetTemplate.
func (in *ControlPlaneMachineSetTemplate) DeepCopy() *ControlPlaneMachineSetTemplate {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneMachineSetTemplateObjectMeta) DeepCopyInto(out *ControlPlaneMachineSetTemplateObjectMeta) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneMachineSetTemplateObjectMeta.
func (in *ControlPlaneMachineSetTemplateObjectMeta) DeepCopy() *ControlPlaneMachineSetTemplateObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneMachineSetTemplateObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDiskProperties) DeepCopyInto(out *DataDiskProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDiskProperties.
func (in *DataDiskProperties) DeepCopy() *DataDiskProperties {
	if in == nil {
		return nil
	}
	out := new(DataDiskProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureDomains) DeepCopyInto(out *FailureDomains) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new([]AWSFailureDomain)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AWSFailureDomain, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new([]AzureFailureDomain)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AzureFailureDomain, len(*in))
			copy(*out, *in)
		}
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		*out = new([]GCPFailureDomain)
		if **in != nil {
			in, out := *in, *out
			*out = make([]GCPFailureDomain, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureDomains.
func (in *FailureDomains) DeepCopy() *FailureDomains {
	if in == nil {
		return nil
	}
	out := new(FailureDomains)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPFailureDomain) DeepCopyInto(out *GCPFailureDomain) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPFailureDomain.
func (in *GCPFailureDomain) DeepCopy() *GCPFailureDomain {
	if in == nil {
		return nil
	}
	out := new(GCPFailureDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerReference) DeepCopyInto(out *LoadBalancerReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerReference.
func (in *LoadBalancerReference) DeepCopy() *LoadBalancerReference {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixCategory) DeepCopyInto(out *NutanixCategory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixCategory.
func (in *NutanixCategory) DeepCopy() *NutanixCategory {
	if in == nil {
		return nil
	}
	out := new(NutanixCategory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixMachineProviderConfig) DeepCopyInto(out *NutanixMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Cluster.DeepCopyInto(&out.Cluster)
	in.Image.DeepCopyInto(&out.Image)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]NutanixResourceIdentifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.MemorySize = in.MemorySize.DeepCopy()
	out.SystemDiskSize = in.SystemDiskSize.DeepCopy()
	in.Project.DeepCopyInto(&out.Project)
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]NutanixCategory, len(*in))
		copy(*out, *in)
	}
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixMachineProviderConfig.
func (in *NutanixMachineProviderConfig) DeepCopy() *NutanixMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(NutanixMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixMachineProviderStatus) DeepCopyInto(out *NutanixMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VmUUID != nil {
		in, out := &in.VmUUID, &out.VmUUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixMachineProviderStatus.
func (in *NutanixMachineProviderStatus) DeepCopy() *NutanixMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(NutanixMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixResourceIdentifier) DeepCopyInto(out *NutanixResourceIdentifier) {
	*out = *in
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixResourceIdentifier.
func (in *NutanixResourceIdentifier) DeepCopy() *NutanixResourceIdentifier {
	if in == nil {
		return nil
	}
	out := new(NutanixResourceIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftMachineV1Beta1MachineTemplate) DeepCopyInto(out *OpenShiftMachineV1Beta1MachineTemplate) {
	*out = *in
	in.FailureDomains.DeepCopyInto(&out.FailureDomains)
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftMachineV1Beta1MachineTemplate.
func (in *OpenShiftMachineV1Beta1MachineTemplate) DeepCopy() *OpenShiftMachineV1Beta1MachineTemplate {
	if in == nil {
		return nil
	}
	out := new(OpenShiftMachineV1Beta1MachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerVSMachineProviderConfig) DeepCopyInto(out *PowerVSMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(PowerVSSecretReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(PowerVSSecretReference)
		**out = **in
	}
	in.ServiceInstance.DeepCopyInto(&out.ServiceInstance)
	in.Image.DeepCopyInto(&out.Image)
	in.Network.DeepCopyInto(&out.Network)
	out.Processors = in.Processors
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]LoadBalancerReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerVSMachineProviderConfig.
func (in *PowerVSMachineProviderConfig) DeepCopy() *PowerVSMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(PowerVSMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PowerVSMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerVSMachineProviderStatus) DeepCopyInto(out *PowerVSMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.ServiceInstanceID != nil {
		in, out := &in.ServiceInstanceID, &out.ServiceInstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerVSMachineProviderStatus.
func (in *PowerVSMachineProviderStatus) DeepCopy() *PowerVSMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(PowerVSMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PowerVSMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerVSResource) DeepCopyInto(out *PowerVSResource) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegEx != nil {
		in, out := &in.RegEx, &out.RegEx
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerVSResource.
func (in *PowerVSResource) DeepCopy() *PowerVSResource {
	if in == nil {
		return nil
	}
	out := new(PowerVSResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerVSSecretReference) DeepCopyInto(out *PowerVSSecretReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerVSSecretReference.
func (in *PowerVSSecretReference) DeepCopy() *PowerVSSecretReference {
	if in == nil {
		return nil
	}
	out := new(PowerVSSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemDiskProperties) DeepCopyInto(out *SystemDiskProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemDiskProperties.
func (in *SystemDiskProperties) DeepCopy() *SystemDiskProperties {
	if in == nil {
		return nil
	}
	out := new(SystemDiskProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
