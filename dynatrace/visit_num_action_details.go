// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VisitNumActionDetails Configuration of a number of user actions-based conversion goal.
// swagger:model VisitNumActionDetails
type VisitNumActionDetails struct {

	// The number of user actions to hit the conversion goal.
	NumUserActions int32 `json:"numUserActions,omitempty"`
}

// Validate validates this visit num action details
func (m *VisitNumActionDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VisitNumActionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisitNumActionDetails) UnmarshalBinary(b []byte) error {
	var res VisitNumActionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}