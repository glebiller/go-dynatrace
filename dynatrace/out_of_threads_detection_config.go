// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OutOfThreadsDetectionConfig Configuration of Java out of threads problems detection.
// swagger:model OutOfThreadsDetectionConfig
type OutOfThreadsDetectionConfig struct {

	// Custom thresholds for Java out of threads detection. If not set, automatic mode is used.
	CustomThresholds *OutOfThreadsThresholds `json:"customThresholds,omitempty"`

	// The detection is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this out of threads detection config
func (m *OutOfThreadsDetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomThresholds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfThreadsDetectionConfig) validateCustomThresholds(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomThresholds) { // not required
		return nil
	}

	if m.CustomThresholds != nil {
		if err := m.CustomThresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customThresholds")
			}
			return err
		}
	}

	return nil
}

func (m *OutOfThreadsDetectionConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutOfThreadsDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfThreadsDetectionConfig) UnmarshalBinary(b []byte) error {
	var res OutOfThreadsDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
