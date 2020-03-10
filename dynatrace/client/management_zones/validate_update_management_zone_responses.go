// Code generated by go-swagger; DO NOT EDIT.

package management_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateManagementZoneReader is a Reader for the ValidateUpdateManagementZone structure.
type ValidateUpdateManagementZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateManagementZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateManagementZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateManagementZoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateManagementZoneNoContent creates a ValidateUpdateManagementZoneNoContent with default headers values
func NewValidateUpdateManagementZoneNoContent() *ValidateUpdateManagementZoneNoContent {
	return &ValidateUpdateManagementZoneNoContent{}
}

/*ValidateUpdateManagementZoneNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response does not have a body.
*/
type ValidateUpdateManagementZoneNoContent struct {
}

func (o *ValidateUpdateManagementZoneNoContent) Error() string {
	return fmt.Sprintf("[POST /managementZones/{id}/validator][%d] validateUpdateManagementZoneNoContent ", 204)
}

func (o *ValidateUpdateManagementZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateManagementZoneBadRequest creates a ValidateUpdateManagementZoneBadRequest with default headers values
func NewValidateUpdateManagementZoneBadRequest() *ValidateUpdateManagementZoneBadRequest {
	return &ValidateUpdateManagementZoneBadRequest{}
}

/*ValidateUpdateManagementZoneBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateManagementZoneBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateManagementZoneBadRequest) Error() string {
	return fmt.Sprintf("[POST /managementZones/{id}/validator][%d] validateUpdateManagementZoneBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateManagementZoneBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateManagementZoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
