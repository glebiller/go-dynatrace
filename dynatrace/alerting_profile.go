// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertingProfile Configuration of an alerting profile.
// swagger:model AlertingProfile
type AlertingProfile struct {

	// The name of the alerting profile, displayed in the UI.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	DisplayName *string `json:"displayName"`

	// The list of event filters.
	//
	//  If several filters are specified, the OR logic applies.
	//
	//  If you specify both severity rule and event filter, the AND logic applies.
	// Max Items: 20
	// Min Items: 0
	EventTypeFilters []*AlertingEventTypeFilter `json:"eventTypeFilters"`

	// The ID of the alerting profile.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The ID of the management zone to which the alerting profile applies.
	ManagementZoneID int64 `json:"managementZoneId,omitempty"`

	// Metadata useful for debugging.
	// Read Only: true
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`

	// A list of severity rules.
	//
	//  The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops.
	//
	//  If you specify both severity rule and event filter, the AND logic applies.
	// Max Items: 20
	// Min Items: 0
	Rules []*AlertingProfileSeverityRule `json:"rules"`
}

// Validate validates this alerting profile
func (m *AlertingProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTypeFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfile) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.MinLength("displayName", "body", string(*m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(*m.DisplayName), 100); err != nil {
		return err
	}

	return nil
}

func (m *AlertingProfile) validateEventTypeFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.EventTypeFilters) { // not required
		return nil
	}

	iEventTypeFiltersSize := int64(len(m.EventTypeFilters))

	if err := validate.MinItems("eventTypeFilters", "body", iEventTypeFiltersSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("eventTypeFilters", "body", iEventTypeFiltersSize, 20); err != nil {
		return err
	}

	for i := 0; i < len(m.EventTypeFilters); i++ {
		if swag.IsZero(m.EventTypeFilters[i]) { // not required
			continue
		}

		if m.EventTypeFilters[i] != nil {
			if err := m.EventTypeFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventTypeFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingProfile) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AlertingProfile) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *AlertingProfile) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	iRulesSize := int64(len(m.Rules))

	if err := validate.MinItems("rules", "body", iRulesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("rules", "body", iRulesSize, 20); err != nil {
		return err
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfile) UnmarshalBinary(b []byte) error {
	var res AlertingProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
