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

// SimpleHostTechComparison simple host tech comparison
// swagger:model SimpleHostTechComparison
type SimpleHostTechComparison struct {
	negateField *bool

	operatorField Enum

	valueField interface{}
}

// Negate gets the negate of this subtype
func (m *SimpleHostTechComparison) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this subtype
func (m *SimpleHostTechComparison) SetNegate(val *bool) {
	m.negateField = val
}

// Operator gets the operator of this subtype
func (m *SimpleHostTechComparison) Operator() Enum {
	return m.operatorField
}

// SetOperator sets the operator of this subtype
func (m *SimpleHostTechComparison) SetOperator(val Enum) {
	m.operatorField = val
}

// Type gets the type of this subtype
func (m *SimpleHostTechComparison) Type() string {
	return "SimpleHostTechComparison"
}

// SetType sets the type of this subtype
func (m *SimpleHostTechComparison) SetType(val string) {

}

// Value gets the value of this subtype
func (m *SimpleHostTechComparison) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this subtype
func (m *SimpleHostTechComparison) SetValue(val interface{}) {
	m.valueField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SimpleHostTechComparison) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Negate *bool `json:"negate"`

		Operator Enum `json:"operator"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result SimpleHostTechComparison

	result.negateField = base.Negate

	result.operatorField = base.Operator

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.valueField = base.Value

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SimpleHostTechComparison) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Negate *bool `json:"negate"`

		Operator Enum `json:"operator"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}{

		Negate: m.Negate(),

		Operator: m.Operator(),

		Type: m.Type(),

		Value: m.Value(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this simple host tech comparison
func (m *SimpleHostTechComparison) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNegate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleHostTechComparison) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

func (m *SimpleHostTechComparison) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleHostTechComparison) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleHostTechComparison) UnmarshalBinary(b []byte) error {
	var res SimpleHostTechComparison
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
