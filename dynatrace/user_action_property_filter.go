// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserActionPropertyFilter user action property filter
// swagger:model UserActionPropertyFilter
type UserActionPropertyFilter struct {

	// Filter by values >= this value (only for properties of numeric type)
	From float64 `json:"from,omitempty"`

	// The key of the user action property
	Key string `json:"key,omitempty"`

	// Filter by values <= this value (only for properties of numeric type)
	To float64 `json:"to,omitempty"`

	// Filter by a string value (only for properties of type string)
	Value string `json:"value,omitempty"`
}

// Validate validates this user action property filter
func (m *UserActionPropertyFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserActionPropertyFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserActionPropertyFilter) UnmarshalBinary(b []byte) error {
	var res UserActionPropertyFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
