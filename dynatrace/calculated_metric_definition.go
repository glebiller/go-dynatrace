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

// CalculatedMetricDefinition The definition of a calculated service metric.
// swagger:model CalculatedMetricDefinition
type CalculatedMetricDefinition struct {

	// The metric to be captured.
	// Required: true
	// Enum: [CPU_TIME DATABASE_CHILD_CALL_COUNT DATABASE_CHILD_CALL_TIME EXCEPTION_COUNT FAILED_REQUEST_COUNT FAILED_REQUEST_COUNT_CLIENT FAILURE_RATE FAILURE_RATE_CLIENT HTTP_4XX_ERROR_COUNT HTTP_4XX_ERROR_COUNT_CLIENT HTTP_5XX_ERROR_COUNT HTTP_5XX_ERROR_COUNT_CLIENT IO_TIME LOCK_TIME NON_DATABASE_CHILD_CALL_COUNT NON_DATABASE_CHILD_CALL_TIME PROCESSING_TIME REQUEST_ATTRIBUTE REQUEST_COUNT RESPONSE_TIME RESPONSE_TIME_CLIENT SUCCESSFUL_REQUEST_COUNT SUCCESSFUL_REQUEST_COUNT_CLIENT WAIT_TIME]
	Metric *string `json:"metric"`

	// The request attribute to be captured.
	//
	//  Only applicable when the **metric** parameter is set to `REQUEST_ATTRIBUTE`.
	RequestAttribute string `json:"requestAttribute,omitempty"`
}

// Validate validates this calculated metric definition
func (m *CalculatedMetricDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var calculatedMetricDefinitionTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CPU_TIME","DATABASE_CHILD_CALL_COUNT","DATABASE_CHILD_CALL_TIME","EXCEPTION_COUNT","FAILED_REQUEST_COUNT","FAILED_REQUEST_COUNT_CLIENT","FAILURE_RATE","FAILURE_RATE_CLIENT","HTTP_4XX_ERROR_COUNT","HTTP_4XX_ERROR_COUNT_CLIENT","HTTP_5XX_ERROR_COUNT","HTTP_5XX_ERROR_COUNT_CLIENT","IO_TIME","LOCK_TIME","NON_DATABASE_CHILD_CALL_COUNT","NON_DATABASE_CHILD_CALL_TIME","PROCESSING_TIME","REQUEST_ATTRIBUTE","REQUEST_COUNT","RESPONSE_TIME","RESPONSE_TIME_CLIENT","SUCCESSFUL_REQUEST_COUNT","SUCCESSFUL_REQUEST_COUNT_CLIENT","WAIT_TIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		calculatedMetricDefinitionTypeMetricPropEnum = append(calculatedMetricDefinitionTypeMetricPropEnum, v)
	}
}

