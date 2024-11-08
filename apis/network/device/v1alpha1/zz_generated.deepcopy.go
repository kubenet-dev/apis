//go:build !ignore_autogenerated

/*
Copyright 2024 Nokia.

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
	corev1alpha1 "github.com/kubenet-dev/apis/apis/network/core/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BFDSpec) DeepCopyInto(out *BFDSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.BFD != nil {
		in, out := &in.BFD, &out.BFD
		*out = new(corev1alpha1.BFDLinkParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BFDSpec.
func (in *BFDSpec) DeepCopy() *BFDSpec {
	if in == nil {
		return nil
	}
	out := new(BFDSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BFDStatus) DeepCopyInto(out *BFDStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BFDStatus.
func (in *BFDStatus) DeepCopy() *BFDStatus {
	if in == nil {
		return nil
	}
	out := new(BFDStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPAddressFamily) DeepCopyInto(out *BGPAddressFamily) {
	*out = *in
	if in.RFC5549 != nil {
		in, out := &in.RFC5549, &out.RFC5549
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPAddressFamily.
func (in *BGPAddressFamily) DeepCopy() *BGPAddressFamily {
	if in == nil {
		return nil
	}
	out := new(BGPAddressFamily)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPDynamicNeighborInterface) DeepCopyInto(out *BGPDynamicNeighborInterface) {
	*out = *in
	in.PartitionEndpointID.DeepCopyInto(&out.PartitionEndpointID)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPDynamicNeighborInterface.
func (in *BGPDynamicNeighborInterface) DeepCopy() *BGPDynamicNeighborInterface {
	if in == nil {
		return nil
	}
	out := new(BGPDynamicNeighborInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPDynamicNeighborSpec) DeepCopyInto(out *BGPDynamicNeighborSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]BGPDynamicNeighborInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPDynamicNeighborSpec.
func (in *BGPDynamicNeighborSpec) DeepCopy() *BGPDynamicNeighborSpec {
	if in == nil {
		return nil
	}
	out := new(BGPDynamicNeighborSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPDynamicNeighborStatus) DeepCopyInto(out *BGPDynamicNeighborStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPDynamicNeighborStatus.
func (in *BGPDynamicNeighborStatus) DeepCopy() *BGPDynamicNeighborStatus {
	if in == nil {
		return nil
	}
	out := new(BGPDynamicNeighborStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPNeighborSpec) DeepCopyInto(out *BGPNeighborSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.BFD != nil {
		in, out := &in.BFD, &out.BFD
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPNeighborSpec.
func (in *BGPNeighborSpec) DeepCopy() *BGPNeighborSpec {
	if in == nil {
		return nil
	}
	out := new(BGPNeighborSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPNeighborStatus) DeepCopyInto(out *BGPNeighborStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPNeighborStatus.
func (in *BGPNeighborStatus) DeepCopy() *BGPNeighborStatus {
	if in == nil {
		return nil
	}
	out := new(BGPNeighborStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerGroup) DeepCopyInto(out *BGPPeerGroup) {
	*out = *in
	if in.AddressFamilies != nil {
		in, out := &in.AddressFamilies, &out.AddressFamilies
		*out = make([]*BGPAddressFamily, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BGPAddressFamily)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.RouteReflector != nil {
		in, out := &in.RouteReflector, &out.RouteReflector
		*out = new(BGPRouteReflector)
		**out = **in
	}
	if in.BFD != nil {
		in, out := &in.BFD, &out.BFD
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerGroup.
func (in *BGPPeerGroup) DeepCopy() *BGPPeerGroup {
	if in == nil {
		return nil
	}
	out := new(BGPPeerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPRouteReflector) DeepCopyInto(out *BGPRouteReflector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPRouteReflector.
func (in *BGPRouteReflector) DeepCopy() *BGPRouteReflector {
	if in == nil {
		return nil
	}
	out := new(BGPRouteReflector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPSpec) DeepCopyInto(out *BGPSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.AddressFamilies != nil {
		in, out := &in.AddressFamilies, &out.AddressFamilies
		*out = make([]*BGPAddressFamily, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BGPAddressFamily)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.PeerGroups != nil {
		in, out := &in.PeerGroups, &out.PeerGroups
		*out = make([]BGPPeerGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPSpec.
func (in *BGPSpec) DeepCopy() *BGPSpec {
	if in == nil {
		return nil
	}
	out := new(BGPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPStatus) DeepCopyInto(out *BGPStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPStatus.
func (in *BGPStatus) DeepCopy() *BGPStatus {
	if in == nil {
		return nil
	}
	out := new(BGPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interface) DeepCopyInto(out *Interface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Interface) DeepCopy() *Interface {
	if in == nil {
		return nil
	}
	out := new(Interface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Interface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceList) DeepCopyInto(out *InterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Interface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceList.
func (in *InterfaceList) DeepCopy() *InterfaceList {
	if in == nil {
		return nil
	}
	out := new(InterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSpec) DeepCopyInto(out *InterfaceSpec) {
	*out = *in
	in.PartitionProviderEndpointID.DeepCopyInto(&out.PartitionProviderEndpointID)
	if in.Ethernet != nil {
		in, out := &in.Ethernet, &out.Ethernet
		*out = new(InterfaceSpecEthernet)
		**out = **in
	}
	if in.VLANTagging != nil {
		in, out := &in.VLANTagging, &out.VLANTagging
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSpec.
func (in *InterfaceSpec) DeepCopy() *InterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(InterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSpecEthernet) DeepCopyInto(out *InterfaceSpecEthernet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSpecEthernet.
func (in *InterfaceSpecEthernet) DeepCopy() *InterfaceSpecEthernet {
	if in == nil {
		return nil
	}
	out := new(InterfaceSpecEthernet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceStatus) DeepCopyInto(out *InterfaceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceStatus.
func (in *InterfaceStatus) DeepCopy() *InterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(InterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceInterface) DeepCopyInto(out *NetworkInstanceInterface) {
	*out = *in
	in.PartitionEndpointID.DeepCopyInto(&out.PartitionEndpointID)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceInterface.
func (in *NetworkInstanceInterface) DeepCopy() *NetworkInstanceInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceSpec) DeepCopyInto(out *NetworkInstanceSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]*NetworkInstanceInterface, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NetworkInstanceInterface)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VXLANInterface != nil {
		in, out := &in.VXLANInterface, &out.VXLANInterface
		*out = new(NetworkInstanceVXLANInterface)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceSpec.
func (in *NetworkInstanceSpec) DeepCopy() *NetworkInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceStatus) DeepCopyInto(out *NetworkInstanceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceStatus.
func (in *NetworkInstanceStatus) DeepCopy() *NetworkInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceVXLANInterface) DeepCopyInto(out *NetworkInstanceVXLANInterface) {
	*out = *in
	in.PartitionEndpointID.DeepCopyInto(&out.PartitionEndpointID)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceVXLANInterface.
func (in *NetworkInstanceVXLANInterface) DeepCopy() *NetworkInstanceVXLANInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceVXLANInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplatePort) DeepCopyInto(out *NodeTemplatePort) {
	*out = *in
	out.Ids = in.Ids
	out.Adaptor = in.Adaptor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplatePort.
func (in *NodeTemplatePort) DeepCopy() *NodeTemplatePort {
	if in == nil {
		return nil
	}
	out := new(NodeTemplatePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplatePortAdaptor) DeepCopyInto(out *NodeTemplatePortAdaptor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplatePortAdaptor.
func (in *NodeTemplatePortAdaptor) DeepCopy() *NodeTemplatePortAdaptor {
	if in == nil {
		return nil
	}
	out := new(NodeTemplatePortAdaptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplatePortIds) DeepCopyInto(out *NodeTemplatePortIds) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplatePortIds.
func (in *NodeTemplatePortIds) DeepCopy() *NodeTemplatePortIds {
	if in == nil {
		return nil
	}
	out := new(NodeTemplatePortIds)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplateSpec) DeepCopyInto(out *NodeTemplateSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*NodeTemplatePort, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeTemplatePort)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplateSpec.
func (in *NodeTemplateSpec) DeepCopy() *NodeTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplateStatus) DeepCopyInto(out *NodeTemplateStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplateStatus.
func (in *NodeTemplateStatus) DeepCopy() *NodeTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(NodeTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixSetPrefix) DeepCopyInto(out *PrefixSetPrefix) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixSetPrefix.
func (in *PrefixSetPrefix) DeepCopy() *PrefixSetPrefix {
	if in == nil {
		return nil
	}
	out := new(PrefixSetPrefix)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixSetSpec) DeepCopyInto(out *PrefixSetSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]*PrefixSetPrefix, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PrefixSetPrefix)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixSetSpec.
func (in *PrefixSetSpec) DeepCopy() *PrefixSetSpec {
	if in == nil {
		return nil
	}
	out := new(PrefixSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixSetStatus) DeepCopyInto(out *PrefixSetStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixSetStatus.
func (in *PrefixSetStatus) DeepCopy() *PrefixSetStatus {
	if in == nil {
		return nil
	}
	out := new(PrefixSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicySpec) DeepCopyInto(out *RoutingPolicySpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
	in.DefaultAction.DeepCopyInto(&out.DefaultAction)
	if in.Statements != nil {
		in, out := &in.Statements, &out.Statements
		*out = make([]*RoutingPolicyStatement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RoutingPolicyStatement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicySpec.
func (in *RoutingPolicySpec) DeepCopy() *RoutingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatement) DeepCopyInto(out *RoutingPolicyStatement) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	in.Match.DeepCopyInto(&out.Match)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatement.
func (in *RoutingPolicyStatement) DeepCopy() *RoutingPolicyStatement {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementAction) DeepCopyInto(out *RoutingPolicyStatementAction) {
	*out = *in
	if in.BGP != nil {
		in, out := &in.BGP, &out.BGP
		*out = new(RoutingPolicyStatementActionBGP)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementAction.
func (in *RoutingPolicyStatementAction) DeepCopy() *RoutingPolicyStatementAction {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementActionBGP) DeepCopyInto(out *RoutingPolicyStatementActionBGP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementActionBGP.
func (in *RoutingPolicyStatementActionBGP) DeepCopy() *RoutingPolicyStatementActionBGP {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementActionBGP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementMatch) DeepCopyInto(out *RoutingPolicyStatementMatch) {
	*out = *in
	if in.BGP != nil {
		in, out := &in.BGP, &out.BGP
		*out = new(RoutingPolicyStatementMatchBGP)
		(*in).DeepCopyInto(*out)
	}
	if in.OSPF != nil {
		in, out := &in.OSPF, &out.OSPF
		*out = new(RoutingPolicyStatementMatchOSPF)
		**out = **in
	}
	if in.ISIS != nil {
		in, out := &in.ISIS, &out.ISIS
		*out = new(RoutingPolicyStatementMatchISIS)
		**out = **in
	}
	if in.PrefixSetRef != nil {
		in, out := &in.PrefixSetRef, &out.PrefixSetRef
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.TagSetRef != nil {
		in, out := &in.TagSetRef, &out.TagSetRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementMatch.
func (in *RoutingPolicyStatementMatch) DeepCopy() *RoutingPolicyStatementMatch {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementMatchBGP) DeepCopyInto(out *RoutingPolicyStatementMatchBGP) {
	*out = *in
	if in.ASPathLength != nil {
		in, out := &in.ASPathLength, &out.ASPathLength
		*out = new(RoutingPolicyStatementMatchBGPASPathLength)
		(*in).DeepCopyInto(*out)
	}
	if in.ASPathSetRef != nil {
		in, out := &in.ASPathSetRef, &out.ASPathSetRef
		*out = new(string)
		**out = **in
	}
	if in.CommunitySetRef != nil {
		in, out := &in.CommunitySetRef, &out.CommunitySetRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementMatchBGP.
func (in *RoutingPolicyStatementMatchBGP) DeepCopy() *RoutingPolicyStatementMatchBGP {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementMatchBGP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementMatchBGPASPathLength) DeepCopyInto(out *RoutingPolicyStatementMatchBGPASPathLength) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Unique != nil {
		in, out := &in.Unique, &out.Unique
		*out = new(bool)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementMatchBGPASPathLength.
func (in *RoutingPolicyStatementMatchBGPASPathLength) DeepCopy() *RoutingPolicyStatementMatchBGPASPathLength {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementMatchBGPASPathLength)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementMatchISIS) DeepCopyInto(out *RoutingPolicyStatementMatchISIS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementMatchISIS.
func (in *RoutingPolicyStatementMatchISIS) DeepCopy() *RoutingPolicyStatementMatchISIS {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementMatchISIS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatementMatchOSPF) DeepCopyInto(out *RoutingPolicyStatementMatchOSPF) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatementMatchOSPF.
func (in *RoutingPolicyStatementMatchOSPF) DeepCopy() *RoutingPolicyStatementMatchOSPF {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatementMatchOSPF)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicyStatus) DeepCopyInto(out *RoutingPolicyStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicyStatus.
func (in *RoutingPolicyStatus) DeepCopy() *RoutingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubInterfaceIPv4) DeepCopyInto(out *SubInterfaceIPv4) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubInterfaceIPv4.
func (in *SubInterfaceIPv4) DeepCopy() *SubInterfaceIPv4 {
	if in == nil {
		return nil
	}
	out := new(SubInterfaceIPv4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubInterfaceIPv6) DeepCopyInto(out *SubInterfaceIPv6) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubInterfaceIPv6.
func (in *SubInterfaceIPv6) DeepCopy() *SubInterfaceIPv6 {
	if in == nil {
		return nil
	}
	out := new(SubInterfaceIPv6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubInterfacePeer) DeepCopyInto(out *SubInterfacePeer) {
	*out = *in
	in.PartitionEndpointID.DeepCopyInto(&out.PartitionEndpointID)
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(SubInterfaceIPv4)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(SubInterfaceIPv6)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubInterfacePeer.
func (in *SubInterfacePeer) DeepCopy() *SubInterfacePeer {
	if in == nil {
		return nil
	}
	out := new(SubInterfacePeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubInterfaceSpec) DeepCopyInto(out *SubInterfaceSpec) {
	*out = *in
	in.PartitionProviderEndpointID.DeepCopyInto(&out.PartitionProviderEndpointID)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(SubInterfaceIPv4)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(SubInterfaceIPv6)
		(*in).DeepCopyInto(*out)
	}
	if in.BFD != nil {
		in, out := &in.BFD, &out.BFD
		*out = new(corev1alpha1.BFDLinkParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Peer != nil {
		in, out := &in.Peer, &out.Peer
		*out = new(SubInterfacePeer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubInterfaceSpec.
func (in *SubInterfaceSpec) DeepCopy() *SubInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(SubInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubInterfaceStatus) DeepCopyInto(out *SubInterfaceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubInterfaceStatus.
func (in *SubInterfaceStatus) DeepCopy() *SubInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(SubInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelInterfaceSpec) DeepCopyInto(out *TunnelInterfaceSpec) {
	*out = *in
	out.PartitionProviderNodeID = in.PartitionProviderNodeID
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelInterfaceSpec.
func (in *TunnelInterfaceSpec) DeepCopy() *TunnelInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(TunnelInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelInterfaceSpecEthernet) DeepCopyInto(out *TunnelInterfaceSpecEthernet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelInterfaceSpecEthernet.
func (in *TunnelInterfaceSpecEthernet) DeepCopy() *TunnelInterfaceSpecEthernet {
	if in == nil {
		return nil
	}
	out := new(TunnelInterfaceSpecEthernet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelInterfaceStatus) DeepCopyInto(out *TunnelInterfaceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelInterfaceStatus.
func (in *TunnelInterfaceStatus) DeepCopy() *TunnelInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(TunnelInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelSubInterfaceSpec) DeepCopyInto(out *TunnelSubInterfaceSpec) {
	*out = *in
	in.PartitionProviderEndpointID.DeepCopyInto(&out.PartitionProviderEndpointID)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelSubInterfaceSpec.
func (in *TunnelSubInterfaceSpec) DeepCopy() *TunnelSubInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(TunnelSubInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelSubInterfaceStatus) DeepCopyInto(out *TunnelSubInterfaceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelSubInterfaceStatus.
func (in *TunnelSubInterfaceStatus) DeepCopy() *TunnelSubInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(TunnelSubInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}
