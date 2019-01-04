// +build !ignore_autogenerated

/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	vmoperator "vmware.com/kubevsphere/pkg/apis/vmoperator"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*VirtualMachine)(nil), (*vmoperator.VirtualMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachine_To_vmoperator_VirtualMachine(a.(*VirtualMachine), b.(*vmoperator.VirtualMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachine)(nil), (*VirtualMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachine_To_v1beta1_VirtualMachine(a.(*vmoperator.VirtualMachine), b.(*VirtualMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineConfigStatus)(nil), (*vmoperator.VirtualMachineConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus(a.(*VirtualMachineConfigStatus), b.(*vmoperator.VirtualMachineConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineConfigStatus)(nil), (*VirtualMachineConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus(a.(*vmoperator.VirtualMachineConfigStatus), b.(*VirtualMachineConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImage)(nil), (*vmoperator.VirtualMachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImage_To_vmoperator_VirtualMachineImage(a.(*VirtualMachineImage), b.(*vmoperator.VirtualMachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImage)(nil), (*VirtualMachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImage_To_v1beta1_VirtualMachineImage(a.(*vmoperator.VirtualMachineImage), b.(*VirtualMachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImageList)(nil), (*vmoperator.VirtualMachineImageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImageList_To_vmoperator_VirtualMachineImageList(a.(*VirtualMachineImageList), b.(*vmoperator.VirtualMachineImageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImageList)(nil), (*VirtualMachineImageList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImageList_To_v1beta1_VirtualMachineImageList(a.(*vmoperator.VirtualMachineImageList), b.(*VirtualMachineImageList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImageSpec)(nil), (*vmoperator.VirtualMachineImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec(a.(*VirtualMachineImageSpec), b.(*vmoperator.VirtualMachineImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImageSpec)(nil), (*VirtualMachineImageSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec(a.(*vmoperator.VirtualMachineImageSpec), b.(*VirtualMachineImageSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImageStatus)(nil), (*vmoperator.VirtualMachineImageStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus(a.(*VirtualMachineImageStatus), b.(*vmoperator.VirtualMachineImageStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImageStatus)(nil), (*VirtualMachineImageStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus(a.(*vmoperator.VirtualMachineImageStatus), b.(*VirtualMachineImageStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImageStatusStrategy)(nil), (*vmoperator.VirtualMachineImageStatusStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImageStatusStrategy_To_vmoperator_VirtualMachineImageStatusStrategy(a.(*VirtualMachineImageStatusStrategy), b.(*vmoperator.VirtualMachineImageStatusStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImageStatusStrategy)(nil), (*VirtualMachineImageStatusStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImageStatusStrategy_To_v1beta1_VirtualMachineImageStatusStrategy(a.(*vmoperator.VirtualMachineImageStatusStrategy), b.(*VirtualMachineImageStatusStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineImageStrategy)(nil), (*vmoperator.VirtualMachineImageStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineImageStrategy_To_vmoperator_VirtualMachineImageStrategy(a.(*VirtualMachineImageStrategy), b.(*vmoperator.VirtualMachineImageStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineImageStrategy)(nil), (*VirtualMachineImageStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineImageStrategy_To_v1beta1_VirtualMachineImageStrategy(a.(*vmoperator.VirtualMachineImageStrategy), b.(*VirtualMachineImageStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineList)(nil), (*vmoperator.VirtualMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineList_To_vmoperator_VirtualMachineList(a.(*VirtualMachineList), b.(*vmoperator.VirtualMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineList)(nil), (*VirtualMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineList_To_v1beta1_VirtualMachineList(a.(*vmoperator.VirtualMachineList), b.(*VirtualMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineResourceSpec)(nil), (*vmoperator.VirtualMachineResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(a.(*VirtualMachineResourceSpec), b.(*vmoperator.VirtualMachineResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineResourceSpec)(nil), (*VirtualMachineResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(a.(*vmoperator.VirtualMachineResourceSpec), b.(*VirtualMachineResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineResourcesSpec)(nil), (*vmoperator.VirtualMachineResourcesSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec(a.(*VirtualMachineResourcesSpec), b.(*vmoperator.VirtualMachineResourcesSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineResourcesSpec)(nil), (*VirtualMachineResourcesSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec(a.(*vmoperator.VirtualMachineResourcesSpec), b.(*VirtualMachineResourcesSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineRuntimeStatus)(nil), (*vmoperator.VirtualMachineRuntimeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus(a.(*VirtualMachineRuntimeStatus), b.(*vmoperator.VirtualMachineRuntimeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineRuntimeStatus)(nil), (*VirtualMachineRuntimeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus(a.(*vmoperator.VirtualMachineRuntimeStatus), b.(*VirtualMachineRuntimeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineSpec)(nil), (*vmoperator.VirtualMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec(a.(*VirtualMachineSpec), b.(*vmoperator.VirtualMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineSpec)(nil), (*VirtualMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec(a.(*vmoperator.VirtualMachineSpec), b.(*VirtualMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineStatus)(nil), (*vmoperator.VirtualMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus(a.(*VirtualMachineStatus), b.(*vmoperator.VirtualMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineStatus)(nil), (*VirtualMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus(a.(*vmoperator.VirtualMachineStatus), b.(*VirtualMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineStatusStrategy)(nil), (*vmoperator.VirtualMachineStatusStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineStatusStrategy_To_vmoperator_VirtualMachineStatusStrategy(a.(*VirtualMachineStatusStrategy), b.(*vmoperator.VirtualMachineStatusStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineStatusStrategy)(nil), (*VirtualMachineStatusStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineStatusStrategy_To_v1beta1_VirtualMachineStatusStrategy(a.(*vmoperator.VirtualMachineStatusStrategy), b.(*VirtualMachineStatusStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VirtualMachineStrategy)(nil), (*vmoperator.VirtualMachineStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VirtualMachineStrategy_To_vmoperator_VirtualMachineStrategy(a.(*VirtualMachineStrategy), b.(*vmoperator.VirtualMachineStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*vmoperator.VirtualMachineStrategy)(nil), (*VirtualMachineStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_vmoperator_VirtualMachineStrategy_To_v1beta1_VirtualMachineStrategy(a.(*vmoperator.VirtualMachineStrategy), b.(*VirtualMachineStrategy), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_VirtualMachine_To_vmoperator_VirtualMachine(in *VirtualMachine, out *vmoperator.VirtualMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VirtualMachine_To_vmoperator_VirtualMachine is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachine_To_vmoperator_VirtualMachine(in *VirtualMachine, out *vmoperator.VirtualMachine, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachine_To_vmoperator_VirtualMachine(in, out, s)
}

func autoConvert_vmoperator_VirtualMachine_To_v1beta1_VirtualMachine(in *vmoperator.VirtualMachine, out *VirtualMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_vmoperator_VirtualMachine_To_v1beta1_VirtualMachine is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachine_To_v1beta1_VirtualMachine(in *vmoperator.VirtualMachine, out *VirtualMachine, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachine_To_v1beta1_VirtualMachine(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus(in *VirtualMachineConfigStatus, out *vmoperator.VirtualMachineConfigStatus, s conversion.Scope) error {
	out.Uuid = in.Uuid
	out.InternalId = in.InternalId
	out.CreateDate = in.CreateDate
	out.ModifiedDate = in.ModifiedDate
	return nil
}

// Convert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus(in *VirtualMachineConfigStatus, out *vmoperator.VirtualMachineConfigStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus(in *vmoperator.VirtualMachineConfigStatus, out *VirtualMachineConfigStatus, s conversion.Scope) error {
	out.Uuid = in.Uuid
	out.InternalId = in.InternalId
	out.CreateDate = in.CreateDate
	out.ModifiedDate = in.ModifiedDate
	return nil
}

// Convert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus(in *vmoperator.VirtualMachineConfigStatus, out *VirtualMachineConfigStatus, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImage_To_vmoperator_VirtualMachineImage(in *VirtualMachineImage, out *vmoperator.VirtualMachineImage, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VirtualMachineImage_To_vmoperator_VirtualMachineImage is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImage_To_vmoperator_VirtualMachineImage(in *VirtualMachineImage, out *vmoperator.VirtualMachineImage, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImage_To_vmoperator_VirtualMachineImage(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImage_To_v1beta1_VirtualMachineImage(in *vmoperator.VirtualMachineImage, out *VirtualMachineImage, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_vmoperator_VirtualMachineImage_To_v1beta1_VirtualMachineImage is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImage_To_v1beta1_VirtualMachineImage(in *vmoperator.VirtualMachineImage, out *VirtualMachineImage, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImage_To_v1beta1_VirtualMachineImage(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImageList_To_vmoperator_VirtualMachineImageList(in *VirtualMachineImageList, out *vmoperator.VirtualMachineImageList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]vmoperator.VirtualMachineImage)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_VirtualMachineImageList_To_vmoperator_VirtualMachineImageList is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImageList_To_vmoperator_VirtualMachineImageList(in *VirtualMachineImageList, out *vmoperator.VirtualMachineImageList, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImageList_To_vmoperator_VirtualMachineImageList(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImageList_To_v1beta1_VirtualMachineImageList(in *vmoperator.VirtualMachineImageList, out *VirtualMachineImageList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]VirtualMachineImage)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_vmoperator_VirtualMachineImageList_To_v1beta1_VirtualMachineImageList is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImageList_To_v1beta1_VirtualMachineImageList(in *vmoperator.VirtualMachineImageList, out *VirtualMachineImageList, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImageList_To_v1beta1_VirtualMachineImageList(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec(in *VirtualMachineImageSpec, out *vmoperator.VirtualMachineImageSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec(in *VirtualMachineImageSpec, out *vmoperator.VirtualMachineImageSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImageSpec_To_vmoperator_VirtualMachineImageSpec(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec(in *vmoperator.VirtualMachineImageSpec, out *VirtualMachineImageSpec, s conversion.Scope) error {
	return nil
}

// Convert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec(in *vmoperator.VirtualMachineImageSpec, out *VirtualMachineImageSpec, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImageSpec_To_v1beta1_VirtualMachineImageSpec(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus(in *VirtualMachineImageStatus, out *vmoperator.VirtualMachineImageStatus, s conversion.Scope) error {
	out.Uuid = in.Uuid
	out.InternalId = in.InternalId
	out.PowerState = in.PowerState
	return nil
}

// Convert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus(in *VirtualMachineImageStatus, out *vmoperator.VirtualMachineImageStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImageStatus_To_vmoperator_VirtualMachineImageStatus(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus(in *vmoperator.VirtualMachineImageStatus, out *VirtualMachineImageStatus, s conversion.Scope) error {
	out.Uuid = in.Uuid
	out.InternalId = in.InternalId
	out.PowerState = in.PowerState
	return nil
}

// Convert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus(in *vmoperator.VirtualMachineImageStatus, out *VirtualMachineImageStatus, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImageStatus_To_v1beta1_VirtualMachineImageStatus(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImageStatusStrategy_To_vmoperator_VirtualMachineImageStatusStrategy(in *VirtualMachineImageStatusStrategy, out *vmoperator.VirtualMachineImageStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1beta1_VirtualMachineImageStatusStrategy_To_vmoperator_VirtualMachineImageStatusStrategy is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImageStatusStrategy_To_vmoperator_VirtualMachineImageStatusStrategy(in *VirtualMachineImageStatusStrategy, out *vmoperator.VirtualMachineImageStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImageStatusStrategy_To_vmoperator_VirtualMachineImageStatusStrategy(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImageStatusStrategy_To_v1beta1_VirtualMachineImageStatusStrategy(in *vmoperator.VirtualMachineImageStatusStrategy, out *VirtualMachineImageStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_vmoperator_VirtualMachineImageStatusStrategy_To_v1beta1_VirtualMachineImageStatusStrategy is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImageStatusStrategy_To_v1beta1_VirtualMachineImageStatusStrategy(in *vmoperator.VirtualMachineImageStatusStrategy, out *VirtualMachineImageStatusStrategy, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImageStatusStrategy_To_v1beta1_VirtualMachineImageStatusStrategy(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineImageStrategy_To_vmoperator_VirtualMachineImageStrategy(in *VirtualMachineImageStrategy, out *vmoperator.VirtualMachineImageStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1beta1_VirtualMachineImageStrategy_To_vmoperator_VirtualMachineImageStrategy is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineImageStrategy_To_vmoperator_VirtualMachineImageStrategy(in *VirtualMachineImageStrategy, out *vmoperator.VirtualMachineImageStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineImageStrategy_To_vmoperator_VirtualMachineImageStrategy(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineImageStrategy_To_v1beta1_VirtualMachineImageStrategy(in *vmoperator.VirtualMachineImageStrategy, out *VirtualMachineImageStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_vmoperator_VirtualMachineImageStrategy_To_v1beta1_VirtualMachineImageStrategy is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineImageStrategy_To_v1beta1_VirtualMachineImageStrategy(in *vmoperator.VirtualMachineImageStrategy, out *VirtualMachineImageStrategy, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineImageStrategy_To_v1beta1_VirtualMachineImageStrategy(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineList_To_vmoperator_VirtualMachineList(in *VirtualMachineList, out *vmoperator.VirtualMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]vmoperator.VirtualMachine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_VirtualMachineList_To_vmoperator_VirtualMachineList is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineList_To_vmoperator_VirtualMachineList(in *VirtualMachineList, out *vmoperator.VirtualMachineList, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineList_To_vmoperator_VirtualMachineList(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineList_To_v1beta1_VirtualMachineList(in *vmoperator.VirtualMachineList, out *VirtualMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]VirtualMachine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_vmoperator_VirtualMachineList_To_v1beta1_VirtualMachineList is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineList_To_v1beta1_VirtualMachineList(in *vmoperator.VirtualMachineList, out *VirtualMachineList, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineList_To_v1beta1_VirtualMachineList(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(in *VirtualMachineResourceSpec, out *vmoperator.VirtualMachineResourceSpec, s conversion.Scope) error {
	out.Cpu = in.Cpu
	out.Memory = in.Memory
	return nil
}

// Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(in *VirtualMachineResourceSpec, out *vmoperator.VirtualMachineResourceSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(in *vmoperator.VirtualMachineResourceSpec, out *VirtualMachineResourceSpec, s conversion.Scope) error {
	out.Cpu = in.Cpu
	out.Memory = in.Memory
	return nil
}

// Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(in *vmoperator.VirtualMachineResourceSpec, out *VirtualMachineResourceSpec, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec(in *VirtualMachineResourcesSpec, out *vmoperator.VirtualMachineResourcesSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(&in.Capacity, &out.Capacity, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(&in.Limits, &out.Limits, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VirtualMachineResourceSpec_To_vmoperator_VirtualMachineResourceSpec(&in.Requests, &out.Requests, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec(in *VirtualMachineResourcesSpec, out *vmoperator.VirtualMachineResourcesSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec(in *vmoperator.VirtualMachineResourcesSpec, out *VirtualMachineResourcesSpec, s conversion.Scope) error {
	if err := Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(&in.Capacity, &out.Capacity, s); err != nil {
		return err
	}
	if err := Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(&in.Limits, &out.Limits, s); err != nil {
		return err
	}
	if err := Convert_vmoperator_VirtualMachineResourceSpec_To_v1beta1_VirtualMachineResourceSpec(&in.Requests, &out.Requests, s); err != nil {
		return err
	}
	return nil
}

// Convert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec(in *vmoperator.VirtualMachineResourcesSpec, out *VirtualMachineResourcesSpec, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus(in *VirtualMachineRuntimeStatus, out *vmoperator.VirtualMachineRuntimeStatus, s conversion.Scope) error {
	out.Host = in.Host
	out.PowerState = in.PowerState
	return nil
}

// Convert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus(in *VirtualMachineRuntimeStatus, out *vmoperator.VirtualMachineRuntimeStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus(in *vmoperator.VirtualMachineRuntimeStatus, out *VirtualMachineRuntimeStatus, s conversion.Scope) error {
	out.Host = in.Host
	out.PowerState = in.PowerState
	return nil
}

// Convert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus(in *vmoperator.VirtualMachineRuntimeStatus, out *VirtualMachineRuntimeStatus, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec(in *VirtualMachineSpec, out *vmoperator.VirtualMachineSpec, s conversion.Scope) error {
	out.Image = in.Image
	if err := Convert_v1beta1_VirtualMachineResourcesSpec_To_vmoperator_VirtualMachineResourcesSpec(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.PowerState = in.PowerState
	return nil
}

// Convert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec(in *VirtualMachineSpec, out *vmoperator.VirtualMachineSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineSpec_To_vmoperator_VirtualMachineSpec(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec(in *vmoperator.VirtualMachineSpec, out *VirtualMachineSpec, s conversion.Scope) error {
	out.Image = in.Image
	if err := Convert_vmoperator_VirtualMachineResourcesSpec_To_v1beta1_VirtualMachineResourcesSpec(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.PowerState = in.PowerState
	return nil
}

// Convert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec(in *vmoperator.VirtualMachineSpec, out *VirtualMachineSpec, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineSpec_To_v1beta1_VirtualMachineSpec(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus(in *VirtualMachineStatus, out *vmoperator.VirtualMachineStatus, s conversion.Scope) error {
	out.State = in.State
	if err := Convert_v1beta1_VirtualMachineConfigStatus_To_vmoperator_VirtualMachineConfigStatus(&in.ConfigStatus, &out.ConfigStatus, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VirtualMachineRuntimeStatus_To_vmoperator_VirtualMachineRuntimeStatus(&in.RuntimeStatus, &out.RuntimeStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus(in *VirtualMachineStatus, out *vmoperator.VirtualMachineStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineStatus_To_vmoperator_VirtualMachineStatus(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus(in *vmoperator.VirtualMachineStatus, out *VirtualMachineStatus, s conversion.Scope) error {
	out.State = in.State
	if err := Convert_vmoperator_VirtualMachineConfigStatus_To_v1beta1_VirtualMachineConfigStatus(&in.ConfigStatus, &out.ConfigStatus, s); err != nil {
		return err
	}
	if err := Convert_vmoperator_VirtualMachineRuntimeStatus_To_v1beta1_VirtualMachineRuntimeStatus(&in.RuntimeStatus, &out.RuntimeStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus(in *vmoperator.VirtualMachineStatus, out *VirtualMachineStatus, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineStatus_To_v1beta1_VirtualMachineStatus(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineStatusStrategy_To_vmoperator_VirtualMachineStatusStrategy(in *VirtualMachineStatusStrategy, out *vmoperator.VirtualMachineStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1beta1_VirtualMachineStatusStrategy_To_vmoperator_VirtualMachineStatusStrategy is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineStatusStrategy_To_vmoperator_VirtualMachineStatusStrategy(in *VirtualMachineStatusStrategy, out *vmoperator.VirtualMachineStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineStatusStrategy_To_vmoperator_VirtualMachineStatusStrategy(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineStatusStrategy_To_v1beta1_VirtualMachineStatusStrategy(in *vmoperator.VirtualMachineStatusStrategy, out *VirtualMachineStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_vmoperator_VirtualMachineStatusStrategy_To_v1beta1_VirtualMachineStatusStrategy is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineStatusStrategy_To_v1beta1_VirtualMachineStatusStrategy(in *vmoperator.VirtualMachineStatusStrategy, out *VirtualMachineStatusStrategy, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineStatusStrategy_To_v1beta1_VirtualMachineStatusStrategy(in, out, s)
}

func autoConvert_v1beta1_VirtualMachineStrategy_To_vmoperator_VirtualMachineStrategy(in *VirtualMachineStrategy, out *vmoperator.VirtualMachineStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1beta1_VirtualMachineStrategy_To_vmoperator_VirtualMachineStrategy is an autogenerated conversion function.
func Convert_v1beta1_VirtualMachineStrategy_To_vmoperator_VirtualMachineStrategy(in *VirtualMachineStrategy, out *vmoperator.VirtualMachineStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_VirtualMachineStrategy_To_vmoperator_VirtualMachineStrategy(in, out, s)
}

func autoConvert_vmoperator_VirtualMachineStrategy_To_v1beta1_VirtualMachineStrategy(in *vmoperator.VirtualMachineStrategy, out *VirtualMachineStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_vmoperator_VirtualMachineStrategy_To_v1beta1_VirtualMachineStrategy is an autogenerated conversion function.
func Convert_vmoperator_VirtualMachineStrategy_To_v1beta1_VirtualMachineStrategy(in *vmoperator.VirtualMachineStrategy, out *VirtualMachineStrategy, s conversion.Scope) error {
	return autoConvert_vmoperator_VirtualMachineStrategy_To_v1beta1_VirtualMachineStrategy(in, out, s)
}
