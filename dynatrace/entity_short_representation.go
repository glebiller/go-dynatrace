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

// EntityShortRepresentation The short representation of a Dynatrace entity.
// swagger:model EntityShortRepresentation
type EntityShortRepresentation struct {

	// A short description of the Dynatrace entity.
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The ID of the Dynatrace entity.
	// Required: true
	ID *string `json:"id"`

	// The name of the Dynatrace entity.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this entity short representation
func (m *EntityShortRepresentation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityShortRepresentation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityShortRepresentation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityShortRepresentation) UnmarshalBinary(b []byte) error {
	var res EntityShortRepresentation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