const (

	// CalculatedMetricDefinitionMetricCPUTIME captures enum value "CPU_TIME"
	CalculatedMetricDefinitionMetricCPUTIME string = "CPU_TIME"

	// CalculatedMetricDefinitionMetricDATABASECHILDCALLCOUNT captures enum value "DATABASE_CHILD_CALL_COUNT"
	CalculatedMetricDefinitionMetricDATABASECHILDCALLCOUNT string = "DATABASE_CHILD_CALL_COUNT"

	// CalculatedMetricDefinitionMetricDATABASECHILDCALLTIME captures enum value "DATABASE_CHILD_CALL_TIME"
	CalculatedMetricDefinitionMetricDATABASECHILDCALLTIME string = "DATABASE_CHILD_CALL_TIME"

	// CalculatedMetricDefinitionMetricEXCEPTIONCOUNT captures enum value "EXCEPTION_COUNT"
	CalculatedMetricDefinitionMetricEXCEPTIONCOUNT string = "EXCEPTION_COUNT"

	// CalculatedMetricDefinitionMetricFAILEDREQUESTCOUNT captures enum value "FAILED_REQUEST_COUNT"
	CalculatedMetricDefinitionMetricFAILEDREQUESTCOUNT string = "FAILED_REQUEST_COUNT"

	// CalculatedMetricDefinitionMetricFAILEDREQUESTCOUNTCLIENT captures enum value "FAILED_REQUEST_COUNT_CLIENT"
	CalculatedMetricDefinitionMetricFAILEDREQUESTCOUNTCLIENT string = "FAILED_REQUEST_COUNT_CLIENT"

	// CalculatedMetricDefinitionMetricFAILURERATE captures enum value "FAILURE_RATE"
	CalculatedMetricDefinitionMetricFAILURERATE string = "FAILURE_RATE"

	// CalculatedMetricDefinitionMetricFAILURERATECLIENT captures enum value "FAILURE_RATE_CLIENT"
	CalculatedMetricDefinitionMetricFAILURERATECLIENT string = "FAILURE_RATE_CLIENT"

	// CalculatedMetricDefinitionMetricHTTP4XXERRORCOUNT captures enum value "HTTP_4XX_ERROR_COUNT"
	CalculatedMetricDefinitionMetricHTTP4XXERRORCOUNT string = "HTTP_4XX_ERROR_COUNT"

	// CalculatedMetricDefinitionMetricHTTP4XXERRORCOUNTCLIENT captures enum value "HTTP_4XX_ERROR_COUNT_CLIENT"
	CalculatedMetricDefinitionMetricHTTP4XXERRORCOUNTCLIENT string = "HTTP_4XX_ERROR_COUNT_CLIENT"

	// CalculatedMetricDefinitionMetricHTTP5XXERRORCOUNT captures enum value "HTTP_5XX_ERROR_COUNT"
	CalculatedMetricDefinitionMetricHTTP5XXERRORCOUNT string = "HTTP_5XX_ERROR_COUNT"

	// CalculatedMetricDefinitionMetricHTTP5XXERRORCOUNTCLIENT captures enum value "HTTP_5XX_ERROR_COUNT_CLIENT"
	CalculatedMetricDefinitionMetricHTTP5XXERRORCOUNTCLIENT string = "HTTP_5XX_ERROR_COUNT_CLIENT"

	// CalculatedMetricDefinitionMetricIOTIME captures enum value "IO_TIME"
	CalculatedMetricDefinitionMetricIOTIME string = "IO_TIME"

	// CalculatedMetricDefinitionMetricLOCKTIME captures enum value "LOCK_TIME"
	CalculatedMetricDefinitionMetricLOCKTIME string = "LOCK_TIME"

	// CalculatedMetricDefinitionMetricNONDATABASECHILDCALLCOUNT captures enum value "NON_DATABASE_CHILD_CALL_COUNT"
	CalculatedMetricDefinitionMetricNONDATABASECHILDCALLCOUNT string = "NON_DATABASE_CHILD_CALL_COUNT"

	// CalculatedMetricDefinitionMetricNONDATABASECHILDCALLTIME captures enum value "NON_DATABASE_CHILD_CALL_TIME"
	CalculatedMetricDefinitionMetricNONDATABASECHILDCALLTIME string = "NON_DATABASE_CHILD_CALL_TIME"

	// CalculatedMetricDefinitionMetricPROCESSINGTIME captures enum value "PROCESSING_TIME"
	CalculatedMetricDefinitionMetricPROCESSINGTIME string = "PROCESSING_TIME"

	// CalculatedMetricDefinitionMetricREQUESTATTRIBUTE captures enum value "REQUEST_ATTRIBUTE"
	CalculatedMetricDefinitionMetricREQUESTATTRIBUTE string = "REQUEST_ATTRIBUTE"

	// CalculatedMetricDefinitionMetricREQUESTCOUNT captures enum value "REQUEST_COUNT"
	CalculatedMetricDefinitionMetricREQUESTCOUNT string = "REQUEST_COUNT"

	// CalculatedMetricDefinitionMetricRESPONSETIME captures enum value "RESPONSE_TIME"
	CalculatedMetricDefinitionMetricRESPONSETIME string = "RESPONSE_TIME"

	// CalculatedMetricDefinitionMetricRESPONSETIMECLIENT captures enum value "RESPONSE_TIME_CLIENT"
	CalculatedMetricDefinitionMetricRESPONSETIMECLIENT string = "RESPONSE_TIME_CLIENT"

	// CalculatedMetricDefinitionMetricSUCCESSFULREQUESTCOUNT captures enum value "SUCCESSFUL_REQUEST_COUNT"
	CalculatedMetricDefinitionMetricSUCCESSFULREQUESTCOUNT string = "SUCCESSFUL_REQUEST_COUNT"

	// CalculatedMetricDefinitionMetricSUCCESSFULREQUESTCOUNTCLIENT captures enum value "SUCCESSFUL_REQUEST_COUNT_CLIENT"
	CalculatedMetricDefinitionMetricSUCCESSFULREQUESTCOUNTCLIENT string = "SUCCESSFUL_REQUEST_COUNT_CLIENT"

	// CalculatedMetricDefinitionMetricWAITTIME captures enum value "WAIT_TIME"
	CalculatedMetricDefinitionMetricWAITTIME string = "WAIT_TIME"
)

// prop value enum
func (m *CalculatedMetricDefinition) validateMetricEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, calculatedMetricDefinitionTypeMetricPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CalculatedMetricDefinition) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", *m.Metric); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalculatedMetricDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculatedMetricDefinition) UnmarshalBinary(b []byte) error {
	var res CalculatedMetricDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}