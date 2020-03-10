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

// EsxiHighCPUSaturationConfig The configuration of the CPU saturation on ESXi host detection.
// swagger:model EsxiHighCpuSaturationConfig
type EsxiHighCPUSaturationConfig struct {

	// Custom thresholds for CPU saturation detection on ESXi. If not set then the automatic mode is used.
	//
	//  **All** conditions must be fulfilled to trigger an alert.
	CustomThresholds *EsxiHighCPUThresholds `json:"customThresholds,omitempty"`

	// The detection is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this esxi high Cpu saturation config
func (m *EsxiHighCPUSaturationConfig) Validate(formats strfmt.Registry) error {
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

func (m *EsxiHighCPUSaturationConfig) validateCustomThresholds(formats strfmt.Registry) error {

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

func (m *EsxiHighCPUSaturationConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EsxiHighCPUSaturationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EsxiHighCPUSaturationConfig) UnmarshalBinary(b []byte) error {
	var res EsxiHighCPUSaturationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}