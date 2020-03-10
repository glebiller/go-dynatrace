// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationFilter The condition of an application detection rule.
// swagger:model ApplicationFilter
type ApplicationFilter struct {

	// Where to look for the the **pattern** value.
	// Required: true
	// Enum: [DOMAIN URL]
	ApplicationMatchTarget *string `json:"applicationMatchTarget"`

	// The operator of the matching.
	// Required: true
	// Enum: [BEGINS_WITH CONTAINS ENDS_WITH EQUALS MATCHES]
	ApplicationMatchType *string `json:"applicationMatchType"`

	// The value to look for.
	// Required: true
	// Max Length: 200
	// Min Length: 1
	Pattern *string `json:"pattern"`
}

// Validate validates this application filter
func (m *ApplicationFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationMatchTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationMatchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationFilterTypeApplicationMatchTargetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOMAIN","URL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationFilterTypeApplicationMatchTargetPropEnum = append(applicationFilterTypeApplicationMatchTargetPropEnum, v)
	}
}

const (

	// ApplicationFilterApplicationMatchTargetDOMAIN captures enum value "DOMAIN"
	ApplicationFilterApplicationMatchTargetDOMAIN string = "DOMAIN"

	// ApplicationFilterApplicationMatchTargetURL captures enum value "URL"
	ApplicationFilterApplicationMatchTargetURL string = "URL"
)

// prop value enum
func (m *ApplicationFilter) validateApplicationMatchTargetEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationFilterTypeApplicationMatchTargetPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationFilter) validateApplicationMatchTarget(formats strfmt.Registry) error {

	if err := validate.Required("applicationMatchTarget", "body", m.ApplicationMatchTarget); err != nil {
		return err
	}

	// value enum
	if err := m.validateApplicationMatchTargetEnum("applicationMatchTarget", "body", *m.ApplicationMatchTarget); err != nil {
		return err
	}

	return nil
}

var applicationFilterTypeApplicationMatchTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BEGINS_WITH","CONTAINS","ENDS_WITH","EQUALS","MATCHES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationFilterTypeApplicationMatchTypePropEnum = append(applicationFilterTypeApplicationMatchTypePropEnum, v)
	}
}

const (

	// ApplicationFilterApplicationMatchTypeBEGINSWITH captures enum value "BEGINS_WITH"
	ApplicationFilterApplicationMatchTypeBEGINSWITH string = "BEGINS_WITH"

	// ApplicationFilterApplicationMatchTypeCONTAINS captures enum value "CONTAINS"
	ApplicationFilterApplicationMatchTypeCONTAINS string = "CONTAINS"

	// ApplicationFilterApplicationMatchTypeENDSWITH captures enum value "ENDS_WITH"
	ApplicationFilterApplicationMatchTypeENDSWITH string = "ENDS_WITH"

	// ApplicationFilterApplicationMatchTypeEQUALS captures enum value "EQUALS"
	ApplicationFilterApplicationMatchTypeEQUALS string = "EQUALS"

	// ApplicationFilterApplicationMatchTypeMATCHES captures enum value "MATCHES"
	ApplicationFilterApplicationMatchTypeMATCHES string = "MATCHES"
)

// prop value enum
func (m *ApplicationFilter) validateApplicationMatchTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationFilterTypeApplicationMatchTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationFilter) validateApplicationMatchType(formats strfmt.Registry) error {

	if err := validate.Required("applicationMatchType", "body", m.ApplicationMatchType); err != nil {
		return err
	}

	// value enum
	if err := m.validateApplicationMatchTypeEnum("applicationMatchType", "body", *m.ApplicationMatchType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationFilter) validatePattern(formats strfmt.Registry) error {

	if err := validate.Required("pattern", "body", m.Pattern); err != nil {
		return err
	}

	if err := validate.MinLength("pattern", "body", string(*m.Pattern), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("pattern", "body", string(*m.Pattern), 200); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationFilter) UnmarshalBinary(b []byte) error {
	var res ApplicationFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}