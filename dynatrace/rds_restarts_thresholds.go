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

// RdsRestartsThresholds Custom thresholds for restarts sequence on RDS. If not set, automatic mode is used.
// swagger:model RdsRestartsThresholds
type RdsRestartsThresholds struct {

	// Alert if number of restarts is *X* per minute or higher in 3 out of 20 samples.
	// Required: true
	// Maximum: 100
	// Minimum: 1
	RestartsPerMinute *int32 `json:"restartsPerMinute"`
}

// Validate validates this rds restarts thresholds
func (m *RdsRestartsThresholds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestartsPerMinute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsRestartsThresholds) validateRestartsPerMinute(formats strfmt.Registry) error {

	if err := validate.Required("restartsPerMinute", "body", m.RestartsPerMinute); err != nil {
		return err
	}

	if err := validate.MinimumInt("restartsPerMinute", "body", int64(*m.RestartsPerMinute), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("restartsPerMinute", "body", int64(*m.RestartsPerMinute), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsRestartsThresholds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsRestartsThresholds) UnmarshalBinary(b []byte) error {
	var res RdsRestartsThresholds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
