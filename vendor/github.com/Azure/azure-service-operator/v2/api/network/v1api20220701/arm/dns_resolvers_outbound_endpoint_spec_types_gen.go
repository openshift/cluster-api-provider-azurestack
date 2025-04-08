// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsResolversOutboundEndpoint_Spec struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the outbound endpoint.
	Properties *OutboundEndpointProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsResolversOutboundEndpoint_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint DnsResolversOutboundEndpoint_Spec) GetAPIVersion() string {
	return "2022-07-01"
}

// GetName returns the Name of the resource
func (endpoint *DnsResolversOutboundEndpoint_Spec) GetName() string {
	return endpoint.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers/outboundEndpoints"
func (endpoint *DnsResolversOutboundEndpoint_Spec) GetType() string {
	return "Microsoft.Network/dnsResolvers/outboundEndpoints"
}

// Represents the properties of an outbound endpoint for a DNS resolver.
type OutboundEndpointProperties struct {
	// Subnet: The reference to the subnet used for the outbound endpoint.
	Subnet *SubResource `json:"subnet,omitempty"`
}
