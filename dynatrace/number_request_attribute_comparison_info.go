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

// NumberRequestAttributeComparisonInfo number request attribute comparison info
// swagger:model NumberRequestAttributeComparisonInfo
type NumberRequestAttributeComparisonInfo struct {
	comparisonField Enum

	negateField *bool

	valueField interface{}

	// request attribute
	// Required: true
	RequestAttribute *string `json:"requestAttribute"`
}

// Comparison gets the comparison of this subtype
func (m *NumberRequestAttributeComparisonInfo) Comparison() Enum {
	return m.comparisonField
}

// SetComparison sets the comparison of this subtype
func (m *NumberRequestAttributeComparisonInfo) SetComparison(val Enum) {
	m.comparisonField = val
}

// Negate gets the negate of this subtype
func (m *NumberRequestAttributeComparisonInfo) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this subtype
func (m *NumberRequestAttributeComparisonInfo) SetNegate(val *bool) {
	m.negateField = val
}

// Type gets the type of this subtype
func (m *NumberRequestAttributeComparisonInfo) Type() string {
	return "NumberRequestAttributeComparisonInfo"
}

// SetType sets the type of this subtype
func (m *NumberRequestAttributeComparisonInfo) SetType(val string) {

}

// Value gets the value of this subtype
func (m *NumberRequestAttributeComparisonInfo) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this subtype
func (m *NumberRequestAttributeComparisonInfo) SetValue(val interface{}) {
	m.valueField = val
}

// RequestAttribute gets the request attribute of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NumberRequestAttributeComparisonInfo) UnmarshalJSON(raw []byte) error {
	var data struct {

		// request attribute
		// Required: true
		RequestAttribute *string `json:"requestAttribute"`
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

	var result NumberRequestAttributeComparisonInfo

	result.comparisonField = base.Comparison

	result.negateField = base.Negate

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.valueField = base.Value

	result.RequestAttribute = data.RequestAttribute

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NumberRequestAttributeComparisonInfo) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// request attribute
		// Required: true
		RequestAttribute *string `json:"requestAttribute"`
	}{

		RequestAttribute: m.RequestAttribute,
	},
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

// Validate validates this number request attribute comparison info
func (m *NumberRequestAttributeComparisonInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparison(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNegate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestAttribute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NumberRequestAttributeComparisonInfo) validateComparison(formats strfmt.Registry) error {

	if err := validate.Required("comparison", "body", m.Comparison()); err != nil {
		return err
	}

	return nil
}

func (m *NumberRequestAttributeComparisonInfo) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

func (m *NumberRequestAttributeComparisonInfo) validateRequestAttribute(formats strfmt.Registry) error {

	if err := validate.Required("requestAttribute", "body", m.RequestAttribute); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NumberRequestAttributeComparisonInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NumberRequestAttributeComparisonInfo) UnmarshalBinary(b []byte) error {
	var res NumberRequestAttributeComparisonInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
