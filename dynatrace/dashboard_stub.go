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

// DashboardStub A short representation of a dashboard.
// swagger:model DashboardStub
type DashboardStub struct {

	// The ID of the dashboard.
	// Required: true
	ID *string `json:"id"`

	// The name of the dashboard.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The owner of the dashboard.
	// Read Only: true
	Owner string `json:"owner,omitempty"`
}

// Validate validates this dashboard stub
func (m *DashboardStub) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardStub) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardStub) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardStub) UnmarshalBinary(b []byte) error {
	var res DashboardStub
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
