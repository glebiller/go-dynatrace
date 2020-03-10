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

// NumberComparisonInfo number comparison info
// swagger:model NumberComparisonInfo
type NumberComparisonInfo struct {
	comparisonField Enum

	negateField *bool

	valueField interface{}
}

// Comparison gets the comparison of this subtype
func (m *NumberComparisonInfo) Comparison() Enum {
	return m.comparisonField
}

// SetComparison sets the comparison of this subtype
func (m *NumberComparisonInfo) SetComparison(val Enum) {
	m.comparisonField = val
}

// Negate gets the negate of this subtype
func (m *NumberComparisonInfo) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this subtype
func (m *NumberComparisonInfo) SetNegate(val *bool) {
	m.negateField = val
}

// Type gets the type of this subtype
func (m *NumberComparisonInfo) Type() string {
	return "NumberComparisonInfo"
}

// SetType sets the type of this subtype
func (m *NumberComparisonInfo) SetType(val string) {

}

// Value gets the value of this subtype
func (m *NumberComparisonInfo) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this subtype
func (m *NumberComparisonInfo) SetValue(val interface{}) {
	m.valueField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NumberComparisonInfo) UnmarshalJSON(raw []byte) error {
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

		Comparison Enum `json:"comparison"`

		Negate *bool `json:"negate"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result NumberComparisonInfo

	result.comparisonField = base.Comparison

	result.negateField = base.Negate

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.valueField = base.Value

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NumberComparisonInfo) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Comparison Enum `json:"comparison"`

		Negate *bool `json:"negate"`

		Type string `json:"type"`

		Value interface{} `json:"value,omitempty"`
	}{

		Comparison: m.Comparison(),

		Negate: m.Negate(),

		Type: m.Type(),

		Value: m.Value(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this number comparison info
func (m *NumberComparisonInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparison(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegate(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NumberComparisonInfo) validateComparison(formats strfmt.Registry) error {

	if err := validate.Required("comparison", "body", m.Comparison()); err != nil {
		return err
	}

	return nil
}

func (m *NumberComparisonInfo) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NumberComparisonInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NumberComparisonInfo) UnmarshalBinary(b []byte) error {
	var res NumberComparisonInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
