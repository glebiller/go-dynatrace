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

// EntityIDAlertingScope entity Id alerting scope
// swagger:model EntityIdAlertingScope
type EntityIDAlertingScope struct {

	// The monitored entities id to match on.
	// Required: true
	EntityID *string `json:"entityId"`
}

// FilterType gets the filter type of this subtype
func (m *EntityIDAlertingScope) FilterType() string {
	return "EntityIdAlertingScope"
}

// SetFilterType sets the filter type of this subtype
func (m *EntityIDAlertingScope) SetFilterType(val string) {

}

// EntityID gets the entity Id of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EntityIDAlertingScope) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The monitored entities id to match on.
		// Required: true
		EntityID *string `json:"entityId"`
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
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result EntityIDAlertingScope

	if base.FilterType != result.FilterType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid filterType value: %q", base.FilterType)
	}

	result.EntityID = data.EntityID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EntityIDAlertingScope) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The monitored entities id to match on.
		// Required: true
		EntityID *string `json:"entityId"`
	}{

		EntityID: m.EntityID,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		FilterType string `json:"filterType"`
	}{

		FilterType: m.FilterType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this entity Id alerting scope
func (m *EntityIDAlertingScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityIDAlertingScope) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityId", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityIDAlertingScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityIDAlertingScope) UnmarshalBinary(b []byte) error {
	var res EntityIDAlertingScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
