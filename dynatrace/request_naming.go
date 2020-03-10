// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestNaming The request naming rule.
// swagger:model RequestNaming
type RequestNaming struct {

	// The set of conditions for the request naming rule usage.
	//
	//  You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger.
	// Required: true
	Conditions []*Condition `json:"conditions"`

	// The rule is enabled (`true`) or disabled (`false`).
	// Required: true
	Enabled *bool `json:"enabled"`

	// The ID of the request naming rule.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Specifies the management zones for which this rule should be applied.
	// Max Items: 1
	// Min Items: 0
	ManagementZones []string `json:"managementZones"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// The name to be assigned to matching requests.
	// Required: true
	// Max Length: 1000
	// Min Length: 3
	NamingPattern *string `json:"namingPattern"`

	// The order string. Sorting request namings alphabetically by their order string determines their relative ordering.
	//
	// Typically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise.
	// Max Length: 2147483647
	// Min Length: 1
	Order string `json:"order,omitempty"`

	// The list of custom placeholders to be used in the naming pattern.
	//
	//  It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern.
	Placeholders []*Placeholder `json:"placeholders"`
}

// Validate validates this request naming
func (m *RequestNaming) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamingPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlaceholders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestNaming) validateConditions(formats strfmt.Registry) error {

	if err := validate.Required("conditions", "body", m.Conditions); err != nil {
		return err
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RequestNaming) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *RequestNaming) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RequestNaming) validateManagementZones(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagementZones) { // not required
		return nil
	}

	iManagementZonesSize := int64(len(m.ManagementZones))

	if err := validate.MinItems("managementZones", "body", iManagementZonesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("managementZones", "body", iManagementZonesSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *RequestNaming) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *RequestNaming) validateNamingPattern(formats strfmt.Registry) error {

	if err := validate.Required("namingPattern", "body", m.NamingPattern); err != nil {
		return err
	}

	if err := validate.MinLength("namingPattern", "body", string(*m.NamingPattern), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("namingPattern", "body", string(*m.NamingPattern), 1000); err != nil {
		return err
	}

	return nil
}

func (m *RequestNaming) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if err := validate.MinLength("order", "body", string(m.Order), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("order", "body", string(m.Order), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *RequestNaming) validatePlaceholders(formats strfmt.Registry) error {

	if swag.IsZero(m.Placeholders) { // not required
		return nil
	}

	for i := 0; i < len(m.Placeholders); i++ {
		if swag.IsZero(m.Placeholders[i]) { // not required
			continue
		}

		if m.Placeholders[i] != nil {
			if err := m.Placeholders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("placeholders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RequestNaming) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestNaming) UnmarshalBinary(b []byte) error {
	var res RequestNaming
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}