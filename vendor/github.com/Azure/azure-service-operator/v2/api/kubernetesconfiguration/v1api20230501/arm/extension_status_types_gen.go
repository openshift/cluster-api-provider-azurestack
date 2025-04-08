// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// The Extension object.
type Extension_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: Identity of the Extension resource
	Identity *Identity_STATUS `json:"identity,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Plan: The plan information.
	Plan *Plan_STATUS `json:"plan,omitempty"`

	// Properties: Properties of an Extension resource
	Properties *Extension_Properties_STATUS `json:"properties,omitempty"`

	// SystemData: Top level metadata
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type Extension_Properties_STATUS struct {
	// AksAssignedIdentity: Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity *Extension_Properties_AksAssignedIdentity_STATUS `json:"aksAssignedIdentity,omitempty"`

	// AutoUpgradeMinorVersion: Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// ConfigurationProtectedSettings: Configuration settings that are sensitive, as name-value pairs for configuring this
	// extension.
	ConfigurationProtectedSettings map[string]string `json:"configurationProtectedSettings,omitempty"`

	// ConfigurationSettings: Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]string `json:"configurationSettings,omitempty"`

	// CurrentVersion: Currently installed version of the extension.
	CurrentVersion *string `json:"currentVersion,omitempty"`

	// CustomLocationSettings: Custom Location settings properties.
	CustomLocationSettings map[string]string `json:"customLocationSettings,omitempty"`

	// ErrorInfo: Error information from the Agent - e.g. errors during installation.
	ErrorInfo *ErrorDetail_STATUS `json:"errorInfo,omitempty"`

	// ExtensionType: Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types
	// registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType *string `json:"extensionType,omitempty"`

	// IsSystemExtension: Flag to note if this extension is a system extension
	IsSystemExtension *bool `json:"isSystemExtension,omitempty"`

	// PackageUri: Uri of the Helm package
	PackageUri *string `json:"packageUri,omitempty"`

	// ProvisioningState: Status of installation of this extension.
	ProvisioningState *ProvisioningStateDefinition_STATUS `json:"provisioningState,omitempty"`

	// ReleaseTrain: ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if
	// autoUpgradeMinorVersion is 'true'.
	ReleaseTrain *string `json:"releaseTrain,omitempty"`

	// Scope: Scope at which the extension is installed.
	Scope *Scope_STATUS `json:"scope,omitempty"`

	// Statuses: Status from this extension.
	Statuses []ExtensionStatus_STATUS `json:"statuses,omitempty"`

	// Version: User-specified version of the extension for this extension to 'pin'. To use 'version', autoUpgradeMinorVersion
	// must be 'false'.
	Version *string `json:"version,omitempty"`
}

// Identity for the resource.
type Identity_STATUS struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *Identity_Type_STATUS `json:"type,omitempty"`
}

// Plan for the resource.
type Plan_STATUS struct {
	// Name: A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`

	// Product: The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the
	// artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`

	// PromotionCode: A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// Publisher: The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`

	// Version: The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// The error detail.
type ErrorDetail_STATUS struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUS `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Details: The error details.
	Details []ErrorDetail_STATUS_Unrolled `json:"details,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}

type Extension_Properties_AksAssignedIdentity_STATUS struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *Extension_Properties_AksAssignedIdentity_Type_STATUS `json:"type,omitempty"`
}

// Status from the extension.
type ExtensionStatus_STATUS struct {
	// Code: Status code provided by the Extension
	Code *string `json:"code,omitempty"`

	// DisplayStatus: Short description of status of the extension.
	DisplayStatus *string `json:"displayStatus,omitempty"`

	// Level: Level of the status.
	Level *ExtensionStatus_Level_STATUS `json:"level,omitempty"`

	// Message: Detailed message of the status from the Extension.
	Message *string `json:"message,omitempty"`

	// Time: DateLiteral (per ISO8601) noting the time of installation status.
	Time *string `json:"time,omitempty"`
}

type Identity_Type_STATUS string

const Identity_Type_STATUS_SystemAssigned = Identity_Type_STATUS("SystemAssigned")

// Mapping from string to Identity_Type_STATUS
var identity_Type_STATUS_Values = map[string]Identity_Type_STATUS{
	"systemassigned": Identity_Type_STATUS_SystemAssigned,
}

// The provisioning state of the resource.
type ProvisioningStateDefinition_STATUS string

