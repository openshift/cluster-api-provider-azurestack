// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DnsForwardingRuleSetsForwardingRule_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the forwarding rule.
	Properties *ForwardingRuleProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DnsForwardingRuleSetsForwardingRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (rule DnsForwardingRuleSetsForwardingRule_Spec) GetAPIVersion() string {
	return "2022-07-01"
}

// GetName returns the Name of the resource
func (rule *DnsForwardingRuleSetsForwardingRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsForwardingRulesets/forwardingRules"
func (rule *DnsForwardingRuleSetsForwardingRule_Spec) GetType() string {
	return "Microsoft.Network/dnsForwardingRulesets/forwardingRules"
}

// Represents the properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRuleProperties struct {
	// DomainName: The domain name for the forwarding rule.
	DomainName *string `json:"domainName,omitempty"`

	// ForwardingRuleState: The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleProperties_ForwardingRuleState `json:"forwardingRuleState,omitempty"`

	// Metadata: Metadata attached to the forwarding rule.
	Metadata map[string]string `json:"metadata,omitempty"`

	// TargetDnsServers: DNS servers to forward the DNS query to.
	TargetDnsServers []TargetDnsServer `json:"targetDnsServers,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ForwardingRuleProperties_ForwardingRuleState string

const (
	ForwardingRuleProperties_ForwardingRuleState_Disabled = ForwardingRuleProperties_ForwardingRuleState("Disabled")
	ForwardingRuleProperties_ForwardingRuleState_Enabled  = ForwardingRuleProperties_ForwardingRuleState("Enabled")
)

// Mapping from string to ForwardingRuleProperties_ForwardingRuleState
var forwardingRuleProperties_ForwardingRuleState_Values = map[string]ForwardingRuleProperties_ForwardingRuleState{
	"disabled": ForwardingRuleProperties_ForwardingRuleState_Disabled,
	"enabled":  ForwardingRuleProperties_ForwardingRuleState_Enabled,
}

// Describes a server to forward the DNS queries to.
type TargetDnsServer struct {
	// IpAddress: DNS server IP address.
	IpAddress *string `json:"ipAddress,omitempty" optionalConfigMapPair:"IpAddress"`

	// Port: DNS server port.
	Port *int `json:"port,omitempty"`
}
