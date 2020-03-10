// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertingEventTypeFilter Configuration of the event filter for the alerting profile.
//
// You have two mutually exclusive options:
// * Select an event type from the list of the predefined events. Specify it in the **predefinedEventFilter** field.
// * Set a rule for custom events. Specify it in the **customEventFilter** field.
// swagger:model AlertingEventTypeFilter
type AlertingEventTypeFilter struct {

	// The custom event filter.
	//
	//  If several event filters specified, the OR logic applies.
	CustomEventFilter *AlertingCustomEventFilter `json:"customEventFilter,omitempty"`

	// The predefined event filter.
	//
	//  If several event filters specified, the OR logic applies.
	PredefinedEventFilter *AlertingPredefinedEventFilter `json:"predefinedEventFilter,omitempty"`
}

// Validate validates this alerting event type filter
func (m *AlertingEventTypeFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomEventFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePredefinedEventFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingEventTypeFilter) validateCustomEventFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomEventFilter) { // not required
		return nil
	}

	if m.CustomEventFilter != nil {
		if err := m.CustomEventFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customEventFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AlertingEventTypeFilter) validatePredefinedEventFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.PredefinedEventFilter) { // not required
		return nil
	}

	if m.PredefinedEventFilter != nil {
		if err := m.PredefinedEventFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("predefinedEventFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertingEventTypeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingEventTypeFilter) UnmarshalBinary(b []byte) error {
	var res AlertingEventTypeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
