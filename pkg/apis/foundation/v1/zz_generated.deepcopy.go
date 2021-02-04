// +build !ignore_autogenerated

//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *About) DeepCopyInto(out *About) {
	*out = *in
	if in.Licenses != nil {
		in, out := &in.Licenses, &out.Licenses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new About.
func (in *About) DeepCopy() *About {
	if in == nil {
		return nil
	}
	out := new(About)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalProperties) DeepCopyInto(out *AdditionalProperties) {
	*out = *in
	if in.IsAuthorized != nil {
		in, out := &in.IsAuthorized, &out.IsAuthorized
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalProperties.
func (in *AdditionalProperties) DeepCopy() *AdditionalProperties {
	if in == nil {
		return nil
	}
	out := new(AdditionalProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectHeaderItems) DeepCopyInto(out *DetectHeaderItems) {
	*out = *in
	in.AdditionalProperties.DeepCopyInto(&out.AdditionalProperties)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectHeaderItems.
func (in *DetectHeaderItems) DeepCopy() *DetectHeaderItems {
	if in == nil {
		return nil
	}
	out := new(DetectHeaderItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Header) DeepCopyInto(out *Header) {
	*out = *in
	if in.DisabledItems != nil {
		in, out := &in.DisabledItems, &out.DisabledItems
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.DetectHeaderItems.DeepCopyInto(&out.DetectHeaderItems)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *License) DeepCopyInto(out *License) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new License.
func (in *License) DeepCopy() *License {
	if in == nil {
		return nil
	}
	out := new(License)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Login) DeepCopyInto(out *Login) {
	*out = *in
	out.LoginDialog = in.LoginDialog
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Login.
func (in *Login) DeepCopy() *Login {
	if in == nil {
		return nil
	}
	out := new(Login)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginDialog) DeepCopyInto(out *LoginDialog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginDialog.
func (in *LoginDialog) DeepCopy() *LoginDialog {
	if in == nil {
		return nil
	}
	out := new(LoginDialog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavConfiguration) DeepCopyInto(out *NavConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavConfiguration.
func (in *NavConfiguration) DeepCopy() *NavConfiguration {
	if in == nil {
		return nil
	}
	out := new(NavConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NavConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavConfigurationList) DeepCopyInto(out *NavConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NavConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavConfigurationList.
func (in *NavConfigurationList) DeepCopy() *NavConfigurationList {
	if in == nil {
		return nil
	}
	out := new(NavConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NavConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavConfigurationSpec) DeepCopyInto(out *NavConfigurationSpec) {
	*out = *in
	if in.LogoutRedirects != nil {
		in, out := &in.LogoutRedirects, &out.LogoutRedirects
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.About.DeepCopyInto(&out.About)
	in.Header.DeepCopyInto(&out.Header)
	out.Login = in.Login
	if in.NavItems != nil {
		in, out := &in.NavItems, &out.NavItems
		*out = make([]NavItems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.License = in.License
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavConfigurationSpec.
func (in *NavConfigurationSpec) DeepCopy() *NavConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(NavConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavConfigurationStatus) DeepCopyInto(out *NavConfigurationStatus) {
	*out = *in
	out.Versions = in.Versions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavConfigurationStatus.
func (in *NavConfigurationStatus) DeepCopy() *NavConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(NavConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NavItems) DeepCopyInto(out *NavItems) {
	*out = *in
	if in.IsAuthorized != nil {
		in, out := &in.IsAuthorized, &out.IsAuthorized
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NavItems.
func (in *NavItems) DeepCopy() *NavItems {
	if in == nil {
		return nil
	}
	out := new(NavItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Versions) DeepCopyInto(out *Versions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Versions.
func (in *Versions) DeepCopy() *Versions {
	if in == nil {
		return nil
	}
	out := new(Versions)
	in.DeepCopyInto(out)
	return out
}
