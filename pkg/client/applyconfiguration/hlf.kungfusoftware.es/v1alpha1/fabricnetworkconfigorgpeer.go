/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricNetworkConfigOrgPeerApplyConfiguration represents an declarative configuration of the FabricNetworkConfigOrgPeer type for use
// with apply.
type FabricNetworkConfigOrgPeerApplyConfiguration struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// FabricNetworkConfigOrgPeerApplyConfiguration constructs an declarative configuration of the FabricNetworkConfigOrgPeer type for use with
// apply.
func FabricNetworkConfigOrgPeer() *FabricNetworkConfigOrgPeerApplyConfiguration {
	return &FabricNetworkConfigOrgPeerApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *FabricNetworkConfigOrgPeerApplyConfiguration) WithName(value string) *FabricNetworkConfigOrgPeerApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *FabricNetworkConfigOrgPeerApplyConfiguration) WithNamespace(value string) *FabricNetworkConfigOrgPeerApplyConfiguration {
	b.Namespace = &value
	return b
}