const (
	ProvisioningStateDefinition_STATUS_Canceled  = ProvisioningStateDefinition_STATUS("Canceled")
	ProvisioningStateDefinition_STATUS_Creating  = ProvisioningStateDefinition_STATUS("Creating")
	ProvisioningStateDefinition_STATUS_Deleting  = ProvisioningStateDefinition_STATUS("Deleting")
	ProvisioningStateDefinition_STATUS_Failed    = ProvisioningStateDefinition_STATUS("Failed")
	ProvisioningStateDefinition_STATUS_Succeeded = ProvisioningStateDefinition_STATUS("Succeeded")
	ProvisioningStateDefinition_STATUS_Updating  = ProvisioningStateDefinition_STATUS("Updating")
)

// Mapping from string to ProvisioningStateDefinition_STATUS
var provisioningStateDefinition_STATUS_Values = map[string]ProvisioningStateDefinition_STATUS{
	"canceled":  ProvisioningStateDefinition_STATUS_Canceled,
	"creating":  ProvisioningStateDefinition_STATUS_Creating,
	"deleting":  ProvisioningStateDefinition_STATUS_Deleting,
	"failed":    ProvisioningStateDefinition_STATUS_Failed,
	"succeeded": ProvisioningStateDefinition_STATUS_Succeeded,
	"updating":  ProvisioningStateDefinition_STATUS_Updating,
}

// Scope of the extension. It can be either Cluster or Namespace; but not both.
type Scope_STATUS struct {
	// Cluster: Specifies that the scope of the extension is Cluster
	Cluster *ScopeCluster_STATUS `json:"cluster,omitempty"`

	// Namespace: Specifies that the scope of the extension is Namespace
	Namespace *ScopeNamespace_STATUS `json:"namespace,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

// The resource management error additional info.
type ErrorAdditionalInfo_STATUS struct {
	// Info: The additional info.
	Info map[string]v1.JSON `json:"info,omitempty"`

	// Type: The additional info type.
	Type *string `json:"type,omitempty"`
}

type ErrorDetail_STATUS_Unrolled struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUS `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}

type Extension_Properties_AksAssignedIdentity_Type_STATUS string

const (
	Extension_Properties_AksAssignedIdentity_Type_STATUS_SystemAssigned = Extension_Properties_AksAssignedIdentity_Type_STATUS("SystemAssigned")
	Extension_Properties_AksAssignedIdentity_Type_STATUS_UserAssigned   = Extension_Properties_AksAssignedIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to Extension_Properties_AksAssignedIdentity_Type_STATUS
var extension_Properties_AksAssignedIdentity_Type_STATUS_Values = map[string]Extension_Properties_AksAssignedIdentity_Type_STATUS{
	"systemassigned": Extension_Properties_AksAssignedIdentity_Type_STATUS_SystemAssigned,
	"userassigned":   Extension_Properties_AksAssignedIdentity_Type_STATUS_UserAssigned,
}

type ExtensionStatus_Level_STATUS string

const (
	ExtensionStatus_Level_STATUS_Error       = ExtensionStatus_Level_STATUS("Error")
	ExtensionStatus_Level_STATUS_Information = ExtensionStatus_Level_STATUS("Information")
	ExtensionStatus_Level_STATUS_Warning     = ExtensionStatus_Level_STATUS("Warning")
)

// Mapping from string to ExtensionStatus_Level_STATUS
var extensionStatus_Level_STATUS_Values = map[string]ExtensionStatus_Level_STATUS{
	"error":       ExtensionStatus_Level_STATUS_Error,
	"information": ExtensionStatus_Level_STATUS_Information,
	"warning":     ExtensionStatus_Level_STATUS_Warning,
}

// Specifies that the scope of the extension is Cluster
type ScopeCluster_STATUS struct {
	// ReleaseNamespace: Namespace where the extension Release must be placed, for a Cluster scoped extension.  If this
	// namespace does not exist, it will be created
	ReleaseNamespace *string `json:"releaseNamespace,omitempty"`
}

// Specifies that the scope of the extension is Namespace
type ScopeNamespace_STATUS struct {
	// TargetNamespace: Namespace where the extension will be created for an Namespace scoped extension.  If this namespace
	// does not exist, it will be created
	TargetNamespace *string `json:"targetNamespace,omitempty"`
}
