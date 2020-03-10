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

// GreaterThanCompareOperation greater than compare operation
// swagger:model GreaterThanCompareOperation
type GreaterThanCompareOperation struct {

	// The value to compare to.
	// Required: true
	Value *int32 `json:"value"`
}

// Type gets the type of this subtype
func (m *GreaterThanCompareOperation) Type() string {
	return "GreaterThanCompareOperation"
}

// SetType sets the type of this subtype
func (m *GreaterThanCompareOperation) SetType(val string) {

}

// Value gets the value of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GreaterThanCompareOperation) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The value to compare to.
		// Required: true
		Value *int32 `json:"value"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result GreaterThanCompareOperation

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Value = data.Value

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GreaterThanCompareOperation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The value to compare to.
		// Required: true
		Value *int32 `json:"value"`
	}{

		Value: m.Value,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this greater than compare operation
func (m *GreaterThanCompareOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GreaterThanCompareOperation) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GreaterThanCompareOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GreaterThanCompareOperation) UnmarshalBinary(b []byte) error {
	var res GreaterThanCompareOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}