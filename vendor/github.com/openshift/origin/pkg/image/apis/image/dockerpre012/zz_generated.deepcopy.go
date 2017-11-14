// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package dockerpre012

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
	time "time"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_dockerpre012_Config, InType: reflect.TypeOf(&Config{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_dockerpre012_DockerConfig, InType: reflect.TypeOf(&DockerConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_dockerpre012_DockerImage, InType: reflect.TypeOf(&DockerImage{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_dockerpre012_ImagePre012, InType: reflect.TypeOf(&ImagePre012{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_dockerpre012_Mount, InType: reflect.TypeOf(&Mount{})},
	)
}

// DeepCopy_dockerpre012_Config is an autogenerated deepcopy function.
func DeepCopy_dockerpre012_Config(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Config)
		out := out.(*Config)
		*out = *in
		if in.PortSpecs != nil {
			in, out := &in.PortSpecs, &out.PortSpecs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ExposedPorts != nil {
			in, out := &in.ExposedPorts, &out.ExposedPorts
			*out = make(map[Port]struct{})
			for key := range *in {
				(*out)[key] = struct{}{}
			}
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Cmd != nil {
			in, out := &in.Cmd, &out.Cmd
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.DNS != nil {
			in, out := &in.DNS, &out.DNS
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make(map[string]struct{})
			for key := range *in {
				(*out)[key] = struct{}{}
			}
		}
		if in.Entrypoint != nil {
			in, out := &in.Entrypoint, &out.Entrypoint
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SecurityOpts != nil {
			in, out := &in.SecurityOpts, &out.SecurityOpts
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.OnBuild != nil {
			in, out := &in.OnBuild, &out.OnBuild
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Mounts != nil {
			in, out := &in.Mounts, &out.Mounts
			*out = make([]Mount, len(*in))
			copy(*out, *in)
		}
		if in.Labels != nil {
			in, out := &in.Labels, &out.Labels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_dockerpre012_DockerConfig is an autogenerated deepcopy function.
func DeepCopy_dockerpre012_DockerConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerConfig)
		out := out.(*DockerConfig)
		*out = *in
		if in.PortSpecs != nil {
			in, out := &in.PortSpecs, &out.PortSpecs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ExposedPorts != nil {
			in, out := &in.ExposedPorts, &out.ExposedPorts
			*out = make(map[string]struct{})
			for key := range *in {
				(*out)[key] = struct{}{}
			}
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Cmd != nil {
			in, out := &in.Cmd, &out.Cmd
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.DNS != nil {
			in, out := &in.DNS, &out.DNS
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make(map[string]struct{})
			for key := range *in {
				(*out)[key] = struct{}{}
			}
		}
		if in.Entrypoint != nil {
			in, out := &in.Entrypoint, &out.Entrypoint
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.SecurityOpts != nil {
			in, out := &in.SecurityOpts, &out.SecurityOpts
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.OnBuild != nil {
			in, out := &in.OnBuild, &out.OnBuild
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Labels != nil {
			in, out := &in.Labels, &out.Labels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_dockerpre012_DockerImage is an autogenerated deepcopy function.
func DeepCopy_dockerpre012_DockerImage(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerImage)
		out := out.(*DockerImage)
		*out = *in
		out.Created = in.Created.DeepCopy()
		if err := DeepCopy_dockerpre012_DockerConfig(&in.ContainerConfig, &out.ContainerConfig, c); err != nil {
			return err
		}
		if in.Config != nil {
			in, out := &in.Config, &out.Config
			*out = new(DockerConfig)
			if err := DeepCopy_dockerpre012_DockerConfig(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_dockerpre012_ImagePre012 is an autogenerated deepcopy function.
func DeepCopy_dockerpre012_ImagePre012(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImagePre012)
		out := out.(*ImagePre012)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Created); err != nil {
			return err
		} else {
			out.Created = *newVal.(*time.Time)
		}
		if err := DeepCopy_dockerpre012_Config(&in.ContainerConfig, &out.ContainerConfig, c); err != nil {
			return err
		}
		if in.Config != nil {
			in, out := &in.Config, &out.Config
			*out = new(Config)
			if err := DeepCopy_dockerpre012_Config(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_dockerpre012_Mount is an autogenerated deepcopy function.
func DeepCopy_dockerpre012_Mount(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Mount)
		out := out.(*Mount)
		*out = *in
		return nil
	}
}
