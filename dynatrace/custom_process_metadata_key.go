// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomProcessMetadataKey custom process metadata key
// swagger:model CustomProcessMetadataKey
type CustomProcessMetadataKey struct {

	// The actual key of the custom metadata.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Key *string `json:"key"`

	// The source of the custom metadata.
	// Required: true
	// Enum: [CLOUD_FOUNDRY ENVIRONMENT GOOGLE_CLOUD KUBERNETES PLUGIN]
	Source *string `json:"source"`
}

// Validate validates this custom process metadata key
func (m *CustomProcessMetadataKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomProcessMetadataKey) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", string(*m.Key), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", string(*m.Key), 256); err != nil {
		return err
	}

	return nil
}

var customProcessMetadataKeyTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLOUD_FOUNDRY","ENVIRONMENT","GOOGLE_CLOUD","KUBERNETES","PLUGIN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customProcessMetadataKeyTypeSourcePropEnum = append(customProcessMetadataKeyTypeSourcePropEnum, v)
	}
}

const (

	// CustomProcessMetadataKeySourceCLOUDFOUNDRY captures enum value "CLOUD_FOUNDRY"
	CustomProcessMetadataKeySourceCLOUDFOUNDRY string = "CLOUD_FOUNDRY"

	// CustomProcessMetadataKeySourceENVIRONMENT captures enum value "ENVIRONMENT"
	CustomProcessMetadataKeySourceENVIRONMENT string = "ENVIRONMENT"

	// CustomProcessMetadataKeySourceGOOGLECLOUD captures enum value "GOOGLE_CLOUD"
	CustomProcessMetadataKeySourceGOOGLECLOUD string = "GOOGLE_CLOUD"

	// CustomProcessMetadataKeySourceKUBERNETES captures enum value "KUBERNETES"
	CustomProcessMetadataKeySourceKUBERNETES string = "KUBERNETES"

	// CustomProcessMetadataKeySourcePLUGIN captures enum value "PLUGIN"
	CustomProcessMetadataKeySourcePLUGIN string = "PLUGIN"
)

// prop value enum
func (m *CustomProcessMetadataKey) validateSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customProcessMetadataKeyTypeSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomProcessMetadataKey) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomProcessMetadataKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomProcessMetadataKey) UnmarshalBinary(b []byte) error {
	var res CustomProcessMetadataKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
