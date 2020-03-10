// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateConfigurationReader is a Reader for the ValidateUpdateConfiguration structure.
type ValidateUpdateConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateConfigurationNoContent creates a ValidateUpdateConfigurationNoContent with default headers values
func NewValidateUpdateConfigurationNoContent() *ValidateUpdateConfigurationNoContent {
	return &ValidateUpdateConfigurationNoContent{}
}

/*ValidateUpdateConfigurationNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body
*/
type ValidateUpdateConfigurationNoContent struct {
}

func (o *ValidateUpdateConfigurationNoContent) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows/{id}/validator][%d] validateUpdateConfigurationNoContent ", 204)
}

func (o *ValidateUpdateConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateConfigurationBadRequest creates a ValidateUpdateConfigurationBadRequest with default headers values
func NewValidateUpdateConfigurationBadRequest() *ValidateUpdateConfigurationBadRequest {
	return &ValidateUpdateConfigurationBadRequest{}
}

/*ValidateUpdateConfigurationBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateConfigurationBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows/{id}/validator][%d] validateUpdateConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateConfigurationBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
