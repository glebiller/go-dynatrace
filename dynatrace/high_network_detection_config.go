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

// HighNetworkDetectionConfig Configuration of high network utilization detection.
// swagger:model HighNetworkDetectionConfig
type HighNetworkDetectionConfig struct {

	// Custom thresholds for high network utilization. If not set, automatic mode is used.
	CustomThresholds *HighNetworkThresholds `json:"customThresholds,omitempty"`

	// The detection is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this high network detection config
func (m *HighNetworkDetectionConfig) Validate(formats strfmt.Registry) error {
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

func (m *HighNetworkDetectionConfig) validateCustomThresholds(formats strfmt.Registry) error {

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

func (m *HighNetworkDetectionConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HighNetworkDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HighNetworkDetectionConfig) UnmarshalBinary(b []byte) error {
	var res HighNetworkDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
