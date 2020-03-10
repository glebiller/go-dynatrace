// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DashboardFilter Filters, applied to a dashboard.
// swagger:model DashboardFilter
type DashboardFilter struct {

	// The management zone to which the dashboard belongs.
	ManagementZone *EntityShortRepresentation `json:"managementZone,omitempty"`

	// The default timeframe of the dashboard.
	Timeframe string `json:"timeframe,omitempty"`
}

// Validate validates this dashboard filter
func (m *DashboardFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagementZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardFilter) validateManagementZone(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagementZone) { // not required
		return nil
	}

	if m.ManagementZone != nil {
		if err := m.ManagementZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementZone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardFilter) UnmarshalBinary(b []byte) error {
	var res DashboardFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}