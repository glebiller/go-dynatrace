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

// LambdaHighErrorRateThresholds Custom thresholds for AWS Lambda high error rate. If not set, automatic mode is used.
// swagger:model LambdaHighErrorRateThresholds
type LambdaHighErrorRateThresholds struct {

	// Alert if failed invocations rate is higher than *X*% in 3 out of 5 samples.
	// Required: true
	// Maximum: 100
	// Minimum: 0
	FailedInvocationsRate *int32 `json:"failedInvocationsRate"`
}

// Validate validates this lambda high error rate thresholds
func (m *LambdaHighErrorRateThresholds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedInvocationsRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LambdaHighErrorRateThresholds) validateFailedInvocationsRate(formats strfmt.Registry) error {

	if err := validate.Required("failedInvocationsRate", "body", m.FailedInvocationsRate); err != nil {
		return err
	}

	if err := validate.MinimumInt("failedInvocationsRate", "body", int64(*m.FailedInvocationsRate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("failedInvocationsRate", "body", int64(*m.FailedInvocationsRate), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LambdaHighErrorRateThresholds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LambdaHighErrorRateThresholds) UnmarshalBinary(b []byte) error {
	var res LambdaHighErrorRateThresholds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}