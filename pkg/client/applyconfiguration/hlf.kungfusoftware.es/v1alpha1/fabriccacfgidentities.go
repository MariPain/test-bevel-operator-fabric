/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricCACFGIdentitiesApplyConfiguration represents an declarative configuration of the FabricCACFGIdentities type for use
// with apply.
type FabricCACFGIdentitiesApplyConfiguration struct {
	AllowRemove *bool `json:"allowRemove,omitempty"`
}

// FabricCACFGIdentitiesApplyConfiguration constructs an declarative configuration of the FabricCACFGIdentities type for use with
// apply.
func FabricCACFGIdentities() *FabricCACFGIdentitiesApplyConfiguration {
	return &FabricCACFGIdentitiesApplyConfiguration{}
}

// WithAllowRemove sets the AllowRemove field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowRemove field is set to the value of the last call.
func (b *FabricCACFGIdentitiesApplyConfiguration) WithAllowRemove(value bool) *FabricCACFGIdentitiesApplyConfiguration {
	b.AllowRemove = &value
	return b
}