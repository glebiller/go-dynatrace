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

// ServiceTopologyComparison service topology comparison
// swagger:model ServiceTopologyComparison
type ServiceTopologyComparison struct {
	negateField *bool

	operatorField Enum

	valueField interface{}
}

// Negate gets the negate of this subtype
func (m *ServiceTopologyComparison) Negate() *bool {
	return m.negateField
}

// SetNegate sets the negate of this subtype
func (m *ServiceTopologyComparison) SetNegate(val *bool) {
	m.negateField = val
}

// Operator gets the operator of this subtype
func (m *ServiceTopologyComparison) Operator() Enum {
	return m.operatorField
}

// SetOperator sets the operator of this subtype
func (m *ServiceTopologyComparison) SetOperator(val Enum) {
	m.operatorField = val
}

// Type gets the type of this subtype
func (m *ServiceTopologyComparison) Type() string {
	return "ServiceTopologyComparison"
}

// SetType sets the type of this subtype
func (m *ServiceTopologyComparison) SetType(val string) {

}

// Value gets the value of this subtype
func (m *ServiceTopologyComparison) Value() interface{} {
	return m.valueField
}

// SetValue sets the value of this subtype
func (m *ServiceTopologyComparison) SetValue(val interface{}) {
	m.valueField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ServiceTopologyComparison) UnmarshalJSON(raw []byte) error {
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

	var result ServiceTopologyComparison

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
func (m ServiceTopologyComparison) MarshalJSON() ([]byte, error) {
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

// Validate validates this service topology comparison
func (m *ServiceTopologyComparison) Validate(formats strfmt.Registry) error {
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

func (m *ServiceTopologyComparison) validateNegate(formats strfmt.Registry) error {

	if err := validate.Required("negate", "body", m.Negate()); err != nil {
		return err
	}

	return nil
}

func (m *ServiceTopologyComparison) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceTopologyComparison) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceTopologyComparison) UnmarshalBinary(b []byte) error {
	var res ServiceTopologyComparison
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
