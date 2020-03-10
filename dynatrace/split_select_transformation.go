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

// SplitSelectTransformation split select transformation
// swagger:model SplitSelectTransformation
type SplitSelectTransformation struct {

	// The delimiter for splitting the detected value. The delimiter itself is not kept.
	// Required: true
	Delimiter *string `json:"delimiter"`

	// The index of the element in the split array to be used. Indexing starts with `1`.
	// Required: true
	// Minimum: 1
	ItemIndex *int32 `json:"itemIndex"`
}

// Type gets the type of this subtype
func (m *SplitSelectTransformation) Type() string {
	return "SplitSelectTransformation"
}

// SetType sets the type of this subtype
func (m *SplitSelectTransformation) SetType(val string) {

}

// Delimiter gets the delimiter of this subtype

// ItemIndex gets the item index of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SplitSelectTransformation) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The delimiter for splitting the detected value. The delimiter itself is not kept.
		// Required: true
		Delimiter *string `json:"delimiter"`

		// The index of the element in the split array to be used. Indexing starts with `1`.
		// Required: true
		// Minimum: 1
		ItemIndex *int32 `json:"itemIndex"`
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

	var result SplitSelectTransformation

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Delimiter = data.Delimiter

	result.ItemIndex = data.ItemIndex

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SplitSelectTransformation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The delimiter for splitting the detected value. The delimiter itself is not kept.
		// Required: true
		Delimiter *string `json:"delimiter"`

		// The index of the element in the split array to be used. Indexing starts with `1`.
		// Required: true
		// Minimum: 1
		ItemIndex *int32 `json:"itemIndex"`
	}{

		Delimiter: m.Delimiter,

		ItemIndex: m.ItemIndex,
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

// Validate validates this split select transformation
func (m *SplitSelectTransformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelimiter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SplitSelectTransformation) validateDelimiter(formats strfmt.Registry) error {

	if err := validate.Required("delimiter", "body", m.Delimiter); err != nil {
		return err
	}

	return nil
}

func (m *SplitSelectTransformation) validateItemIndex(formats strfmt.Registry) error {

	if err := validate.Required("itemIndex", "body", m.ItemIndex); err != nil {
		return err
	}

	if err := validate.MinimumInt("itemIndex", "body", int64(*m.ItemIndex), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SplitSelectTransformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SplitSelectTransformation) UnmarshalBinary(b []byte) error {
	var res SplitSelectTransformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
