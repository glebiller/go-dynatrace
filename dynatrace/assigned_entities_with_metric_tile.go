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

// AssignedEntitiesWithMetricTile assigned entities with metric tile
// swagger:model AssignedEntitiesWithMetricTile
type AssignedEntitiesWithMetricTile struct {
	boundsField *TileBounds

	configuredField bool

	nameField *string

	tileFilterField *TileFilter

	// The list of Dynatrace entities, assigned to the tile.
	// Required: true
	AssignedEntities []string `json:"assignedEntities"`

	// The metric assigned to the tile.
	Metric string `json:"metric,omitempty"`
}

// Bounds gets the bounds of this subtype
func (m *AssignedEntitiesWithMetricTile) Bounds() *TileBounds {
	return m.boundsField
}

// SetBounds sets the bounds of this subtype
func (m *AssignedEntitiesWithMetricTile) SetBounds(val *TileBounds) {
	m.boundsField = val
}

// Configured gets the configured of this subtype
func (m *AssignedEntitiesWithMetricTile) Configured() bool {
	return m.configuredField
}

// SetConfigured sets the configured of this subtype
func (m *AssignedEntitiesWithMetricTile) SetConfigured(val bool) {
	m.configuredField = val
}

// Name gets the name of this subtype
func (m *AssignedEntitiesWithMetricTile) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *AssignedEntitiesWithMetricTile) SetName(val *string) {
	m.nameField = val
}

// TileFilter gets the tile filter of this subtype
func (m *AssignedEntitiesWithMetricTile) TileFilter() *TileFilter {
	return m.tileFilterField
}

// SetTileFilter sets the tile filter of this subtype
func (m *AssignedEntitiesWithMetricTile) SetTileFilter(val *TileFilter) {
	m.tileFilterField = val
}

// TileType gets the tile type of this subtype
func (m *AssignedEntitiesWithMetricTile) TileType() string {
	return "AssignedEntitiesWithMetricTile"
}

// SetTileType sets the tile type of this subtype
func (m *AssignedEntitiesWithMetricTile) SetTileType(val string) {

}

// AssignedEntities gets the assigned entities of this subtype

// Metric gets the metric of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AssignedEntitiesWithMetricTile) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The list of Dynatrace entities, assigned to the tile.
		// Required: true
		AssignedEntities []string `json:"assignedEntities"`

		// The metric assigned to the tile.
		Metric string `json:"metric,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Bounds *TileBounds `json:"bounds"`

		Configured bool `json:"configured,omitempty"`

		Name *string `json:"name"`

		TileFilter *TileFilter `json:"tileFilter,omitempty"`

		TileType string `json:"tileType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AssignedEntitiesWithMetricTile

	result.boundsField = base.Bounds

	result.configuredField = base.Configured

	result.nameField = base.Name

	result.tileFilterField = base.TileFilter

	if base.TileType != result.TileType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid tileType value: %q", base.TileType)
	}

	result.AssignedEntities = data.AssignedEntities

	result.Metric = data.Metric

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AssignedEntitiesWithMetricTile) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The list of Dynatrace entities, assigned to the tile.
		// Required: true
		AssignedEntities []string `json:"assignedEntities"`

		// The metric assigned to the tile.
		Metric string `json:"metric,omitempty"`
	}{

		AssignedEntities: m.AssignedEntities,

		Metric: m.Metric,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Bounds *TileBounds `json:"bounds"`

		Configured bool `json:"configured,omitempty"`

		Name *string `json:"name"`

		TileFilter *TileFilter `json:"tileFilter,omitempty"`

		TileType string `json:"tileType"`
	}{

		Bounds: m.Bounds(),

		Configured: m.Configured(),

		Name: m.Name(),

		TileFilter: m.TileFilter(),

		TileType: m.TileType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this assigned entities with metric tile
func (m *AssignedEntitiesWithMetricTile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBounds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTileFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssignedEntitiesWithMetricTile) validateBounds(formats strfmt.Registry) error {

	if err := validate.Required("bounds", "body", m.Bounds()); err != nil {
		return err
	}

	if m.Bounds() != nil {
		if err := m.Bounds().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bounds")
			}
			return err
		}
	}

	return nil
}

func (m *AssignedEntitiesWithMetricTile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *AssignedEntitiesWithMetricTile) validateTileFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.TileFilter()) { // not required
		return nil
	}

	if m.TileFilter() != nil {
		if err := m.TileFilter().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tileFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AssignedEntitiesWithMetricTile) validateAssignedEntities(formats strfmt.Registry) error {

	if err := validate.Required("assignedEntities", "body", m.AssignedEntities); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssignedEntitiesWithMetricTile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignedEntitiesWithMetricTile) UnmarshalBinary(b []byte) error {
	var res AssignedEntitiesWithMetricTile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
