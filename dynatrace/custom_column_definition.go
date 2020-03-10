// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomColumnDefinition custom column definition
// swagger:model CustomColumnDefinition
type CustomColumnDefinition struct {
	nameField *string

	// prefix
	// Required: true
	// Read Only: true
	// Min Length: 1
	Prefix string `json:"prefix"`

	// suffix
	// Required: true
	// Read Only: true
	Suffix string `json:"suffix"`
}

// Name gets the name of this subtype
func (m *CustomColumnDefinition) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *CustomColumnDefinition) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this subtype
func (m *CustomColumnDefinition) Type() string {
	return "CustomColumnDefinition"
}

// SetType sets the type of this subtype
func (m *CustomColumnDefinition) SetType(val string) {

}

// Prefix gets the prefix of this subtype

// Suffix gets the suffix of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CustomColumnDefinition) UnmarshalJSON(raw []byte) error {
	var data struct {

		// prefix
		// Required: true
		// Read Only: true
		// Min Length: 1
		Prefix string `json:"prefix"`

		// suffix
		// Required: true
		// Read Only: true
		Suffix string `json:"suffix"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name *string `json:"name"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result CustomColumnDefinition

	result.nameField = base.Name

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Prefix = data.Prefix

	result.Suffix = data.Suffix

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CustomColumnDefinition) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// prefix
		// Required: true
		// Read Only: true
		// Min Length: 1
		Prefix string `json:"prefix"`

		// suffix
		// Required: true
		// Read Only: true
		Suffix string `json:"suffix"`
	}{

		Prefix: m.Prefix,

		Suffix: m.Suffix,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name *string `json:"name"`

		Type string `json:"type"`
	}{

		Name: m.Name(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this custom column definition
func (m *CustomColumnDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuffix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomColumnDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name()), 1); err != nil {
		return err
	}

	return nil
}

func (m *CustomColumnDefinition) validatePrefix(formats strfmt.Registry) error {

	if err := validate.RequiredString("prefix", "body", string(m.Prefix)); err != nil {
		return err
	}

	if err := validate.MinLength("prefix", "body", string(m.Prefix), 1); err != nil {
		return err
	}

	return nil
}

func (m *CustomColumnDefinition) validateSuffix(formats strfmt.Registry) error {

	if err := validate.RequiredString("suffix", "body", string(m.Suffix)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomColumnDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomColumnDefinition) UnmarshalBinary(b []byte) error {
	var res CustomColumnDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
