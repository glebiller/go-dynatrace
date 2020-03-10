// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureCredentials Configuration of Azure credentials.
// swagger:model AzureCredentials
type AzureCredentials struct {

	// The monitoring is enabled (`true`) or disabled (`false`) for Azure credentials.
	//
	// If not set on creation, the `true` value is used.
	//
	// If the field is omitted during an update, the old value is used.
	Active bool `json:"active,omitempty"`

	// Id of an application (client) for which the Azure ServicePrinicipal credentials are created
	// The combination of appId and directoryId must be unique
	// Required: true
	// Read Only: true
	// Min Length: 1
	AppID string `json:"appId"`

	// Whether or not tags created in Azure should be automatically imported as tags in Dynatrace.
	// Required: true
	AutoTagging *bool `json:"autoTagging"`

	// Directory (Azure tenant) ID of the ServicePrincipal credentialsThe combination of appId and directoryId must be unique
	// Required: true
	// Read Only: true
	// Min Length: 1
	DirectoryID string `json:"directoryId"`

	// The ID of the AZURE_CREDENTIALS configuration.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The secretKey to the Azure ServicePrincipal App Registration.
	//
	// Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.
	//
	// If the field is omitted during an update, the old value is used.
	Key string `json:"key,omitempty"`

	// The name of the Azure credentials configuration.
	// The name must be unique.
	//
	// Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`.
	// Required: true
	// Min Length: 1
	// Pattern: ^([a-zA-Z0-9_ +-.]*)$
	Label *string `json:"label"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// A list of Azure tags to be monitored.
	//
	// Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`.
	//
	// You can specify up to 10 tags. A resource tagged with any of the TagPairs specified will be monitored. If value in TagPair is null, then resources with any value of this tag will be monitored.
	// Required: true
	// Max Items: 10
	// Min Items: 0
	// Unique: true
	MonitorOnlyTagPairs []*TagPair `json:"monitorOnlyTagPairs"`

	// Whether or not only resources which have the specified Azure tags (`true`) or all resources (`false`) should be monitored.
	// Required: true
	MonitorOnlyTaggedEntities *bool `json:"monitorOnlyTaggedEntities"`
}

// Validate validates this azure credentials
func (m *AzureCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoTagging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorOnlyTagPairs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorOnlyTaggedEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCredentials) validateAppID(formats strfmt.Registry) error {

	if err := validate.RequiredString("appId", "body", string(m.AppID)); err != nil {
		return err
	}

	if err := validate.MinLength("appId", "body", string(m.AppID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentials) validateAutoTagging(formats strfmt.Registry) error {

	if err := validate.Required("autoTagging", "body", m.AutoTagging); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentials) validateDirectoryID(formats strfmt.Registry) error {

	if err := validate.RequiredString("directoryId", "body", string(m.DirectoryID)); err != nil {
		return err
	}

	if err := validate.MinLength("directoryId", "body", string(m.DirectoryID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentials) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	if err := validate.MinLength("label", "body", string(*m.Label), 1); err != nil {
		return err
	}

	if err := validate.Pattern("label", "body", string(*m.Label), `^([a-zA-Z0-9_ +-.]*)$`); err != nil {
		return err
	}

	return nil
}

func (m *AzureCredentials) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCredentials) validateMonitorOnlyTagPairs(formats strfmt.Registry) error {

	if err := validate.Required("monitorOnlyTagPairs", "body", m.MonitorOnlyTagPairs); err != nil {
		return err
	}

	iMonitorOnlyTagPairsSize := int64(len(m.MonitorOnlyTagPairs))

	if err := validate.MinItems("monitorOnlyTagPairs", "body", iMonitorOnlyTagPairsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("monitorOnlyTagPairs", "body", iMonitorOnlyTagPairsSize, 10); err != nil {
		return err
	}

	if err := validate.UniqueItems("monitorOnlyTagPairs", "body", m.MonitorOnlyTagPairs); err != nil {
		return err
	}

	for i := 0; i < len(m.MonitorOnlyTagPairs); i++ {
		if swag.IsZero(m.MonitorOnlyTagPairs[i]) { // not required
			continue
		}

		if m.MonitorOnlyTagPairs[i] != nil {
			if err := m.MonitorOnlyTagPairs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monitorOnlyTagPairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureCredentials) validateMonitorOnlyTaggedEntities(formats strfmt.Registry) error {

	if err := validate.Required("monitorOnlyTaggedEntities", "body", m.MonitorOnlyTaggedEntities); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCredentials) UnmarshalBinary(b []byte) error {
	var res AzureCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
