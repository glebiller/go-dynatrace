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

// RemotePluginEndpoint Configuration of a plugin endpoint.
// swagger:model RemotePluginEndpoint
type RemotePluginEndpoint struct {

	// The ActiveGate plugin module that hosts the endpoint.
	// Required: true
	ActiveGatePluginModule *EntityShortRepresentation `json:"activeGatePluginModule"`

	// The endpoint is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled,omitempty"`

	// The ID of the endpoint.
	ID string `json:"id,omitempty"`

	// The name of the endpoint, displayed in Dynatrace.
	Name string `json:"name,omitempty"`

	// The ID of the plugin to which the endpoint belongs.
	PluginID string `json:"pluginId,omitempty"`

	// The list of endpoint parameters.
	//
	//  Each parameter is a property-value pair.
	Properties map[string]string `json:"properties,omitempty"`
}

// Validate validates this remote plugin endpoint
func (m *RemotePluginEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveGatePluginModule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemotePluginEndpoint) validateActiveGatePluginModule(formats strfmt.Registry) error {

	if err := validate.Required("activeGatePluginModule", "body", m.ActiveGatePluginModule); err != nil {
		return err
	}

	if m.ActiveGatePluginModule != nil {
		if err := m.ActiveGatePluginModule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeGatePluginModule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemotePluginEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemotePluginEndpoint) UnmarshalBinary(b []byte) error {
	var res RemotePluginEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}