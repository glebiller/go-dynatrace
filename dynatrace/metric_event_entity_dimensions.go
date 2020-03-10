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

// MetricEventEntityDimensions metric event entity dimensions
// swagger:model MetricEventEntityDimensions
type MetricEventEntityDimensions struct {
	indexField *int32

	nameField string

	// The name filter which should be applied for the dimensions.
	// Required: true
	NameFilter *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto `json:"nameFilter"`
}

// FilterType gets the filter type of this subtype
func (m *MetricEventEntityDimensions) FilterType() string {
	return "MetricEventEntityDimensions"
}

// SetFilterType sets the filter type of this subtype
func (m *MetricEventEntityDimensions) SetFilterType(val string) {

}

// Index gets the index of this subtype
func (m *MetricEventEntityDimensions) Index() *int32 {
	return m.indexField
}

// SetIndex sets the index of this subtype
func (m *MetricEventEntityDimensions) SetIndex(val *int32) {
	m.indexField = val
}

// Name gets the name of this subtype
func (m *MetricEventEntityDimensions) Name() string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *MetricEventEntityDimensions) SetName(val string) {
	m.nameField = val
}

// NameFilter gets the name filter of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MetricEventEntityDimensions) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The name filter which should be applied for the dimensions.
		// Required: true
		NameFilter *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto `json:"nameFilter"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		FilterType string `json:"filterType"`

		Index *int32 `json:"index"`

		Name string `json:"name,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MetricEventEntityDimensions

	if base.FilterType != result.FilterType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid filterType value: %q", base.FilterType)
	}

	result.indexField = base.Index

	result.nameField = base.Name

	result.NameFilter = data.NameFilter

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MetricEventEntityDimensions) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The name filter which should be applied for the dimensions.
		// Required: true
		NameFilter *MetricEventTextFilterMetricEventDimensionsFilterOperatorDto `json:"nameFilter"`
	}{

		NameFilter: m.NameFilter,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		FilterType string `json:"filterType"`

		Index *int32 `json:"index"`

		Name string `json:"name,omitempty"`
	}{

		FilterType: m.FilterType(),

		Index: m.Index(),

		Name: m.Name(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this metric event entity dimensions
func (m *MetricEventEntityDimensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricEventEntityDimensions) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index()); err != nil {
		return err
	}

	return nil
}

func (m *MetricEventEntityDimensions) validateNameFilter(formats strfmt.Registry) error {

	if err := validate.Required("nameFilter", "body", m.NameFilter); err != nil {
		return err
	}

	if m.NameFilter != nil {
		if err := m.NameFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nameFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricEventEntityDimensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricEventEntityDimensions) UnmarshalBinary(b []byte) error {
	var res MetricEventEntityDimensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
