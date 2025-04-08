// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
)

// Storage version of v1api20240402preview.AdvancedNetworking_STATUS
// Advanced Networking profile for enabling observability on a cluster. Note that enabling advanced networking features may
// incur additional costs. For more information see aka.ms/aksadvancednetworking.
type AdvancedNetworking_STATUS struct {
	Observability *AdvancedNetworkingObservability_STATUS `json:"observability,omitempty"`
	PropertyBag   genruntime.PropertyBag                  `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AdvancedNetworking_STATUS populates our AdvancedNetworking_STATUS from the provided source AdvancedNetworking_STATUS
func (networking *AdvancedNetworking_STATUS) AssignProperties_From_AdvancedNetworking_STATUS(source *storage.AdvancedNetworking_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Enabled
	if source.Enabled != nil {
		propertyBag.Add("Enabled", *source.Enabled)
	} else {
		propertyBag.Remove("Enabled")
	}

	// Observability
	if source.Observability != nil {
		var observability AdvancedNetworkingObservability_STATUS
		err := observability.AssignProperties_From_AdvancedNetworkingObservability_STATUS(source.Observability)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AdvancedNetworkingObservability_STATUS() to populate field Observability")
		}
		networking.Observability = &observability
	} else {
		networking.Observability = nil
	}

	// Security
	if source.Security != nil {
		propertyBag.Add("Security", *source.Security)
	} else {
		propertyBag.Remove("Security")
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		networking.PropertyBag = propertyBag
	} else {
		networking.PropertyBag = nil
	}

	// Invoke the augmentConversionForAdvancedNetworking_STATUS interface (if implemented) to customize the conversion
	var networkingAsAny any = networking
	if augmentedNetworking, ok := networkingAsAny.(augmentConversionForAdvancedNetworking_STATUS); ok {
		err := augmentedNetworking.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AdvancedNetworking_STATUS populates the provided destination AdvancedNetworking_STATUS from our AdvancedNetworking_STATUS
func (networking *AdvancedNetworking_STATUS) AssignProperties_To_AdvancedNetworking_STATUS(destination *storage.AdvancedNetworking_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(networking.PropertyBag)

	// Enabled
	if propertyBag.Contains("Enabled") {
		var enabled bool
		err := propertyBag.Pull("Enabled", &enabled)
		if err != nil {
			return errors.Wrap(err, "pulling 'Enabled' from propertyBag")
		}

		destination.Enabled = &enabled
	} else {
		destination.Enabled = nil
	}

	// Observability
	if networking.Observability != nil {
		var observability storage.AdvancedNetworkingObservability_STATUS
		err := networking.Observability.AssignProperties_To_AdvancedNetworkingObservability_STATUS(&observability)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AdvancedNetworkingObservability_STATUS() to populate field Observability")
		}
		destination.Observability = &observability
	} else {
		destination.Observability = nil
	}

	// Security
	if propertyBag.Contains("Security") {
		var security storage.AdvancedNetworkingSecurity_STATUS
		err := propertyBag.Pull("Security", &security)
		if err != nil {
			return errors.Wrap(err, "pulling 'Security' from propertyBag")
		}

		destination.Security = &security
	} else {
		destination.Security = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAdvancedNetworking_STATUS interface (if implemented) to customize the conversion
	var networkingAsAny any = networking
	if augmentedNetworking, ok := networkingAsAny.(augmentConversionForAdvancedNetworking_STATUS); ok {
		err := augmentedNetworking.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20240402preview.AdvancedNetworkingObservability_STATUS
// Observability profile to enable advanced network metrics and flow logs with historical contexts.
type AdvancedNetworkingObservability_STATUS struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AdvancedNetworkingObservability_STATUS populates our AdvancedNetworkingObservability_STATUS from the provided source AdvancedNetworkingObservability_STATUS
func (observability *AdvancedNetworkingObservability_STATUS) AssignProperties_From_AdvancedNetworkingObservability_STATUS(source *storage.AdvancedNetworkingObservability_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Enabled
	if source.Enabled != nil {
		enabled := *source.Enabled
		observability.Enabled = &enabled
	} else {
		observability.Enabled = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		observability.PropertyBag = propertyBag
	} else {
		observability.PropertyBag = nil
	}

	// Invoke the augmentConversionForAdvancedNetworkingObservability_STATUS interface (if implemented) to customize the conversion
	var observabilityAsAny any = observability
	if augmentedObservability, ok := observabilityAsAny.(augmentConversionForAdvancedNetworkingObservability_STATUS); ok {
		err := augmentedObservability.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_AdvancedNetworkingObservability_STATUS populates the provided destination AdvancedNetworkingObservability_STATUS from our AdvancedNetworkingObservability_STATUS
func (observability *AdvancedNetworkingObservability_STATUS) AssignProperties_To_AdvancedNetworkingObservability_STATUS(destination *storage.AdvancedNetworkingObservability_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(observability.PropertyBag)

	// Enabled
	if observability.Enabled != nil {
		enabled := *observability.Enabled
		destination.Enabled = &enabled
	} else {
		destination.Enabled = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForAdvancedNetworkingObservability_STATUS interface (if implemented) to customize the conversion
	var observabilityAsAny any = observability
	if augmentedObservability, ok := observabilityAsAny.(augmentConversionForAdvancedNetworkingObservability_STATUS); ok {
		err := augmentedObservability.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForAdvancedNetworking_STATUS interface {
	AssignPropertiesFrom(src *storage.AdvancedNetworking_STATUS) error
	AssignPropertiesTo(dst *storage.AdvancedNetworking_STATUS) error
}

type augmentConversionForAdvancedNetworkingObservability_STATUS interface {
	AssignPropertiesFrom(src *storage.AdvancedNetworkingObservability_STATUS) error
	AssignPropertiesTo(dst *storage.AdvancedNetworkingObservability_STATUS) error
}
