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

// FailureRateIncreaseDetectionConfig Configuration of failure rate increase detection.
// swagger:model FailureRateIncreaseDetectionConfig
type FailureRateIncreaseDetectionConfig struct {

	// Parameters of automatic failure rate increase detection.
	//
	//  Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
	//
	// The absolute and relative thresholds **both** must be exceeded to trigger an alert.
	//
	// Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and you set a relative increase of 50%, the thresholds will be:
	// Absolute: 1.5% + **1%** = 2.5%
	// Relative: 1.5% + 1.5% * **50%** = 2.25%
	AutomaticDetection *FailureRateIncreaseAutodetectionConfig `json:"automaticDetection,omitempty"`

	// How to detect failure rate increase: automatically, or based on fixed thresholds, or do not detect.
	// Required: true
	// Enum: [DETECT_AUTOMATICALLY DETECT_USING_FIXED_THRESHOLDS DONT_DETECT]
	DetectionMode *string `json:"detectionMode"`

	// Fixed thresholds for failure rate increase detection.
	//
	//  Required if the **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
	Thresholds *FailureRateIncreaseThresholdConfig `json:"thresholds,omitempty"`
}

// Validate validates this failure rate increase detection config
func (m *FailureRateIncreaseDetectionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutomaticDetection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetectionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FailureRateIncreaseDetectionConfig) validateAutomaticDetection(formats strfmt.Registry) error {

	if swag.IsZero(m.AutomaticDetection) { // not required
		return nil
	}

	if m.AutomaticDetection != nil {
		if err := m.AutomaticDetection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automaticDetection")
			}
			return err
		}
	}

	return nil
}

var failureRateIncreaseDetectionConfigTypeDetectionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DETECT_AUTOMATICALLY","DETECT_USING_FIXED_THRESHOLDS","DONT_DETECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		failureRateIncreaseDetectionConfigTypeDetectionModePropEnum = append(failureRateIncreaseDetectionConfigTypeDetectionModePropEnum, v)
	}
}

const (

	// FailureRateIncreaseDetectionConfigDetectionModeDETECTAUTOMATICALLY captures enum value "DETECT_AUTOMATICALLY"
	FailureRateIncreaseDetectionConfigDetectionModeDETECTAUTOMATICALLY string = "DETECT_AUTOMATICALLY"

	// FailureRateIncreaseDetectionConfigDetectionModeDETECTUSINGFIXEDTHRESHOLDS captures enum value "DETECT_USING_FIXED_THRESHOLDS"
	FailureRateIncreaseDetectionConfigDetectionModeDETECTUSINGFIXEDTHRESHOLDS string = "DETECT_USING_FIXED_THRESHOLDS"

	// FailureRateIncreaseDetectionConfigDetectionModeDONTDETECT captures enum value "DONT_DETECT"
	FailureRateIncreaseDetectionConfigDetectionModeDONTDETECT string = "DONT_DETECT"
)

// prop value enum
func (m *FailureRateIncreaseDetectionConfig) validateDetectionModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, failureRateIncreaseDetectionConfigTypeDetectionModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FailureRateIncreaseDetectionConfig) validateDetectionMode(formats strfmt.Registry) error {

	if err := validate.Required("detectionMode", "body", m.DetectionMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateDetectionModeEnum("detectionMode", "body", *m.DetectionMode); err != nil {
		return err
	}

	return nil
}

func (m *FailureRateIncreaseDetectionConfig) validateThresholds(formats strfmt.Registry) error {

	if swag.IsZero(m.Thresholds) { // not required
		return nil
	}

	if m.Thresholds != nil {
		if err := m.Thresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thresholds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FailureRateIncreaseDetectionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailureRateIncreaseDetectionConfig) UnmarshalBinary(b []byte) error {
	var res FailureRateIncreaseDetectionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
