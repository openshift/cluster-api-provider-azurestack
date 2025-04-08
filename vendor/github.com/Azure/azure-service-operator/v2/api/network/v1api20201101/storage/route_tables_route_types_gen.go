// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20201101.RouteTablesRoute
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}
type RouteTablesRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTablesRoute_Spec   `json:"spec,omitempty"`
	Status            RouteTablesRoute_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTablesRoute{}

// GetConditions returns the conditions of the resource
func (route *RouteTablesRoute) GetConditions() conditions.Conditions {
	return route.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (route *RouteTablesRoute) SetConditions(conditions conditions.Conditions) {
	route.Status.Conditions = conditions
}

var _ conversion.Convertible = &RouteTablesRoute{}

// ConvertFrom populates our RouteTablesRoute from the provided hub RouteTablesRoute
func (route *RouteTablesRoute) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1api20240301/storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_From_RouteTablesRoute(source)
}

// ConvertTo populates the provided hub RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1api20240301/storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_To_RouteTablesRoute(destination)
}

var _ configmaps.Exporter = &RouteTablesRoute{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (route *RouteTablesRoute) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if route.Spec.OperatorSpec == nil {
		return nil
	}
	return route.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RouteTablesRoute{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (route *RouteTablesRoute) SecretDestinationExpressions() []*core.DestinationExpression {
	if route.Spec.OperatorSpec == nil {
		return nil
	}
	return route.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (route *RouteTablesRoute) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (route RouteTablesRoute) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceScope returns the scope of the resource
func (route *RouteTablesRoute) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (route *RouteTablesRoute) GetSpec() genruntime.ConvertibleSpec {
	return &route.Spec
}

// GetStatus returns the status of this resource
func (route *RouteTablesRoute) GetStatus() genruntime.ConvertibleStatus {
	return &route.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (route *RouteTablesRoute) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (route *RouteTablesRoute) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// NewEmptyStatus returns a new empty (blank) status
func (route *RouteTablesRoute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTablesRoute_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (route *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return route.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (route *RouteTablesRoute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTablesRoute_STATUS); ok {
		route.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTablesRoute_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// AssignProperties_From_RouteTablesRoute populates our RouteTablesRoute from the provided source RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_From_RouteTablesRoute(source *storage.RouteTablesRoute) error {

	// ObjectMeta
	route.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTablesRoute_Spec
	err := spec.AssignProperties_From_RouteTablesRoute_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTablesRoute_Spec() to populate field Spec")
	}
	route.Spec = spec

	// Status
	var status RouteTablesRoute_STATUS
	err = status.AssignProperties_From_RouteTablesRoute_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTablesRoute_STATUS() to populate field Status")
	}
	route.Status = status

	// Invoke the augmentConversionForRouteTablesRoute interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute populates the provided destination RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_To_RouteTablesRoute(destination *storage.RouteTablesRoute) error {

	// ObjectMeta
	destination.ObjectMeta = *route.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.RouteTablesRoute_Spec
	err := route.Spec.AssignProperties_To_RouteTablesRoute_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTablesRoute_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.RouteTablesRoute_STATUS
	err = route.Status.AssignProperties_To_RouteTablesRoute_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTablesRoute_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRouteTablesRoute interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (route *RouteTablesRoute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: route.Spec.OriginalVersion,
		Kind:    "RouteTablesRoute",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20201101.RouteTablesRoute
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

type augmentConversionForRouteTablesRoute interface {
	AssignPropertiesFrom(src *storage.RouteTablesRoute) error
	AssignPropertiesTo(dst *storage.RouteTablesRoute) error
}

// Storage version of v1api20201101.RouteTablesRoute_Spec
type RouteTablesRoute_Spec struct {
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string                        `json:"azureName,omitempty"`
	NextHopIpAddress *string                       `json:"nextHopIpAddress,omitempty"`
	NextHopType      *string                       `json:"nextHopType,omitempty"`
	OperatorSpec     *RouteTablesRouteOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion  string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/RouteTable resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"RouteTable"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RouteTablesRoute_Spec{}

// ConvertSpecFrom populates our RouteTablesRoute_Spec from the provided source
func (route *RouteTablesRoute_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.RouteTablesRoute_Spec)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTablesRoute_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.RouteTablesRoute_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTablesRoute_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.RouteTablesRoute_Spec)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTablesRoute_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RouteTablesRoute_Spec{}
	err := route.AssignProperties_To_RouteTablesRoute_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RouteTablesRoute_Spec populates our RouteTablesRoute_Spec from the provided source RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignProperties_From_RouteTablesRoute_Spec(source *storage.RouteTablesRoute_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// AzureName
	route.AzureName = source.AzureName

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	route.NextHopType = genruntime.ClonePointerToString(source.NextHopType)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec RouteTablesRouteOperatorSpec
		err := operatorSpec.AssignProperties_From_RouteTablesRouteOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_RouteTablesRouteOperatorSpec() to populate field OperatorSpec")
		}
		route.OperatorSpec = &operatorSpec
	} else {
		route.OperatorSpec = nil
	}

	// OriginalVersion
	route.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		route.Owner = &owner
	} else {
		route.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		route.PropertyBag = propertyBag
	} else {
		route.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRoute_Spec interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute_Spec); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute_Spec populates the provided destination RouteTablesRoute_Spec from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignProperties_To_RouteTablesRoute_Spec(destination *storage.RouteTablesRoute_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(route.PropertyBag)

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(route.AddressPrefix)

	// AzureName
	destination.AzureName = route.AzureName

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(route.NextHopIpAddress)

	// NextHopType
	destination.NextHopType = genruntime.ClonePointerToString(route.NextHopType)

	// OperatorSpec
	if route.OperatorSpec != nil {
		var operatorSpec storage.RouteTablesRouteOperatorSpec
		err := route.OperatorSpec.AssignProperties_To_RouteTablesRouteOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_RouteTablesRouteOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = route.OriginalVersion

	// Owner
	if route.Owner != nil {
		owner := route.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRoute_Spec interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute_Spec); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20201101.RouteTablesRoute_STATUS
type RouteTablesRoute_STATUS struct {
	AddressPrefix     *string                `json:"addressPrefix,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	HasBgpOverride    *bool                  `json:"hasBgpOverride,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NextHopIpAddress  *string                `json:"nextHopIpAddress,omitempty"`
	NextHopType       *string                `json:"nextHopType,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTablesRoute_STATUS{}

// ConvertStatusFrom populates our RouteTablesRoute_STATUS from the provided source
func (route *RouteTablesRoute_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.RouteTablesRoute_STATUS)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTablesRoute_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.RouteTablesRoute_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTablesRoute_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.RouteTablesRoute_STATUS)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTablesRoute_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RouteTablesRoute_STATUS{}
	err := route.AssignProperties_To_RouteTablesRoute_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_RouteTablesRoute_STATUS populates our RouteTablesRoute_STATUS from the provided source RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignProperties_From_RouteTablesRoute_STATUS(source *storage.RouteTablesRoute_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// Conditions
	route.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	route.Etag = genruntime.ClonePointerToString(source.Etag)

	// HasBgpOverride
	if source.HasBgpOverride != nil {
		hasBgpOverride := *source.HasBgpOverride
		route.HasBgpOverride = &hasBgpOverride
	} else {
		route.HasBgpOverride = nil
	}

	// Id
	route.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	route.Name = genruntime.ClonePointerToString(source.Name)

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	route.NextHopType = genruntime.ClonePointerToString(source.NextHopType)

	// ProvisioningState
	route.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// Type
	route.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		route.PropertyBag = propertyBag
	} else {
		route.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRoute_STATUS interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute_STATUS); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute_STATUS populates the provided destination RouteTablesRoute_STATUS from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignProperties_To_RouteTablesRoute_STATUS(destination *storage.RouteTablesRoute_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(route.PropertyBag)

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(route.AddressPrefix)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(route.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(route.Etag)

	// HasBgpOverride
	if route.HasBgpOverride != nil {
		hasBgpOverride := *route.HasBgpOverride
		destination.HasBgpOverride = &hasBgpOverride
	} else {
		destination.HasBgpOverride = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(route.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(route.Name)

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(route.NextHopIpAddress)

	// NextHopType
	destination.NextHopType = genruntime.ClonePointerToString(route.NextHopType)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(route.ProvisioningState)

	// Type
	destination.Type = genruntime.ClonePointerToString(route.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRoute_STATUS interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute_STATUS); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRouteTablesRoute_Spec interface {
	AssignPropertiesFrom(src *storage.RouteTablesRoute_Spec) error
	AssignPropertiesTo(dst *storage.RouteTablesRoute_Spec) error
}

type augmentConversionForRouteTablesRoute_STATUS interface {
	AssignPropertiesFrom(src *storage.RouteTablesRoute_STATUS) error
	AssignPropertiesTo(dst *storage.RouteTablesRoute_STATUS) error
}

// Storage version of v1api20201101.RouteTablesRouteOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RouteTablesRouteOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_RouteTablesRouteOperatorSpec populates our RouteTablesRouteOperatorSpec from the provided source RouteTablesRouteOperatorSpec
func (operator *RouteTablesRouteOperatorSpec) AssignProperties_From_RouteTablesRouteOperatorSpec(source *storage.RouteTablesRouteOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRouteOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForRouteTablesRouteOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRouteOperatorSpec populates the provided destination RouteTablesRouteOperatorSpec from our RouteTablesRouteOperatorSpec
func (operator *RouteTablesRouteOperatorSpec) AssignProperties_To_RouteTablesRouteOperatorSpec(destination *storage.RouteTablesRouteOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTablesRouteOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForRouteTablesRouteOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRouteTablesRouteOperatorSpec interface {
	AssignPropertiesFrom(src *storage.RouteTablesRouteOperatorSpec) error
	AssignPropertiesTo(dst *storage.RouteTablesRouteOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
