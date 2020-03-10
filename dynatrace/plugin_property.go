// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PluginProperty A property of a plugin.
// swagger:model PluginProperty
type PluginProperty struct {

	// The default value of the property.
	DefaultValue string `json:"defaultValue,omitempty"`

	// The list of possible values of the property.
	//
	//  If such a list is defined, only values from this list can be assigned to the property.
	DropdownValues []string `json:"dropdownValues"`

	// The key of the property.
	Key string `json:"key,omitempty"`

	// The type of the property.
	Type string `json:"type,omitempty"`
}

// Validate validates this plugin property
func (m *PluginProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginProperty) UnmarshalBinary(b []byte) error {
	var res PluginProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
