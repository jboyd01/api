// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftServiceCatalogConfig) DeepCopyInto(out *OpenShiftServiceCatalogConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ServerArguments != nil {
		in, out := &in.ServerArguments, &out.ServerArguments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftServiceCatalogConfig.
func (in *OpenShiftServiceCatalogConfig) DeepCopy() *OpenShiftServiceCatalogConfig {
	if in == nil {
		return nil
	}
	out := new(OpenShiftServiceCatalogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftServiceCatalogConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
