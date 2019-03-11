// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package aliases

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasInterfaceMap) DeepCopyInto(out *AliasInterfaceMap) {
	{
		in := &in
		*out = make(AliasInterfaceMap, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val.DeepCopyAliasInterface()
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterfaceMap.
func (in AliasInterfaceMap) DeepCopy() AliasInterfaceMap {
	if in == nil {
		return nil
	}
	out := new(AliasInterfaceMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasInterfaceSlice) DeepCopyInto(out *AliasInterfaceSlice) {
	{
		in := &in
		*out = make(AliasInterfaceSlice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyAliasInterface()
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterfaceSlice.
func (in AliasInterfaceSlice) DeepCopy() AliasInterfaceSlice {
	if in == nil {
		return nil
	}
	out := new(AliasInterfaceSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasMap) DeepCopyInto(out *AliasMap) {
	{
		in := &in
		*out = make(AliasMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasMap.
func (in AliasMap) DeepCopy() AliasMap {
	if in == nil {
		return nil
	}
	out := new(AliasMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AliasSlice) DeepCopyInto(out *AliasSlice) {
	{
		in := &in
		*out = make(AliasSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSlice.
func (in AliasSlice) DeepCopy() AliasSlice {
	if in == nil {
		return nil
	}
	out := new(AliasSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStruct) DeepCopyInto(out *AliasStruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStruct.
func (in *AliasStruct) DeepCopy() *AliasStruct {
	if in == nil {
		return nil
	}
	out := new(AliasStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Foo) DeepCopyInto(out *Foo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Foo.
func (in *Foo) DeepCopy() *Foo {
	if in == nil {
		return nil
	}
	out := new(Foo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAliasAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasAliasInterface.
func (in *Foo) DeepCopyAliasAliasInterface() AliasAliasInterface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyAliasInterface is an autogenerated deepcopy function, copying the receiver, creating a new AliasInterface.
func (in *Foo) DeepCopyAliasInterface() AliasInterface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Foo) DeepCopyInterface() Interface {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FooAlias) DeepCopyInto(out *FooAlias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooAlias.
func (in *FooAlias) DeepCopy() *FooAlias {
	if in == nil {
		return nil
	}
	out := new(FooAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FooMap) DeepCopyInto(out *FooMap) {
	{
		in := &in
		*out = make(FooMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooMap.
func (in FooMap) DeepCopy() FooMap {
	if in == nil {
		return nil
	}
	out := new(FooMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FooSlice) DeepCopyInto(out *FooSlice) {
	{
		in := &in
		*out = make(FooSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FooSlice.
func (in FooSlice) DeepCopy() FooSlice {
	if in == nil {
		return nil
	}
	out := new(FooSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Map) DeepCopyInto(out *Map) {
	{
		in := &in
		*out = make(Map, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Map.
func (in Map) DeepCopy() Map {
	if in == nil {
		return nil
	}
	out := new(Map)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Slice) DeepCopyInto(out *Slice) {
	{
		in := &in
		*out = make(Slice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slice.
func (in Slice) DeepCopy() Slice {
	if in == nil {
		return nil
	}
	out := new(Slice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct) DeepCopyInto(out *Struct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct.
func (in *Struct) DeepCopy() *Struct {
	if in == nil {
		return nil
	}
	out := new(Struct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ttest) DeepCopyInto(out *Ttest) {
	*out = *in
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = make(Slice, len(*in))
		copy(*out, *in)
	}
	if in.Pointer != nil {
		in, out := &in.Pointer, &out.Pointer
		*out = new(int)
		**out = **in
	}
	if in.PointerAlias != nil {
		in, out := &in.PointerAlias, &out.PointerAlias
		*out = new(Builtin)
		**out = **in
	}
	out.Struct = in.Struct
	if in.Map != nil {
		in, out := &in.Map, &out.Map
		*out = make(Map, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SliceSlice != nil {
		in, out := &in.SliceSlice, &out.SliceSlice
		*out = make([]Slice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(Slice, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.MapSlice != nil {
		in, out := &in.MapSlice, &out.MapSlice
		*out = make(map[string]Slice, len(*in))
		for key, val := range *in {
			var outVal []int
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Slice, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	out.FooAlias = in.FooAlias
	if in.FooSlice != nil {
		in, out := &in.FooSlice, &out.FooSlice
		*out = make(FooSlice, len(*in))
		copy(*out, *in)
	}
	if in.FooPointer != nil {
		in, out := &in.FooPointer, &out.FooPointer
		*out = new(Foo)
		**out = **in
	}
	if in.FooMap != nil {
		in, out := &in.FooMap, &out.FooMap
		*out = make(FooMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AliasSlice != nil {
		in, out := &in.AliasSlice, &out.AliasSlice
		*out = make(AliasSlice, len(*in))
		copy(*out, *in)
	}
	if in.AliasPointer != nil {
		in, out := &in.AliasPointer, &out.AliasPointer
		*out = new(int)
		**out = **in
	}
	out.AliasStruct = in.AliasStruct
	if in.AliasMap != nil {
		in, out := &in.AliasMap, &out.AliasMap
		*out = make(AliasMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AliasInterface != nil {
		out.AliasInterface = in.AliasInterface.DeepCopyAliasInterface()
	}
	if in.AliasAliasInterface != nil {
		out.AliasAliasInterface = in.AliasAliasInterface.DeepCopyAliasAliasInterface()
	}
	if in.AliasInterfaceMap != nil {
		in, out := &in.AliasInterfaceMap, &out.AliasInterfaceMap
		*out = make(AliasInterfaceMap, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = val.DeepCopyAliasInterface()
			}
		}
	}
	if in.AliasInterfaceSlice != nil {
		in, out := &in.AliasInterfaceSlice, &out.AliasInterfaceSlice
		*out = make(AliasInterfaceSlice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyAliasInterface()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ttest.
func (in *Ttest) DeepCopy() *Ttest {
	if in == nil {
		return nil
	}
	out := new(Ttest)
	in.DeepCopyInto(out)
	return out
}
